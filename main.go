package main

import (
	"fmt"
	"os"
	"github.com/urfave/cli"
    "errors"
    "path/filepath"
    "strings"
	"github.com/disintegration/imaging"
    "path"
    "crypto/md5"
    "bytes"
    "image"
)

const namearg = "name"
const sourcearg = "source"
const targetarg = "target"
const baseurlarg = "baseurl"
const disqusarg = "disqus"
const thumbwidtharg = "thumbwidth"
const thumbheightarg = "thumbheight"
const viewerwidtharg = "viewerwidth"
const viewerheightarg = "viewerheight"
const skipextcheckarg = "skipextcheck"
const debugarg = "debug"

const thumbnail = "thumbnail.jpg"
const cacheFolder = "cache"
const filemode = 0755

var options Options
var masterAlbum *Album
var filesTouched int

func main() {
	app := cli.NewApp()
	app.Name = "ssgallery"
	app.Usage = "stupidly simple gallery"

    options = Options{}

	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: namearg,
			Usage: "Set gallery name to `NAME`",
			Destination: &options.name,
		},
		cli.StringFlag{
			Name: sourcearg,
			Usage: "Read albums from `FOLDER`",
			Destination: &options.source,
		},
		cli.StringFlag{
			Name: targetarg,
			Usage: "Write web page to `FOLDER`",
			Destination: &options.target,
		},
		cli.StringFlag{
			Name: baseurlarg,
			Usage: "Gallery will be hosted at http://server/`REL_URL`",
			Destination: &options.baseurl,
		},
		cli.StringFlag{
			Name: disqusarg,
			Usage: "[Optional] Custom `URL` for disqus topic (go to disqus.com to get one)",
			Destination: &options.disqus,
		},
		cli.IntFlag{
			Name: thumbwidtharg,
			Usage: "Set max thumbnail width to `WIDTH`",
			Destination: &options.thumbwidth,
		},
		cli.IntFlag{
			Name: thumbheightarg,
			Usage: "Set max thumbnail height to `HEIGHT`",
			Destination: &options.thumbheight,
		},
		cli.IntFlag{
			Name: viewerwidtharg,
			Usage: "Set max image viewer width to `WIDTH`",
			Destination: &options.viewerwidth,
		},
		cli.IntFlag{
			Name: viewerheightarg,
			Usage: "Set max image viewer height to `HEIGHT`",
			Destination: &options.viewerheight,
		},
        cli.BoolFlag{
            Name: skipextcheckarg,
            Usage: "Skip the check that limits to known supported file extensions (may include more images but may be slower if there are non-image files in the source folder)",
            Destination: &options.skipextcheck,
        },
        cli.BoolFlag{
            Name: debugarg,
            Usage: "Enable debug logging",
            Destination: &options.debug,
        },
	}

	app.Action = runApp

	app.Run(os.Args)
}

func runApp(c *cli.Context) error {
	var args []string = []string{namearg, sourcearg, targetarg, baseurlarg, thumbwidtharg, thumbheightarg, viewerwidtharg, viewerheightarg}

	for _,arg := range args {
		if (!c.IsSet(arg)) {
			cli.ShowAppHelp(c)
			return cli.NewExitError(fmt.Sprintf("\n\nArgument '%s' is required", arg), 1)
		}
	}

	fmt.Printf("\narguments:\nname: '%s'\nsource: '%s'\ntarget: '%s'\nbaseurl: '%s'\ndisqus: '%s'" +
																						"\nthumbwidth: %d\nthumbheight: %d\nviewerwidth: %d\nviewerheight: %d\n",
		options.name, options.source, options.target, options.baseurl, options.disqus, options.thumbwidth,
		options.thumbheight, options.viewerwidth, options.viewerheight)

	if !strings.HasPrefix(options.baseurl, "/") {
		options.baseurl = fmt.Sprintf("/%s", options.baseurl)
		fmt.Printf("Warning: baseurl does not include leading slash. I added one for you. New baseurl = '%s'\n", options.baseurl)
	}

	BuildGallery()
	CopyResources()

	if !masterAlbum.HasImages() {
		return fmt.Errorf("No images found in %s or any of its subdirectories. Exiting.", masterAlbum.folder)
	}

	masterAlbum.UpdateImageRenditions(options.target)
	masterAlbum.UpdatePages(options.target, options.baseurl)

	fmt.Printf("%d files touched (not including contents of %s)\n", filesTouched, filepath.Join(options.target, "data"))

	return nil
}

func BuildGallery() {
    if exists, _ := exists(options.source); !exists {
        panic(errors.New(fmt.Sprintf("could not find path %s", options.source)))
    }

    masterAlbum = NewTopAlbum(options.name, options.source, options.baseurl, nil)

    masterAlbum.LoadAlbum(options.source)
}

func CopyResources() {
    _ = os.Mkdir(options.target, filemode)
    RestoreAssets(options.target, "data")
}

func SaveResizedImage(imageInfo *Image, width, height int, filename string, isThumb, skipIfNewer bool) {
	printErr := func(err error) {
		fmt.Printf("SaveRezisedImage unexpected error: %s\n", err.Error())
	}

    updateImageDimensions := func(i image.Image) {
        newWidth := i.Bounds().Size().X
        newHeight := i.Bounds().Size().Y
        if isThumb {
            imageInfo.thumbWidth, imageInfo.thumbHeight = newWidth, newHeight
        } else {
            imageInfo.viewerWidth, imageInfo.viewerHeight = newWidth, newHeight
        }
    }

	fileExists,_ := exists(filename)

	if skipIfNewer && fileExists {
		imageInfo, err := os.Stat(imageInfo.sourcePath)
		if err != nil {
			printErr(err)
			return
		}
		existingInfo, err := os.Stat(filename)
		if err != nil {
			printErr(err)
			return
		}

		if imageInfo.ModTime().Before(existingInfo.ModTime()) {
			// Still need to read the actual width and height for building html
			existingRendition, err := imaging.Open(filename);

			if err != nil {
				printErr(err)
				return
			}

			updateImageDimensions(existingRendition)

			fmt.Printf("Skipping resizing for %s (target's last write time is newer than source)\n", path.Base(filename))
			return
		}
	}

	fmt.Printf("Generating %dx%d for %s\n", width, height, imageInfo.name)

	image, err := imaging.Open(imageInfo.sourcePath)

	if err != nil {
		printErr(err)
		return
	}

    image = imaging.Fit(image, width, height, imaging.Lanczos)

    updateImageDimensions(image)

    if fileExists {
        buf := new(bytes.Buffer)
        err = imaging.Encode(buf, image, imaging.JPEG)

        if err != nil {
            printErr(err)
            return
        }

        existingHash, err := hash_file_md5(filename)

        if err != nil {
            printErr(err)
            return
        }

        resizedHash := md5.Sum(buf.Bytes())

        if existingHash == resizedHash {
            fmt.Printf("Skipping rendition, generated image has same md5sum as what is already in target folder (%s)\n", filename)
            return
        }
    }

	err = imaging.Save(image, filename)

    if err != nil {
        printErr(err)
    }

    filesTouched++
}
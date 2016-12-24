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
	masterAlbum.UpdateImageRenditions(options.target)
	masterAlbum.UpdatePages(options.target, options.baseurl)

	fmt.Printf("%d files touched (not including contents of %s)", filesTouched, filepath.Join(options.target, "data"))

	return nil
}

func BuildGallery() {
    if exists, _ := exists(options.source); !exists {
        panic(errors.New(fmt.Sprintf("could not find path %s", options.source)))
    }

    masterAlbum = NewAlbum(options.name, options.source, nil)

    masterAlbum.LoadAlbum(options.source)
}

func CopyResources() {
    _ = os.Mkdir(options.target, filemode)
    RestoreAssets(options.target, "data")
}

func SaveResizedImage(image *Image, width, height int, filename string, skipIfNewer bool) {
	printErr := func(err error) {
		fmt.Printf("SaveRezisedImage unexpected error: %s\n", err.Error())
	}

	fileExists,_ := exists(filename)

	if skipIfNewer && fileExists {
		imageInfo, err := os.Stat(image.sourcePath)
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
			existingThumb, err := imaging.Open(filename);

			if err != nil {
				printErr(err)
				return
			}

			image.width = existingThumb.Bounds().Size().X
			image.height = existingThumb.Bounds().Size().Y

			fmt.Printf("Skipping resizing for %s (target's last write time is newer than source)\n", path.Base(filename))
			return
		}
	}

	fmt.Printf("Generating %dx%d for %s\n", width, height, image.name)

	img, err := imaging.Open(image.sourcePath)

	if err != nil {
		printErr(err)
		return
	}

	img = imaging.Fit(img, width, height, imaging.Lanczos)
	image.width = img.Bounds().Size().X
	image.height = img.Bounds().Size().Y

    if fileExists {
        buf := new(bytes.Buffer)
        err = imaging.Encode(buf, img, imaging.JPEG)

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

        if strings.Contains(filename, `\gallery1_build\thumbnail.jpg`) {
            fmt.Println("wot")
        }

        if existingHash == resizedHash {
            fmt.Printf("Skipping rendition, generated image has same md5sum as what is already in target folder (%s)\n", filename)
            return
        }
    }

	err = imaging.Save(img, filename)

    if err != nil {
        printErr(err)
    }

    filesTouched++
}
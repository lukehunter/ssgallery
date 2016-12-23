package main

import (
	"fmt"
	"os"
	"github.com/urfave/cli"
    //"github.com/disintegration/imaging"
    "errors"
    "path"
    "path/filepath"
    "strings"
    "net/http"
	"github.com/disintegration/imaging"
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

const thumbnail = "thumbnail.jpg"
const cacheFolder = "cache"
const filemode = 0644

var options Options
var gallery Gallery

func main() {
	app := cli.NewApp()
	app.Name = "ssgallery"
	app.Usage = "stupidly simple gallery"

	var args []string = []string{namearg, sourcearg, targetarg, baseurlarg, thumbwidtharg, thumbheightarg, viewerwidtharg, viewerheightarg}

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
	}

    i := Image { name: "image_test", path: "path" }

    fmt.Printf("%s", i.name)

	app.Action = func(c *cli.Context) error {
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

        BuildGallery()
        CopyResources()
        PopulateImageCache()
        CreatePages()

		return nil
	}

	app.Run(os.Args)
}

func BuildGallery() {
    if exists, _ := exists(options.source); !exists {
        panic(errors.New(fmt.Sprintf("could not find path %s", options.source)))
    }

    gallery = Gallery {
        name: options.name,
        albums: []Album {},
    }

    dirList := getDirList(options.source)

    for _,album := range dirList {
        // does base() work?
        curAlbum := Album { name: filepath.Base(album), folder: album, images: []Image{}}
        files := getFileList(album)

        for _,image := range files {
            if strings.Contains(image, thumbnail) {
                continue
            }

            curImage := Image { name: fileNameWithoutExtension(image), path: image }
            append(curAlbum.images, curImage)
        }

        append(gallery.albums, curAlbum)
    }
}

func CopyResources() {
    _ = os.Mkdir(options.target, filemode)
    RestoreAssets(options.target, "data")
}

func PopulateImageCache() {
	for _,album := range gallery.albums {
		os.Mkdir(path.Join(options.target, album.name, cacheFolder), filemode)

		fmt.Printf("Caching album: %s", album.name)

		for _,image := range album.images {
			Copy(image.path, path.Join(options.target, album.name, path.Base(image.path)))

			resize := func(width, height int) {
				path := path.Join(options.target, album.name, cacheFolder,
					formatFilename(image.name, width, height))
				SaveResizedImage(image, width, height, path, true)
			}

			// image thumbnail
			resize(options.thumbwidth, options.thumbheight)

			// image viewer
			resize(options.viewerwidth, options.viewerheight)
		}

		albumThumbnailPath := path.Join(album.folder, thumbnail)
		targetAlbumThumbnailPath := path.Join(options.target, album.name, thumbnail)

		// Always generate album thumbnail -- otherwise if thumbnail.jpg is removed from source it will never be re-generated
		if exists(albumThumbnailPath) {
			SaveResizedImage(Image{name: fmt.Sprintf("%s thumbnail", album.name), path: albumThumbnailPath},
				options.thumbwidth, options.thumbheight, targetAlbumThumbnailPath, false)
		} else {
			image := album.images[0]
			fmt.Printf("File %s not found, using %s as album thumbnail", albumThumbnailPath, image.path)
			SaveResizedImage(image, options.thumbwidth, options.thumbheight, targetAlbumThumbnailPath, false)
		}
	}
}

func SaveResizedImage(image *Image, width, height int, filename string, skipIfNewer bool) {
	printErr := func(err error) {
		fmt.Printf("SaveRezisedImage unexpected error: %s", err.Error())
	}

	if skipIfNewer && exists(filename) {
		imageInfo, err := os.Stat(image.path)
		if err != nil {
			printErr(err)
			return
		}
		existingInfo, err := os.Stat(filename)
		if err != nil {
			printErr(err)
			return
		}

		if imageInfo.ModTime() <= existingInfo.ModTime() {
			// Still need to read the actual width and height for building html
			existingThumb, err := imaging.Open(filename);

			if err != nil {
				printErr(err)
				return
			}

			image.width = existingThumb.Bounds().Size().X
			image.height = existingThumb.Bounds().Size().Y

			fmt.Printf("Skipping resizing for %s (target's last write time is newer than source", path.Base(filename))
			return
		}
	}

	fmt.Printf("Generating %dx%d for %s", width, height, image.name)

	img, err := imaging.Open(image.path)

	if err != nil {
		printErr(err)
		return
	}

	imaging.Fit(img, width, height, imaging.Lanczos)

	image.width = img.Bounds().Size().X
	image.height = img.Bounds().Size().Y

	imaging.Save(img, filename)
}




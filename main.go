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
)

const namearg string = "name"
const sourcearg string = "source"
const targetarg string = "target"
const baseurlarg string = "baseurl"
const disqusarg string = "disqus"
const thumbwidtharg string = "thumbwidth"
const thumbheightarg string = "thumbheight"
const viewerwidtharg string = "viewerwidth"
const viewerheightarg string = "viewerheight"
const testarg string = "test"

const thumbnail = "thumbnail.jpg"

var options Options
var gallery Gallery

func average(xs []float64) float64 {
	panic("Not Implemented")
    return 1.0;
}

func main() {
	app := cli.NewApp()
	app.Name = "ssgallery"
	app.Usage = "stupidly simple gallery"

	//var args []string = []string{namearg, sourcearg, targetarg, baseurlarg, thumbwidtharg, thumbheightarg, viewerwidtharg, viewerheightarg}

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
		//for _,arg := range args {
		//	if (!c.IsSet(arg)) {
		//		cli.ShowAppHelp(c)
		//		return cli.NewExitError(fmt.Sprintf("\n\nArgument '%s' is required", arg), 1)
		//	}
		//}

		fmt.Printf("\narguments:\nname: '%s'\nsource: '%s'\ntarget: '%s'\nbaseurl: '%s'\ndisqus: '%s'\nthumbwidth: %d\nthumbheight: %d\nviewerwidth: %d\nviewerheight: %d\n",
			options.name, options.source, options.target, options.baseurl, options.disqus, options.thumbwidth, options.thumbheight, options.viewerwidth, options.viewerheight)

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
    _ = os.Mkdir(options.target, 0644)
    RestoreAssets(options.target, "data")
}


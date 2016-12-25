package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"path/filepath"
	"strings"
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
const dataFolder = "data"
const filemode = 0755

type Options struct {
	name, source, target, baseurl, disqus              string
	thumbwidth, thumbheight, viewerwidth, viewerheight int
	skipextcheck, debug                                bool
}

var options Options
var masterAlbum *Album
var htmlFilesTouched, imageFilesTouched int

func main() {
	app := cli.NewApp()
	app.Name = "ssgallery"
	app.Usage = "stupidly simple gallery"

	options = Options{}

	app.Flags = []cli.Flag{
		cli.StringFlag{Name: namearg, Usage: "Set gallery name to `NAME`", Destination: &options.name},
		cli.StringFlag{Name: sourcearg, Usage: "Read albums from `FOLDER`", Destination: &options.source},
		cli.StringFlag{Name: targetarg, Usage: "Write web page to `FOLDER`", Destination: &options.target},
		cli.StringFlag{Name: baseurlarg, Usage: "Gallery will be hosted at http://server/`REL_URL`", Destination: &options.baseurl},
		cli.StringFlag{Name: disqusarg, Usage: "[Optional] Custom `URL` for disqus topic (go to disqus.com to get one)", Destination: &options.disqus},
		cli.IntFlag{Name: thumbwidtharg, Usage: "Set max thumbnail width to `WIDTH`", Destination: &options.thumbwidth},
		cli.IntFlag{Name: thumbheightarg, Usage: "Set max thumbnail height to `HEIGHT`", Destination: &options.thumbheight},
		cli.IntFlag{Name: viewerwidtharg, Usage: "Set max image viewer width to `WIDTH`", Destination: &options.viewerwidth},
		cli.IntFlag{Name: viewerheightarg, Usage: "Set max image viewer height to `HEIGHT`", Destination: &options.viewerheight},
		cli.BoolFlag{Name: skipextcheckarg, Usage: "Skip the check that limits to known supported file extensions (may include more images but may be slower if there are non-image files in the source folder)", Destination: &options.skipextcheck},
		cli.BoolFlag{Name: debugarg, Usage: "Enable debug logging", Destination: &options.debug},
	}

	app.Action = runApp

	app.Run(os.Args)
}

func runApp(c *cli.Context) error {
	var args []string = []string{namearg, sourcearg, targetarg, baseurlarg, thumbwidtharg,
		thumbheightarg, viewerwidtharg, viewerheightarg}

	// Check required args
	for _, arg := range args {
		if !c.IsSet(arg) {
			cli.ShowAppHelp(c)
			return cli.NewExitError(fmt.Sprintf("\n\nArgument '%s' is required", arg), 1)
		}
	}

	// Print back what the user entered
	fmt.Printf("\narguments:\nname: '%s'\nsource: '%s'\ntarget: '%s'\nbaseurl: '%s'\ndisqus: '%s'"+
		"\nthumbwidth: %d\nthumbheight: %d\nviewerwidth: %d\nviewerheight: %d\n",
		options.name, options.source, options.target, options.baseurl, options.disqus, options.thumbwidth,
		options.thumbheight, options.viewerwidth, options.viewerheight)

	if !strings.HasPrefix(options.baseurl, "/") {
		options.baseurl = fmt.Sprintf("/%s", options.baseurl)
		fmt.Printf("Warning: baseurl does not include leading slash. I added one for you. New baseurl = '%s'\n",
			options.baseurl)
	}

	if exists, _ := exists(options.source); !exists {
		panic(fmt.Errorf("could not find path %s", options.source))
	}

	_ = os.Mkdir(options.target, filemode)
	RestoreAssets(options.target, dataFolder)

	masterAlbum = NewAlbum(options.name, options.source, nil)
	masterAlbum.LoadAlbum(options.source)

	if !masterAlbum.HasImages() {
		return fmt.Errorf("No images found in %s or any of its subdirectories. Exiting.", masterAlbum.folder)
	}

	masterAlbum.UpdateImageRenditions(options.target)

	RenderAlbumHtml(masterAlbum, options.target, options.baseurl)

	fmt.Printf("%d files touched (%d html, %d images), not including contents of %s\n",
		htmlFilesTouched + imageFilesTouched, htmlFilesTouched, imageFilesTouched,
		filepath.Join(options.target, dataFolder))

	return nil
}

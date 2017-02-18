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
const verifyimagesarg = "verifyimages"

const thumbnail = "thumbnail.jpg"
const cacheFolder = "cache"
const dataFolder = "data"
const filemode = 0755

type Options struct {
	name, source, target, baseurl, disqus              string
	thumbwidth, thumbheight, viewerwidth, viewerheight int
	skipextcheck, debug, verifyimages                  bool
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
		// Required
		cli.StringFlag{Name: namearg, Usage: "Set gallery name to `NAME`", Destination: &options.name},
		cli.StringFlag{Name: sourcearg, Usage: "Read albums from `FOLDER`", Destination: &options.source},
		cli.StringFlag{Name: targetarg, Usage: "Write web page to `FOLDER`", Destination: &options.target},
		cli.StringFlag{Name: baseurlarg, Usage: "Gallery will be hosted at http://server/`REL_URL`", Destination: &options.baseurl},
		// Optional
		cli.StringFlag{Name: disqusarg, Usage: "[Optional] Custom `URL` for disqus topic (go to disqus.com to get one)", Destination: &options.disqus},
		cli.IntFlag{Name: thumbwidtharg, Usage: "[Optional] Set max thumbnail width to `WIDTH`", Destination: &options.thumbwidth},
		cli.IntFlag{Name: thumbheightarg, Usage: "[Optional] Set max thumbnail height to `HEIGHT`", Destination: &options.thumbheight},
		cli.IntFlag{Name: viewerwidtharg, Usage: "[Optional] Set max image viewer width to `WIDTH`", Destination: &options.viewerwidth},
		cli.IntFlag{Name: viewerheightarg, Usage: "[Optional] Set max image viewer height to `HEIGHT`", Destination: &options.viewerheight},
		cli.BoolFlag{Name: verifyimagesarg, Usage: "[Optional] Attempt to load each image first to ensure it is valid (may be dangerous in combination with --skipextcheck)", Destination: &options.verifyimages},
		cli.BoolFlag{Name: skipextcheckarg, Usage: "[Optional] Skip the check that limits to known supported file extensions (may include more images but may be slower if there are non-image files in the source folder)", Destination: &options.skipextcheck},
		cli.BoolFlag{Name: debugarg, Usage: "[Optional] Enable debug logging", Destination: &options.debug},
	}

	app.Action = runApp

	app.Run(os.Args)
}

func runApp(c *cli.Context) error {
	var requiredArgs []string = []string{namearg, sourcearg, targetarg, baseurlarg}

	// Check required args
	for _, arg := range requiredArgs {
		if !c.IsSet(arg) {
			cli.ShowAppHelp(c)
			return cli.NewExitError(fmt.Sprintf("\n\nArgument '%s' is required", arg), 1)
		}
	}

	// Defaults for optional args
	if !c.IsSet(disqusarg) {
		options.disqus = ""
	}
	if !c.IsSet(thumbwidtharg) {
		options.thumbwidth = 170
	}
	if !c.IsSet(thumbheightarg) {
		options.thumbheight = 130
	}
	if !c.IsSet(viewerwidtharg) {
		options.viewerwidth = 1500
	}
	if !c.IsSet(viewerheightarg) {
		options.viewerheight = 1000
	}
	if !c.IsSet(verifyimagesarg) {
		options.verifyimages = false
	}
	if !c.IsSet(skipextcheckarg) {
		options.skipextcheck = false
	}
	if !c.IsSet(debugarg) {
		options.debug = false
	}

	// Print back what will be used
	fmt.Printf("\narguments:\nname: '%s'\nsource: '%s'\ntarget: '%s'\nbaseurl: '%s'\ndisqus: '%s'"+
		"\nthumbwidth: %d\nthumbheight: %d\nviewerwidth: %d\nviewerheight: %d\n",
		options.name, options.source, options.target, options.baseurl, options.disqus, options.thumbwidth,
		options.thumbheight, options.viewerwidth, options.viewerheight)

	// Ensure base url has leading and trailing slash
	modified := false
	if !strings.HasPrefix(options.baseurl, "/") {
		options.baseurl = fmt.Sprintf("/%s", options.baseurl)
		modified = true
	}
	if !strings.HasSuffix(options.baseurl, "/") {
		options.baseurl = fmt.Sprintf("%s/", options.baseurl)
		modified = true
	}
	if modified {
		fmt.Printf("formatted baseurl: %s\n", options.baseurl)
	}

	if exists, _ := exists(options.source); !exists {
		panic(fmt.Errorf("could not find path %s", options.source))
	}

	_ = os.Mkdir(options.target, filemode)
	err := RestoreAssets(options.target, dataFolder)

	if err != nil {
		return fmt.Errorf("Unable to restore assets")
	}

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

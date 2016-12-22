package main

import (
	"fmt"
	"os"
	"github.com/urfave/cli"
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

func main() {
	app := cli.NewApp()
	app.Name = "gogallery"
	app.Usage = "use it good"

	var args []string = []string{namearg, sourcearg, targetarg, baseurlarg, thumbwidtharg, thumbheightarg, viewerwidtharg, viewerheightarg}

	var name string
	var source string
	var target string
	var baseurl string
	var disqus string
	var thumbwidth int
	var thumbheight int
	var viewerwidth int
	var viewerheight int

	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: namearg,
			Usage: "Set gallery name to `NAME`",
			Destination: &name,
		},
		cli.StringFlag{
			Name: sourcearg,
			Usage: "Read albums from `FOLDER`",
			Destination: &source,
		},
		cli.StringFlag{
			Name: targetarg,
			Usage: "Write web page to `FOLDER`",
			Destination: &target,
		},
		cli.StringFlag{
			Name: baseurlarg,
			Usage: "Gallery will be hosted at http://server/`REL_URL`",
			Destination: &baseurl,
		},
		cli.StringFlag{
			Name: disqusarg,
			Usage: "[Optional] Custom `URL` for disqus topic (go to disqus.com to get one)",
			Destination: &disqus,
		},
		cli.IntFlag{
			Name: thumbwidtharg,
			Usage: "Set max thumbnail width to `WIDTH`",
			Destination: &thumbwidth,
		},
		cli.IntFlag{
			Name: thumbheightarg,
			Usage: "Set max thumbnail height to `HEIGHT`",
			Destination: &thumbheight,
		},
		cli.IntFlag{
			Name: viewerwidtharg,
			Usage: "Set max image viewer width to `WIDTH`",
			Destination: &viewerwidth,
		},
		cli.IntFlag{
			Name: viewerheightarg,
			Usage: "Set max image viewer height to `HEIGHT`",
			Destination: &viewerheight,
		},
	}

	//sort.Sort(cli.FlagsByName(app.Flags))

	app.Action = func(c *cli.Context) error {
		for _,arg := range args {
			if (!c.IsSet(arg)) {
				cli.ShowAppHelp(c)
				return cli.NewExitError(fmt.Sprintf("\n\narg %s is required", arg), 1)
			}
		}

		fmt.Printf("\narguments:\nname: '%s'\nsource: '%s'\ntarget: '%s'\nbaseurl: '%s'\ndisqus: '%s'\nthumbwidth: %d\nthumbheight: %d\nviewerwidth: %d\nviewerheight: %d\n",
			name, source, target, baseurl, disqus, thumbwidth, thumbheight, viewerwidth, viewerheight)
		return nil
	}

	app.Run(os.Args)
}

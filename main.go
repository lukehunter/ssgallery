package main

import (
	"fmt"
	"os"
	"github.com/urfave/cli"
    "github.com/disintegration/imaging"
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

func average(xs []float64) float64 {
	panic("Not Implemented")
}

func main() {
	app := cli.NewApp()
	app.Name = "ssgallery"
	app.Usage = "stupidly simple gallery"

	//var args []string = []string{namearg, sourcearg, targetarg, baseurlarg, thumbwidtharg, thumbheightarg, viewerwidtharg, viewerheightarg}

	var name string
	var source string
	var target string
	var baseurl string
	var disqus string
	var thumbwidth int
	var thumbheight int
	var viewerwidth int
	var viewerheight int
    var test string

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
        cli.StringFlag{
            Name: testarg,
            Usage: "Testing",
            Destination: &test,
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
			name, source, target, baseurl, disqus, thumbwidth, thumbheight, viewerwidth, viewerheight)

        image, err := imaging.Open(test);

        if err != nil {
            panic(err)
        }

        fmt.Printf("%d x %d\n", image.Bounds().Size().X, image.Bounds().Size().Y)

        resized := imaging.Fit(image, thumbwidth, thumbheight, imaging.Bartlett)

        imaging.Save(resized, fmt.Sprintf("%s-resized.jpg", test));

		return nil
	}

	app.Run(os.Args)
}

package main

import (
	"fmt"
	"os"
	"github.com/urfave/cli"
    //"github.com/disintegration/imaging"
    "errors"
    "path/filepath"
    "strings"
	"github.com/disintegration/imaging"
	"strconv"
    "path"
    "crypto/md5"
    "encoding/hex"
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
const filemode = 0644

var options Options
var gallery Gallery
var filesTouched int

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
        cli.BoolFlag{
            Name: skipextcheckarg,
            Usage: "Skip the check that limits to known supported file extensions (may include more images but may be slower if there are non-image files in the source folder)",
            Destination: &options.skipextcheck,
        },
	}

    i := Image { name: "image_test", path: "path" }

    fmt.Printf("%s\n", i.name)

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

        if !strings.HasPrefix(options.baseurl, "/") {
            options.baseurl = fmt.Sprintf("/%s", options.baseurl)
            fmt.Printf("Warning: baseurl does not include leading slash. I added one for you. New baseurl = '%s'\n", options.baseurl)
        }

        BuildGallery()
        CopyResources()
        PopulateImageCache()
        CreatePages()

        fmt.Printf("%d files touched (not including contents of %s)", filesTouched, filepath.Join(options.target, "data"))

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

    fmt.Printf("Scanning %s...\n", options.source)
    dirList := getDirList(options.source)

    fmt.Printf("Found %d albums", len(dirList))
    for _,album := range dirList {
        fmt.Printf("Scanning %s...\n", album)
        curAlbum := Album { name: filepath.Base(album), folder: album, images: []Image{}}
        files := getFileList(album)

        fmt.Printf("Found %d images\n", len(files))
        for _,image := range files {
            if strings.Contains(image, thumbnail) {
                continue
            }

            if !options.skipextcheck {
                valid := []string{".jpg", ".jpeg", ".png", ".gif", ".tiff", ".bmp", ".tif"}

                ext := strings.ToLower(filepath.Ext(image))

                if !stringInSlice(ext, valid) {
                    fmt.Printf("Unrecognized file extension on %s, ignoring (accepted are %s, override this behavior with --%s)", image, strings.Join(valid, ", "), skipextcheckarg)
                    continue
                }
            }

            _, err := imaging.Open(image)
            if err != nil {
                fmt.Printf("Unable to open %s, skipping (%s)\n", image, err.Error())
                continue
            }

            curImage := Image { name: fileNameWithoutExtension(image), path: image }
			curAlbum.AddImage(curImage)
        }

		gallery.AddAlbum(curAlbum)
    }
}

func CopyResources() {
    _ = os.Mkdir(options.target, filemode)
    RestoreAssets(options.target, "data")
}

func PopulateImageCache() {
	for _,album := range gallery.albums {
        curCacheFolder := filepath.Join(options.target, album.name, cacheFolder)
		err := os.MkdirAll(curCacheFolder, filemode)

        if err != nil {
            fmt.Printf("Error making directory %s: %s", curCacheFolder, err.Error())
            continue
        }

		fmt.Printf("Caching album: %s\n", album.name)

		for i := 0; i < len(album.images); i++ {
			image := &album.images[i] // necessary because we are modifying the image class (hack)
			target := filepath.Join(options.target, album.name, filepath.Base(image.path))
			err := Copy(image.path, target)

			if err != nil {
				fmt.Printf("PopulateImageCache error copying %s to %s: %s\n", image.path, target, err.Error())
			}

			resize := func(width, height int) {
				path := filepath.Join(options.target, album.name, cacheFolder,
					formatFilename(image.name, width, height))
				SaveResizedImage(image, width, height, path, true)
			}

			// image thumbnail
			resize(options.thumbwidth, options.thumbheight)

			// image viewer
			resize(options.viewerwidth, options.viewerheight)
		}

		albumThumbnailPath := filepath.Join(album.folder, thumbnail)
		targetAlbumThumbnailPath := filepath.Join(options.target, album.name, thumbnail)

		// Always generate album thumbnail -- otherwise if thumbnail.jpg is removed from source it will never be re-generated
		if exists,_ := exists(albumThumbnailPath); exists {
			SaveResizedImage(&Image{name: fmt.Sprintf("%s thumbnail", album.name), path: albumThumbnailPath},
				options.thumbwidth, options.thumbheight, targetAlbumThumbnailPath, false)
		} else {
			image := album.images[0]
			fmt.Printf("File %s not found, using %s as album thumbnail\n", albumThumbnailPath, image.path)
			SaveResizedImage(&image, options.thumbwidth, options.thumbheight, targetAlbumThumbnailPath, false)
		}
	}
}

func SaveResizedImage(image *Image, width, height int, filename string, skipIfNewer bool) {
	printErr := func(err error) {
		fmt.Printf("SaveRezisedImage unexpected error: %s\n", err.Error())
	}

	fileExists,_ := exists(filename)

	if skipIfNewer && fileExists {
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

	img, err := imaging.Open(image.path)

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

        if existingHash == resizedHash {
            fmt.Printf("Skipping album thumbnail, generated thumbnail has same md5sum as %s\n", filename)
            return
        }
    }

	err = imaging.Save(img, filename)

    if err != nil {
        printErr(err)
    }

    filesTouched++
}

func CreatePages() {
	galleryTemplate := NewTemplate(galleryTemplateRaw)

	galleryValues := map[string]string {
		"SSG_GALLERY_NAME": gallery.name,
		"SSG_HOME_URL": options.baseurl,
		"SSG_GALLERY_URL": options.baseurl + gallery.name + "/",
		"SSG_DISQUS_URL": options.disqus,
	}

	galleryTemplate.AddValues(galleryValues);

	for _,album := range gallery.albums {
		albumThumb := filepath.Join(options.target, album.name, thumbnail)
		albumThumbImg, err := imaging.Open(albumThumb)

		if err != nil {
			fmt.Printf("Unable to open %s\n", albumThumb)
			continue
		}

		albumValues := map[string]string {
			"SSG_ALBUM_NAME": album.name,
			"SSG_ALBUM_URL": filepath.Join(options.baseurl, album.name),
			"SSG_ALBUM_THUMBNAIL_WIDTH": strconv.Itoa(albumThumbImg.Bounds().Size().X),
			"SSG_ALBUM_THUMBNAIL_HEIGHT": strconv.Itoa(albumThumbImg.Bounds().Size().Y),
		}

		albumTemplate := NewTemplate(albumTemplateRaw)

		albumTemplate.AddValues(galleryValues)
		albumTemplate.AddValues(albumValues)

		galleryTemplate.AddItem(TemplateItem{tag: "SSG_ALBUM_LIST_ITEM", values: albumValues})

		for i,image := range album.images {
			imageTemplate := NewTemplate(imageTemplateRaw)

			var nextPage, prevPage, picToPreload string

			if i > 0 {
				prevPage = fmt.Sprintf("%s.html", album.images[i - 1].name)
				imageTemplate.SetHiddenRegion("SSG_PREV_IMAGE_LINK", false)
			} else {
				imageTemplate.SetHiddenRegion("SSG_PREV_IMAGE_LINK", true)
			}
			
			if (i < len(album.images) - 1) {
				nextImage := album.images[i + 1]
				nextPage = fmt.Sprintf("%s.html", nextImage.name)
				picToPreload = filepath.Join(cacheFolder, formatFilename(nextImage.name, options.viewerwidth, options.viewerheight))
				
				imageTemplate.SetHiddenRegion("SSG_NEXT_IMAGE_LINK", false)
			} else {
				imageTemplate.SetHiddenRegion("SSG_NEXT_IMAGE_LINK", true)
			}

            if len(options.disqus) == 0 {
                imageTemplate.SetHiddenRegion("SSG_DISQUS", true)
            } else {
                imageTemplate.SetHiddenRegion("SSG_DISQUS", false)
            }

			imageThumb := filepath.Join(options.target, album.name, cacheFolder,
				formatFilename(image.name, options.thumbwidth, options.thumbheight))
			imageThumbImage, err := imaging.Open(imageThumb)

			if err != nil {
				fmt.Printf("Unable to open %s\n", imageThumb)
				continue
			}

			imageUrl := filepath.Join(cacheFolder, formatFilename(image.name, options.viewerwidth, options.viewerheight))
			imageThumbUrl := filepath.Join(cacheFolder, formatFilename(image.name, options.thumbwidth, options.thumbheight))
            pageUrl := fmt.Sprintf("%s.html", image.name)
            imageIdBytes, err := hash_file_md5(image.path)

            if err != nil {
                fmt.Printf("Failed to calculate hash on %s (%s), falling back on gallery/album/image style ID for disqus", image.path, err.Error())
                imageIdBytes = md5.Sum([]byte(fmt.Sprintf("%s-%s-%s", gallery.name, album.name, image.name)))
            }

            imageId := hex.EncodeToString(imageIdBytes[:])

			imageValues := map[string]string {
				"SSG_IMAGE_NAME": image.name,
				"SSG_PREV_IMAGE_PAGE_URL": prevPage,
				"SSG_NEXT_IMAGE_PAGE_URL": nextPage,
				"SSG_PRELOAD_URL": picToPreload,
				"SSG_IMAGE_WIDTH": strconv.Itoa(image.width),
				"SSG_IMAGE_HEIGHT": strconv.Itoa(image.height),
				"SSG_IMAGE_URL": imageUrl,
				"SSG_IMAGE_PAGE_URL": pageUrl,
				"SSG_IMAGE_DISQUS_ID": imageId,
				"SSG_IMAGE_THUMBNAIL_URL": imageThumbUrl,
				"SSG_ORIG_IMAGE_URL": image.filename(),
				"SSG_IMAGE_THUMBNAIL_WIDTH": strconv.Itoa(imageThumbImage.Bounds().Size().X),
				"SSG_IMAGE_THUMBNAIL_HEIGHT": strconv.Itoa(imageThumbImage.Bounds().Size().Y),
			}

			imageTemplate.AddValues(galleryValues)
			imageTemplate.AddValues(albumValues)
			imageTemplate.AddValues(imageValues)

			albumTemplate.AddItem(TemplateItem{ tag: "SSG_IMAGE_LIST_ITEM", values: imageValues})

			imageTemplate.RenderHtml(filepath.Join(options.target, album.name, fmt.Sprintf("%s.html", image.name)))
		}

		albumTemplate.RenderHtml(filepath.Join(options.target, album.name, "index.html"))
	}

	galleryTemplate.RenderHtml(filepath.Join(options.target, "index.html"))
}




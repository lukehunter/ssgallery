package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/disintegration/imaging"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

type Options struct {
	name, source, target, baseurl, disqus              string
	thumbwidth, thumbheight, viewerwidth, viewerheight int
	skipextcheck, debug                                bool
}

type Image struct {
	name, sourcePath, targetPath                       string
	thumbWidth, thumbHeight, viewerWidth, viewerHeight int
}

func (i *Image) filename() string {
	return filepath.Base(i.sourcePath)
}

type Album struct {
	name, folder, relUrl string
	parent               *Album
	albums               []Album
	images               []Image
}

func NewAlbum(name, folder string, parent *Album) *Album {
	a := new(Album)
	a.name = name
	a.folder = folder
	a.parent = parent
	a.albums = []Album{}
	a.images = []Image{}
	if a.parent != nil {
		a.relUrl = path.Join(a.parent.relUrl, name)
	}
	return a
}

func NewTopAlbum(name, folder, relUrl string, parent *Album) *Album {
	a := NewAlbum(name, folder, parent)

	a.relUrl = relUrl

	return a
}

func (a *Album) AddImage(image Image) {
	a.images = append(a.images[:], image)
}

func (a *Album) AddAlbum(album Album) {
	a.albums = append(a.albums[:], album)
}

func (a *Album) HasImages() bool {
	if a == nil {
		return false
	}

	if len(a.images) > 0 {
		return true
	}

	for _, album := range a.albums {
		if album.HasImages() {
			return true
		}
	}

	return false
}

func (a *Album) FindThumbnail() *Image {
	if len(a.images) > 0 {
		return &a.images[0]
	} else {
		for _, a := range a.albums {
			result := a.FindThumbnail()

			if result != nil {
				return result
			}
		}
	}

	return nil
}

func (a *Album) GetBreadcrumbs(albums []Album) []Album {
	albums = append(albums, *a)

	if a.parent == nil {
		// reverse it so 0 is root and n is leaf
		return reverse(albums)
	}

	return a.parent.GetBreadcrumbs(albums)
}

func reverse(albums []Album) []Album {
	for i := 0; i < len(albums)/2; i++ {
		j := len(albums) - i - 1
		albums[i], albums[j] = albums[j], albums[i]
	}
	return albums
}

func (a *Album) GetBreadcrumbPath(rest string) string {
	rest = fmt.Sprintf("%s/%s", a.name, rest)

	if a.parent == nil {
		return rest
	}

	return a.parent.GetBreadcrumbPath(rest)
}

func (a *Album) LoadAlbum(path string) {
	fmt.Printf("Scanning %s for albums...\n", path)
	dirList := getDirList(path)
	fmt.Printf("Found %d albums\n", len(dirList))

	for _, album := range dirList {
		curAlbum := NewAlbum(filepath.Base(album), album, a)

		curAlbum.LoadAlbum(album)

		fmt.Printf("Scanning %s for images...\n", album)
		files := getFileList(album)
		fmt.Printf("Found %d images\n", len(files))

		for _, image := range files {
			if strings.Contains(image, thumbnail) {
				continue
			}

			if !options.skipextcheck {
				valid := []string{".jpg", ".jpeg", ".png", ".gif", ".tiff", ".bmp", ".tif"}

				ext := strings.ToLower(filepath.Ext(image))

				if !stringInSlice(ext, valid) {
					fmt.Printf("Unrecognized file extension on %s, ignoring (accepted are %s, override this behavior with --%s)\n", image, strings.Join(valid, ", "), skipextcheckarg)
					continue
				}
			}

			_, err := imaging.Open(image)
			if err != nil {
				fmt.Printf("Unable to open %s, skipping (%s)\n", image, err.Error())
				continue
			}

			curImage := Image{name: fileNameWithoutExtension(image), sourcePath: image}
			curAlbum.AddImage(curImage)
		}

		a.AddAlbum(*curAlbum)
	}
}

func (a *Album) UpdateImageRenditions(targetPath string) {
	if !a.HasImages() {
		fmt.Printf("Album %s has no images, skipping resize step")
		return
	}

	fmt.Printf("Updating renditions for album: %s\n", a.name)

	for _, album := range a.albums {
		album.UpdateImageRenditions(filepath.Join(targetPath, album.name))
	}

	curCacheFolder := filepath.Join(targetPath, cacheFolder)
	err := os.MkdirAll(curCacheFolder, filemode)

	if err != nil {
		fmt.Printf("Error making target directory %s: %s\n", curCacheFolder, err.Error())
		return
	}

	for i := 0; i < len(a.images); i++ {
		image := &a.images[i] // necessary because we are modifying the image class (hack)
		target := filepath.Join(targetPath, filepath.Base(image.sourcePath))
		err := Copy(image.sourcePath, target)

		if err != nil {
			fmt.Printf("UpdateImageRenditions error copying %s to %s: %s\n", image.sourcePath, target, err.Error())
		}

		resize := func(width, height int, isThumb bool) {
			path := filepath.Join(targetPath, cacheFolder,
				formatFilename(image.name, width, height))
			SaveResizedImage(image, width, height, path, isThumb, true)
		}

		// image thumbnail
		resize(options.thumbwidth, options.thumbheight, true)

		// image viewer
		resize(options.viewerwidth, options.viewerheight, false)
	}

	albumThumbnailPath := filepath.Join(a.folder, thumbnail)
	targetAlbumThumbnailPath := filepath.Join(targetPath, thumbnail)

	// Always generate album thumbnail -- otherwise if thumbnail.jpg is removed from source it will never be re-generated
	if exists, _ := exists(albumThumbnailPath); exists {
		SaveResizedImage(&Image{name: fmt.Sprintf("%s thumbnail", a.name), sourcePath: albumThumbnailPath},
			options.thumbwidth, options.thumbheight, targetAlbumThumbnailPath, true, false)
	} else {
		image := a.FindThumbnail()
		fmt.Printf("File %s not found, using %s as album thumbnail\n", albumThumbnailPath, image.sourcePath)
		SaveResizedImage(image, options.thumbwidth, options.thumbheight, targetAlbumThumbnailPath, true, false)
	}
}

func (a *Album) UpdatePages(targetPhysicalPath, relativeUrl string) {
	albumTemplate := NewTemplate(albumTemplateRaw)

	albumUrl := relativeUrl

	albumValues := map[string]string{
		"SSG_ALBUM_NAME": a.name,
		"SSG_HOME_URL":   options.baseurl,
		"SSG_ALBUM_URL":  albumUrl,
		"SSG_DISQUS_URL": options.disqus,
		"SSG_CSS_URL":    path.Join(options.baseurl, dataFolder, "ssgallery.css"),
	}

	albumTemplate.AddValues(albumValues)

	if a.parent != nil {
		breadcrumbs := a.parent.GetBreadcrumbs([]Album{})

		for _, breadcrumb := range breadcrumbs {
			breadcrumbTemplateItem := NewTemplateItem("SSG_BREADCRUMB_LIST_ITEM")

			(*breadcrumbTemplateItem.values)["SSG_ALBUM_URL"] = breadcrumb.relUrl
			(*breadcrumbTemplateItem.values)["SSG_ALBUM_NAME"] = breadcrumb.name

			albumTemplate.AddItem(*breadcrumbTemplateItem)
		}
	}

	for _, subAlbum := range a.albums {
		subAlbum.UpdatePages(filepath.Join(targetPhysicalPath, subAlbum.name),
			path.Join(relativeUrl, subAlbum.name))

		albumThumb := filepath.Join(targetPhysicalPath, subAlbum.name, thumbnail)
		albumThumbImg, err := imaging.Open(albumThumb)

		if err != nil {
			fmt.Printf("Unable to open %s\n", albumThumb)
			continue
		}

		subAlbumValues := map[string]string{
			"SSG_ALBUM_NAME":             subAlbum.name,
			"SSG_ALBUM_URL":              filepath.Join(albumUrl, subAlbum.name),
			"SSG_ALBUM_THUMBNAIL_WIDTH":  strconv.Itoa(albumThumbImg.Bounds().Size().X),
			"SSG_ALBUM_THUMBNAIL_HEIGHT": strconv.Itoa(albumThumbImg.Bounds().Size().Y),
		}

		subAlbumTemplateItem := NewTemplateItem("SSG_ALBUM_LIST_ITEM")
		subAlbumTemplateItem.values = &subAlbumValues

		albumTemplate.AddItem(*subAlbumTemplateItem)
	}

	for i, image := range a.images {
		imageTemplate := NewTemplate(imageTemplateRaw)
		imageTemplateItem := NewTemplateItem("SSG_IMAGE_LIST_ITEM")

		breadcrumbs := a.GetBreadcrumbs([]Album{})

		for _, breadcrumb := range breadcrumbs {
			breadcrumbTemplateItem := NewTemplateItem("SSG_BREADCRUMB_LIST_ITEM")

			(*breadcrumbTemplateItem.values)["SSG_ALBUM_URL"] = breadcrumb.relUrl
			(*breadcrumbTemplateItem.values)["SSG_ALBUM_NAME"] = breadcrumb.name

			imageTemplate.AddItem(*breadcrumbTemplateItem)
		}

		var nextPage, prevPage, picToPreload string

		if i > 0 {
			prevPage = fmt.Sprintf("%s.html", a.images[i-1].name)
			imageTemplate.SetHiddenRegion("SSG_PREV_IMAGE_LINK", false)
		} else {
			imageTemplate.SetHiddenRegion("SSG_PREV_IMAGE_LINK", true)
		}

		if i < len(a.images)-1 {
			nextImage := a.images[i+1]
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

		imageThumb := filepath.Join(targetPhysicalPath, cacheFolder,
			formatFilename(image.name, options.thumbwidth, options.thumbheight))
		imageThumbImage, err := imaging.Open(imageThumb)

		if err != nil {
			fmt.Printf("Unable to open %s\n", imageThumb)
			continue
		}

		imageUrl := filepath.Join(cacheFolder, formatFilename(image.name, options.viewerwidth, options.viewerheight))
		imageThumbUrl := filepath.Join(cacheFolder, formatFilename(image.name, options.thumbwidth, options.thumbheight))
		pageUrl := fmt.Sprintf("%s.html", image.name)
		imageIdBytes, err := hash_file_md5(image.sourcePath)

		if err != nil {
			fmt.Printf("Failed to calculate hash on %s (%s), falling back on gallery/album/image style ID for disqus", image.sourcePath, err.Error())
			imageId := fmt.Sprintf("%s-%s-%s", a.GetBreadcrumbPath(""), a.name, image.name)
			imageIdBytes = md5.Sum([]byte(imageId))
		}

		imageId := hex.EncodeToString(imageIdBytes[:])

		imageValues := map[string]string{
			"SSG_HOME_URL":               options.baseurl,
			"SSG_IMAGE_NAME":             image.name,
			"SSG_PREV_IMAGE_PAGE_URL":    prevPage,
			"SSG_NEXT_IMAGE_PAGE_URL":    nextPage,
			"SSG_PRELOAD_URL":            picToPreload,
			"SSG_IMAGE_WIDTH":            strconv.Itoa(image.viewerWidth),
			"SSG_IMAGE_HEIGHT":           strconv.Itoa(image.viewerHeight),
			"SSG_IMAGE_URL":              imageUrl,
			"SSG_IMAGE_PAGE_URL":         pageUrl,
			"SSG_IMAGE_DISQUS_ID":        imageId,
			"SSG_IMAGE_THUMBNAIL_URL":    imageThumbUrl,
			"SSG_ORIG_IMAGE_URL":         image.filename(),
			"SSG_IMAGE_THUMBNAIL_WIDTH":  strconv.Itoa(imageThumbImage.Bounds().Size().X),
			"SSG_IMAGE_THUMBNAIL_HEIGHT": strconv.Itoa(imageThumbImage.Bounds().Size().Y),
		}

		imageTemplate.AddValues(albumValues)
		imageTemplate.AddValues(imageValues)
		imageTemplateItem.AddValues(imageValues)

		albumTemplate.AddItem(*imageTemplateItem)

		imageTemplate.RenderHtml(filepath.Join(targetPhysicalPath, fmt.Sprintf("%s.html", image.name)))
	}

	targetPath := filepath.Join(targetPhysicalPath, "index.html")
	if strings.HasSuffix(targetPath, "wallpapers/index.html") {
		fmt.Println()
	}

	albumTemplate.RenderHtml(targetPath)
}

package main

import (
	"strconv"
	"path/filepath"
	"fmt"
	"path"
	"github.com/disintegration/imaging"
)

func RenderAlbumHtml(a *Album, targetPhysicalPath, relativeUrl string) {
	if !a.HasImages() {
		fmt.Printf("Skipping empty album '%s'\n", a.name)
		return
	}

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
		RenderBreadcrumbHtml(a.parent, albumTemplate)
	}

	for _, subAlbum := range a.albums {
		RenderAlbumThumbHtml(&subAlbum, albumTemplate, targetPhysicalPath, relativeUrl, albumUrl)
	}

	for i, image := range a.images {
		RenderImageHtml(a, albumValues, albumTemplate, &image, i, targetPhysicalPath)
	}

	targetPath := filepath.Join(targetPhysicalPath, "index.html")

	albumTemplate.RenderHtml(targetPath)
}

func RenderAlbumThumbHtml(subAlbum *Album, albumTemplate *Template, targetPhysicalPath, relativeUrl, albumUrl string) {
	albumPath := filepath.Join(targetPhysicalPath, subAlbum.name)
	albumRelUrl := path.Join(relativeUrl, subAlbum.name)

	RenderAlbumHtml(subAlbum, albumPath, albumRelUrl)

	albumThumb := filepath.Join(targetPhysicalPath, subAlbum.name, thumbnail)
	albumThumbImg, err := imaging.Open(albumThumb)

	if err != nil {
		fmt.Printf("Unable to open %s\n", albumThumb)
		return
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

func RenderImageHtml(a *Album, albumValues map[string]string, albumTemplate *Template, image *Image, i int, targetPhysicalPath string ) {
	imageTemplate := NewTemplate(imageTemplateRaw)
	imageTemplateItem := NewTemplateItem("SSG_IMAGE_LIST_ITEM")

	RenderBreadcrumbHtml(a, imageTemplate)

	nextPage, prevPage, picToPreload := SetNextPrev(i, a, imageTemplate)

	thumbWidth, thumbHeight := image.GetThumbSize(targetPhysicalPath)

	imageValues := map[string]string{
		"SSG_HOME_URL":               options.baseurl,
		"SSG_IMAGE_NAME":             image.name,
		"SSG_PREV_IMAGE_PAGE_URL":    prevPage,
		"SSG_NEXT_IMAGE_PAGE_URL":    nextPage,
		"SSG_PRELOAD_URL":            picToPreload,
		"SSG_IMAGE_WIDTH":            strconv.Itoa(image.viewerWidth),
		"SSG_IMAGE_HEIGHT":           strconv.Itoa(image.viewerHeight),
		"SSG_IMAGE_URL":              image.imageUrl(),
		"SSG_IMAGE_PAGE_URL":         image.pageUrl(),
		"SSG_IMAGE_DISQUS_ID":        image.GetDisqusId(a),
		"SSG_IMAGE_THUMBNAIL_URL":    image.thumbUrl(),
		"SSG_ORIG_IMAGE_URL":         image.filename(),
		"SSG_IMAGE_THUMBNAIL_WIDTH":  strconv.Itoa(thumbWidth),
		"SSG_IMAGE_THUMBNAIL_HEIGHT": strconv.Itoa(thumbHeight),
	}

	imageTemplate.SetHiddenRegion("SSG_DISQUS", len(options.disqus) == 0)

	imageTemplate.AddValues(albumValues)
	imageTemplate.AddValues(imageValues)
	imageTemplateItem.AddValues(imageValues)

	albumTemplate.AddItem(*imageTemplateItem)

	imageTemplate.RenderHtml(filepath.Join(targetPhysicalPath, fmt.Sprintf("%s.html", image.name)))
}

func RenderBreadcrumbHtml(a *Album, template *Template) {
	breadcrumbs := a.GetAlbumPath([]Album{})

	breadcrumbUrl := options.baseurl

	for _, breadcrumb := range breadcrumbs {
		breadcrumbTemplateItem := NewTemplateItem("SSG_BREADCRUMB_LIST_ITEM")

		// don't include the root album, its relurl is options.baseurl
		if breadcrumb.parent != nil {
			breadcrumbUrl = path.Join(breadcrumbUrl, breadcrumb.name)
		}

		(*breadcrumbTemplateItem.values)["SSG_ALBUM_URL"] = breadcrumbUrl
		(*breadcrumbTemplateItem.values)["SSG_ALBUM_NAME"] = breadcrumb.name

		template.AddItem(*breadcrumbTemplateItem)
	}
}

func SetNextPrev(i int, a *Album, imageTemplate *Template) (nextPage, prevPage, picToPreload string) {
	anyPrev := i > 0
	anyNext := i < len(a.images) - 1

	imageTemplate.SetHiddenRegion("SSG_PREV_IMAGE_LINK", !anyPrev)
	imageTemplate.SetHiddenRegion("SSG_NEXT_IMAGE_LINK", !anyNext)

	if anyPrev {
		prevPage = fmt.Sprintf("%s.html", a.images[i - 1].name)
	}

	if anyNext {
		nextImage := a.images[i + 1]
		nextPage = fmt.Sprintf("%s.html", nextImage.name)
		picToPreload = filepath.Join(cacheFolder, formatFilename(nextImage.name, options.viewerwidth, options.viewerheight))
	}

	return nextPage, prevPage, picToPreload
}

package main

import (
	"path/filepath"
	"github.com/disintegration/imaging"
	"fmt"
	"crypto/md5"
	"encoding/hex"
)

type Image struct {
	name, sourcePath, targetPath                       string
	thumbWidth, thumbHeight, viewerWidth, viewerHeight int
}

func (i *Image) filename() string {
	return filepath.Base(i.sourcePath)
}

func ReadImageMetadata(image *Image, a *Album, targetPhysicalPath string) (imageUrl, pageUrl, imageId, imageThumbUrl string, thumbWidth, thumbHeight int) {
	imageThumb := filepath.Join(targetPhysicalPath, cacheFolder,
		formatFilename(image.name, options.thumbwidth, options.thumbheight))
	imageThumbImage, err := imaging.Open(imageThumb)

	if err != nil {
		fmt.Printf("Unable to open %s\n", imageThumb)
		return
	}

	imageUrl = filepath.Join(cacheFolder, formatFilename(image.name, options.viewerwidth, options.viewerheight))
	imageThumbUrl = filepath.Join(cacheFolder, formatFilename(image.name, options.thumbwidth, options.thumbheight))
	pageUrl = fmt.Sprintf("%s.html", image.name)
	imageIdBytes, err := hash_file_md5(image.sourcePath)

	if err != nil {
		fmt.Printf("Failed to calculate hash on %s (%s), falling back on gallery/album/image style ID for disqus", image.sourcePath, err.Error())
		imageId := fmt.Sprintf("%s-%s-%s", a.GetBreadcrumbPath(""), a.name, image.name)
		imageIdBytes = md5.Sum([]byte(imageId))
	}

	imageId = hex.EncodeToString(imageIdBytes[:])

	thumbWidth = imageThumbImage.Bounds().Size().X
	thumbHeight = imageThumbImage.Bounds().Size().Y

	return imageUrl, pageUrl, imageId, imageThumbUrl, thumbWidth, thumbHeight
}
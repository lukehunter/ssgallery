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

// Only the portion after the albums
func (i *Image) pageUrl() string {
	return fmt.Sprintf("%s.html", i.name)
}

// Relative to album
func (i *Image) imageUrl() string {
	return filepath.Join(cacheFolder, formatFilename(i.name, options.viewerwidth, options.viewerheight))
}

// Relative to album
func (i *Image) thumbUrl() string {
	return filepath.Join(cacheFolder, formatFilename(i.name, options.thumbwidth, options.thumbheight))
}

func (i *Image) GetDisqusId(album *Album) string {
	imageIdBytes, err := hash_file_md5(i.sourcePath)

	if err != nil {
		// There is some choice here whether to use only the image name, or the entire path. In one direction
		// there is more chance of collision with another image. In the other is more chance of orphaning comments
		// if the file is moved. The compromoise is to only use the album and image name assuming it will stay in the
		// direct containing folder.
		imageId := fmt.Sprintf("%s-%s", album.name, i.name)
		fmt.Printf("Failed to calculate hash on %s (%s), using md5 of %s for disqus ID", i.sourcePath, err.Error(), imageId)
		imageIdBytes = md5.Sum([]byte(imageId))
	}

	return hex.EncodeToString(imageIdBytes[:])
}

// slow. reads in the entire image.
func (i *Image) GetThumbSize(targetPhysicalPath string) (thumbWidth, thumbHeight int) {
	imageThumb := filepath.Join(targetPhysicalPath, cacheFolder,
		formatFilename(i.name, options.thumbwidth, options.thumbheight))
	imageThumbImage, err := imaging.Open(imageThumb)

	if err != nil {
		fmt.Printf("Unable to open %s\n", imageThumb)
		return
	}

	thumbWidth = imageThumbImage.Bounds().Size().X
	thumbHeight = imageThumbImage.Bounds().Size().Y

	return thumbWidth, thumbHeight
}
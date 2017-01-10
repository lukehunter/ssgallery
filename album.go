package main

import (
	"fmt"
	"path/filepath"
	"strings"
	"os"
)

type Album struct {
	name, folder string
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

// Traces path upward to root album. root is at 0 index of returned array.
func (a *Album) GetAlbumPath(albums []Album) []Album {
	albums = append(albums, *a)

	if a.parent == nil {
		// reverse it so 0 is root and n is leaf
		return reverse(albums)
	}

	return a.parent.GetAlbumPath(albums)
}

func reverse(albums []Album) []Album {
	for i := 0; i < len(albums)/2; i++ {
		j := len(albums) - i - 1
		albums[i], albums[j] = albums[j], albums[i]
	}
	return albums
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
					fmt.Printf("Unrecognized file extension on %s, ignoring. " +
						"Accepted are %s, override this behavior with --%s\n",
						image, strings.Join(valid, ", "), skipextcheckarg)
					continue
				}
			}

			if !options.verifyimages || ValidImage(image) {
				curAlbum.AddImage(Image{name: fileNameWithoutExtension(image), sourcePath: image})
			}
		}

		a.AddAlbum(*curAlbum)
	}
}

func (a *Album) UpdateImageRenditions(targetPath string) {
	if !a.HasImages() {
		fmt.Printf("Album %s has no images, skipping resize step\n", a.name)
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

	// Always generate album thumbnail in case thumbnail.jpg removed
	if exists, _ := exists(albumThumbnailPath); exists {
		SaveResizedImage(&Image{name: fmt.Sprintf("%s thumbnail", a.name), sourcePath: albumThumbnailPath},
			options.thumbwidth, options.thumbheight, targetAlbumThumbnailPath, true, false)
	} else {
		image := a.FindThumbnail()
		fmt.Printf("File %s not found, using %s as album thumbnail\n", albumThumbnailPath, image.sourcePath)
		SaveResizedImage(image, options.thumbwidth, options.thumbheight, targetAlbumThumbnailPath, true, false)
	}
}
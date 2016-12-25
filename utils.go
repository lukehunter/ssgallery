package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// exists returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getDirList(path string) []string {
	return getItemList(path, false, true, false, false)
}

func getFileList(path string) []string {
	return getItemList(path, false, false, true, false)
}

func getItemList(path string, includeSelf, includeDirs, includeFiles, recursive bool) []string {
	itemList := []string{}

	err := filepath.Walk(path, func(curPath string, fi os.FileInfo, err error) error {
		checkDir := includeDirs && fi.IsDir()
		checkSelf := includeSelf && curPath == path
		checkFile := includeFiles && !fi.IsDir()
		recurse := recursive || filepath.Dir(curPath) == path

		if (checkSelf || checkDir || checkFile) && recurse {
			printlnIfTrue(fmt.Sprintf("Walk: Include %s", curPath), options.debug)
			itemList = append(itemList, curPath)
		} else {
			printlnIfTrue(fmt.Sprintf("Walk: Exclude %s", curPath), options.debug)
		}

		return nil
	})

	check(err)

	return itemList
}

func fileNameWithoutExtension(path string) string {
	filename := filepath.Base(path)

	basename := strings.TrimSuffix(filename, filepath.Ext(filename))

	return basename
}

func Copy(src, dst string) error {
	inHash, err := hash_file_md5(src)
	if err != nil {
		return err
	}

	outHash, outErr := hash_file_md5(dst)

	if outErr == nil && inHash == outHash {
		fmt.Printf("Skipping copy, %s has same md5sum as %s\n", src, dst)
		return nil
	} else if outErr != nil {
		fmt.Printf("Couldn't read hash for %s (%s), will attempt to copy\n", dst, outErr.Error())
	}

	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	fmt.Printf("Copying %s to %s\n", src, dst)

	_, err = io.Copy(out, in)
	cerr := out.Close()
	if err != nil {
		return err
	}

	filesTouched++

	return cerr
}

func formatFilename(imagename string, width, height int) string {
	return fmt.Sprintf("%s_%d_%d.jpg", imagename, width, height)
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func printlnIfTrue(msg string, flag bool) {
	if flag {
		fmt.Println(msg)
	}
}

func SaveResizedImage(imageInfo *Image, width, height int, filename string, isThumb, skipIfNewer bool) {
	printErr := func(err error) {
		fmt.Printf("SaveRezisedImage unexpected error: %s\n", err.Error())
	}

	updateImageDimensions := func(i image.Image) {
		newWidth := i.Bounds().Size().X
		newHeight := i.Bounds().Size().Y
		if isThumb {
			imageInfo.thumbWidth, imageInfo.thumbHeight = newWidth, newHeight
		} else {
			imageInfo.viewerWidth, imageInfo.viewerHeight = newWidth, newHeight
		}
	}

	fileExists, _ := exists(filename)

	if skipIfNewer && fileExists {
		imageInfo, err := os.Stat(imageInfo.sourcePath)
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
			existingRendition, err := imaging.Open(filename)

			if err != nil {
				printErr(err)
				return
			}

			updateImageDimensions(existingRendition)

			fmt.Printf("Skipping resizing for %s (target's last write time is newer than source)\n", path.Base(filename))
			return
		}
	}

	fmt.Printf("Generating %dx%d for %s\n", width, height, imageInfo.name)

	image, err := imaging.Open(imageInfo.sourcePath)

	if err != nil {
		printErr(err)
		return
	}

	image = imaging.Fit(image, width, height, imaging.Lanczos)

	updateImageDimensions(image)

	if fileExists {
		buf := new(bytes.Buffer)
		err = imaging.Encode(buf, image, imaging.JPEG)

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
			fmt.Printf("Skipping rendition, generated image has same md5sum as what is already in target folder (%s)\n", filename)
			return
		}
	}

	err = imaging.Save(image, filename)

	if err != nil {
		printErr(err)
	}

	filesTouched++
}

package main

import (
    "os"
    "path/filepath"
    "strings"
    "io"
    "fmt"
)

// exists returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return true, err
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func getDirList(path string) []string {
    return getItemList(path, true, false)
}

func getFileList(path string) []string {
    return getItemList(path, false, true)
}

func getItemList(path string, includeDirs bool, includeFiles bool) []string {
    itemList := []string{}

    err := filepath.Walk(options.source, func(path string, fi os.FileInfo, err error) error {
        if includeDirs && fi.IsDir() {
            append(itemList, path)
            return nil
        }

        if includeFiles && !fi.IsDir() {
            append(itemList, path)
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

func Copy(dst, src string) error {
    in, err := os.Open(src)
    if err != nil { return err }
    defer in.Close()
    out, err := os.Create(dst)
    if err != nil { return err }
    defer out.Close()
    _, err = io.Copy(out, in)
    cerr := out.Close()
    if err != nil { return err }
    return cerr
}

func formatFilename(imagename string, width, height int) string {
    return fmt.Sprintf("%s_%d_%d.jpg", imagename, width, height)
}
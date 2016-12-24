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

    err := filepath.Walk(path, func(curPath string, fi os.FileInfo, err error) error {
        if includeDirs && fi.IsDir() && curPath != options.source {
            itemList = append(itemList, curPath)
        }

        if includeFiles && !fi.IsDir() {
            itemList = append(itemList, curPath)
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
    in, err := os.Open(src)
    if err != nil { return err }
    defer in.Close()

    out, err := os.Create(dst)
    if err != nil { return err }
    defer out.Close()

    fmt.Printf("Copying %s to %s\n", src, dst)

    _, err = io.Copy(out, in)
    cerr := out.Close()
    if err != nil { return err }

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
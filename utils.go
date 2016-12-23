package main

import (
    "os"
    "path/filepath"
    "strings"
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
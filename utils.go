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
    if err != nil { return err }

    outHash, outErr := hash_file_md5(dst)

    if outErr == nil && inHash == outHash {
        fmt.Printf("Skipping copy, %s has same md5sum as %s", src, dst)
        return nil
    } else if outErr != nil {
        fmt.Printf("Couldn't read hash for %s (%s), will attempt to copy", dst, outErr.Error())
    }

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

    filesTouched++;

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
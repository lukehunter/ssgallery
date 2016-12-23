package main

import (
    "path"
)

type Image struct {
    name, path string
    width, height int
}

func (i *Image) filename() string {
    _, file := path.Split(i.name)
    return file
}

type Album struct {
    name, folder string
    images []Image
}

type Gallery struct {
    name string
    albums []Album
}
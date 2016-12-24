package main

import (
    "path/filepath"
)

type Options struct {
    name, source, target, baseurl, disqus string
    thumbwidth, thumbheight, viewerwidth, viewerheight int
    skipextcheck bool
}

type Image struct {
    name, path string
    width, height int
}

func (i *Image) filename() string {
    return filepath.Base(i.path)
}

type Album struct {
    name, folder string
    images []Image
}

func (a *Album) AddImage(image Image) {
    a.images = append(a.images[:], image)
}

type Gallery struct {
    name string
    albums []Album
}

func (g *Gallery) AddAlbum(album Album) {
    g.albums = append(g.albums[:], album)
}
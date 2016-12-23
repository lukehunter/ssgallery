package main

import (
    "fmt"
    "strings"
    "os"
    "crypto/md5"
    "io"
    "io/ioutil"
    "errors"
    "bytes"
)

type TemplateItem struct {
    tag string
    values map[string]string
}

type Template struct {
    rawHtml string
    values  map[string]string
    items   []TemplateItem
}

func formatToken(s string) string {
    return fmt.Sprintf("%%%s%%", s)
}

func (t *Template) SetHiddenRegion(regionTag string, hidden bool) {
    startToken := fmt.Sprintf("%s_REGION_START", regionTag)
    endToken := fmt.Sprintf("%s_REGION_END", regionTag)

    var start, end string

    if hidden {
        start = "<!--"
        end = "-->"
    } else {
        start = ""
        end = ""
    }

    t.values[startToken] = start
    t.values[endToken] = end
}

func (t *Template) RenderHtml(filename string) {
    rendered := t.RenderItems()

    for k, v := range t.values {
        token := formatToken(k)
        strings.Replace(rendered, token, v, -1)
    }

    renderedBytes := []byte(rendered)

    // if the file exists, check if it has actually changed
    if exists, _ := exists(filename); exists {
        var existingHash, renderedHash [16]byte
        existingHash, err := hash_file_md5(filename)

        check(err)

        renderedHash = md5.Sum(renderedBytes)

        if (renderedHash == existingHash) {
            fmt.Printf("Skipping %s, existing file is up to date", filename)
            return;
        }
    }

    fmt.Printf("Writing %s", filename)
    err := ioutil.WriteFile(filename, renderedBytes, 0644)

    check(err)
}

func (t *Template) RenderItems() string {
    if len(t.items) == 0 {
        return t.rawHtml
    }

    rendered := t.rawHtml
    itemInfo := t.items[0]

    startTag := fmt.Sprintf("<!-- %%%s_START%% -->", itemInfo.tag)
    endTag := fmt.Sprintf("<!-- %%%s_END%% -->", itemInfo.tag)

    startIndex := strings.Index(rendered, startTag)

    if startIndex < 0 {
        panic(errors.New(fmt.Sprintf("could not find item start tag %s", startTag)))
    }

    endIndex := strings.Index(rendered, endTag)

    if endIndex < 0 {
        panic(errors.New(fmt.Sprintf("could not find item end tag %s", endTag)))
    }

    itemTemplate := rendered[startIndex:endIndex]
    aboveItems := rendered[0:startIndex]
    belowItems := rendered[endIndex:]

    var itemsHtml bytes.Buffer

    for _,item := range t.items {
        curItemHtml := itemTemplate

        for k,v := range item.values {
            token := formatToken(k)
            curItemHtml = strings.Replace(curItemHtml, token, v, -1)
        }

        itemsHtml.WriteString(curItemHtml)
        itemsHtml.WriteString("\n")
    }

    rendered = aboveItems + itemsHtml.String() + belowItems

    return rendered
}

func hash_file_md5(filePath string) ([16]byte, error) {
    var result [16]byte

    //Open the passed argument and check for any error
    file, err := os.Open(filePath)
    if err != nil {
        return result, err
    }

    defer file.Close()

    hash := md5.New()

    if _, err := io.Copy(hash, file); err != nil {
        return result, err
    }

    var hashInBytes [16]byte
    copy(hashInBytes[:], hash.Sum(nil)[:16])

    return hashInBytes, nil

}
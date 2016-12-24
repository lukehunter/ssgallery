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
    values *map[string]string
}

func NewTemplateItem(tag string) *TemplateItem {
    ti := new(TemplateItem)
    ti.tag = tag
    ti.values = &map[string]string{}
    return ti
}

func (t *TemplateItem) AddValues(values map[string]string) {
    for k, v := range values {
        (*t.values)[k] = v
    }
}

type Template struct {
    rawHtml string
    values  map[string]string
    lists   map[string]*[]TemplateItem
}

func formatToken(s string) string {
    return fmt.Sprintf("%%%s%%", s)
}

func NewTemplate(rawHtml string) *Template {
    t := new(Template)
    t.rawHtml = rawHtml
    t.values = map[string]string{}
    t.lists = make(map[string]*[]TemplateItem)
    return t
}

func (t *Template) AddValues(values map[string]string) {
    for k, v := range values {
        t.values[k] = v
    }
}

func (t *Template) AddItem(item TemplateItem) {
    if t.lists[item.tag] == nil {
        t.lists[item.tag] = &[]TemplateItem{}
    }
    newList := append((*t.lists[item.tag])[:], item)
    t.lists[item.tag] = &newList
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

    printlnIfTrue("template values {", options.debug)
    for k, v := range t.values {
        printlnIfTrue(fmt.Sprintf("%s : %s", k, v), options.debug)
        token := formatToken(k)
        rendered = strings.Replace(rendered, token, v, -1)
    }
    printlnIfTrue("} end template values", options.debug)

    renderedBytes := []byte(rendered)

    // if the file exists, check if it has actually changed
    if exists, _ := exists(filename); exists {
        var existingHash, renderedHash [16]byte
        existingHash, err := hash_file_md5(filename)

        check(err)

        renderedHash = md5.Sum(renderedBytes)

        if (renderedHash == existingHash) {
            fmt.Printf("Skipping %s, existing file is up to date\n", filename)
            return;
        }
    }

    fmt.Printf("Writing %s\n", filename)
    err := ioutil.WriteFile(filename, renderedBytes, 0755)

    check(err)

    filesTouched++
}

func (t *Template) RenderItems() string {
    if len(t.lists) == 0 {
        printlnIfTrue("no template items", options.debug)
        return t.rawHtml
    }

    if options.debug {
        for k, v := range t.lists {
            fmt.Printf("%s: \n", k)
            for _, item := range *v {
                fmt.Printf("%s{", item.tag)
                for k, v := range *item.values {
                    fmt.Printf("\n\t%s:%s", k, v)
                }
                fmt.Printf("}\n")
            }
            fmt.Printf("\n\n")
        }
    }

    rendered := t.rawHtml
    var itemsHtml bytes.Buffer

    for tag,_ := range t.lists {
        startTag := fmt.Sprintf("<!-- %%%s_START%%", tag)
        endTag := fmt.Sprintf("%%%s_END%% -->", tag)

        startIndex := strings.Index(rendered, startTag)

        if startIndex < 0 {
            fmt.Println(rendered)
            panic(errors.New(fmt.Sprintf("could not find item start tag %s in above template", startTag)))
        }

        aboveItems := rendered[0:startIndex]
        // offset templateStartIndex to skip past the actual tag
        startIndex += len(startTag)

        endIndex := strings.Index(rendered, endTag)

        if endIndex < 0 {
            fmt.Println(rendered)
            panic(errors.New(fmt.Sprintf("could not find item end tag %s in above template", endTag)))
        }

        belowItems := rendered[endIndex + len(endTag):]

        for _, tplItemList := range *(t.lists[tag]) {
            curItemTemplate := rendered[startIndex:endIndex]

            for k,v := range *tplItemList.values {
                token := formatToken(k)
                curItemTemplate = strings.Replace(curItemTemplate, token, v, -1)
            }

            itemsHtml.WriteString(curItemTemplate)
            itemsHtml.WriteString("\n")
        }

        rendered = aboveItems + itemsHtml.String() + belowItems

        itemsHtml.Reset()
    }

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
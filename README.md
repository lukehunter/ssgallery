# Stupidly Simple Gallery
## Introduction
Stupidly Simple Gallery (ssgallery) is a command line based cross platform static image gallery generator written in Go

I created this software out of frustration with all the bloated and unnecessarily complicated gallery software out there. Would you believe that some gallery software actually waits until somebody requests a thumbnail to resize the source image? This strategy does not work well when using high resolution (10+ megapixel) source images. This software is designed to support as large of images as you like, pre-cache everything and be extremely simple (run it on a folder full of files, upload the result to your web host, and you're done).

## Sample Gallery
See http://lukehunter.net/wallpapers/

## Usage
The following is an example command line for a gallery hosted at http://server/mywebgallery/:

### Windows

ssgallery_windows_amd64.exe --source "C:\Users\luke\pictures\" --target "c:\inetpub\wwwroot\mywebgallery" --name "My Web Gallery" --thumbwidth 170 --thumbheight 130 --viewerwidth 1500 --viewerheight 1000 --baseurl "/mywebgallery" --disqus "//your-disqus-url.disqus.com/embed.js"

### OSX/Linux

ssgallery_darwin_amd64 --source "/users/luke/pictures/" --target "~/Sites/vhost/mywebgallery" --name "My Web Gallery" --thumbwidth 170 --thumbheight 130 --viewerwidth 1500 --viewerheight 1000 --baseurl "/mywebgallery" --disqus "//your-disqus-url.disqus.com/embed.js"

## Download

- Windows [32-bit](http://lukehunter.net/ssgallery_release/ssgallery_windows_386.exe) | [64-bit](http://lukehunter.net/ssgallery_release/ssgallery_windows_amd64.exe)
- OSX [32-bit](http://lukehunter.net/ssgallery_release/ssgallery_darwin_386) | [64-bit](http://lukehunter.net/ssgallery_release/ssgallery_darwin_amd64)
- Linux [32-bit](http://lukehunter.net/ssgallery_release/ssgallery_linux_386) | [64-bit](http://lukehunter.net/ssgallery_release/ssgallery_linux_amd64)
- Other: See binaries [here](http://lukehunter.net/ssgallery_release)

--or-- if you have go installed

go get github.com/lukehunter/ssgallery

## Design

ssgallery uses an extremely basic template system to generate html pages. Tokens such as %SSG_ALBUM_NAME% and %SSG_IMAGE_URL% are inserted into an html file, and at run-time are replaced with the correct strings. There is also support for lists of items with nested tokens (e.g. on the album page which shows lists of thumbnails), although the templating engine has not been heavily tested with custom layouts.

An ssgallery theme consists of two templates:

- Album
  - Shows list of albums and/or images
- Image
  - Shows single image at resolution specified on command line
  - Includes download and view links for original resolution
  - Includes disqus comment area specific to the image (omitted if no disqus argument is provided)
  - Navigate forward and backward by swiping on tablets and phones
  - Navigate forward and backward with onhover buttons
  - Navigate forward by clicking on the image
  
## Input
Input to ssgallery is a hierarchical folder structure, with an optional thumbnail.jpg in each folder that will be used as the album cover photo (if thumbnail.jpg is not present the first image in the album will be used). In addition there are command line options to control thumbnail and image viewing sizes and to specify the base relative url.

## Output
ssgallery writes a complete web page to the target folder. Files in the target folder are only touched if they have changed.

The resulting folder can be transferred via FTP to a webhost. Since files that are up to date are not touched, using the "Overwrite if source file is newer" option in your FTP client (e.g. FileZilla) will make incremental updates faster.

## Deployment Steps
### 1. Prepare files
ssgallery expects a folder structure like the following (it can have as many subfolders as you like). It is recommended that all images (including thumbnail.jpg) be at the largest resolution you would like available for users to download. They will be resized for thumbnails and the image viewer depending on the provided command line arguments, and the original file will be available under the View/Download link on the Image page.

- pictures
  - album1
    - thumbnail.jpg
    - image1.jpg
    - image2.jpg
    - image3.jpg
  - album2
    - album2a
      - image1.jpg
      - image2.jpg
    - album2b
      - thumbnail.jpg
      - image1.jpg
      - image2.jpg
    - thumbnail.jpg 
    - image1.jpg
    - image2.jpg
  - album3
    - image1.jpg
    - image2.jpg
    - image3.jpg
    - image4.jpg
  - image1.jpg
  - image2.jpg
      
Note that album3 does not have a thumbnail.jpg -- in this case album3/image1.jpg will be used as the album3 thumbnail by default. Also if you happen to have extra files in these folders that aren't in a supported image format they will be ignored.
      
### 2. Run ssgallery

See above for usage instructions.

### 3. Upload the target folder to your web host via FTP

You're done!

## A Note about Disqus
Disqus identifies comment pages using a page identifier and/or a page url (see: https://help.disqus.com/customer/portal/articles/472098-javascript-configuration-variables). ssgallery uses an md5 hash of the original source image for the identifier, and does not bother to provide a url. This means as long as the source image is not modified, the comments should still show up correctly even if the filename or url changes. If you change domains after disqus comments have been entered, you may need to use the [Disqus migration tool](https://help.disqus.com/customer/portal/articles/286778-migration-tools) to move them.

## known issues / limitations
- Thumbnail sizes other than 170x130 are not currently supported. Some changes to the CSS are required.
- There is no support for paging. Really large albums may not scale well since all thumbnails are on one page. 
- Layout is a little small on mobile devices. Could use some responsive design.
- If you run ssgallery multiple times and remove some images in between runs, you will end up with extra files in your destination folder that aren't linked in the gallery but are still present. Over time this could grow and start wasting disk space (and you may want removed images to disappear completely from the published gallery). A simple workaround is to completely delete the target folder and re-generate from scratch as needed.
- Occasionally File.Copy throws an IOException related to disk space even when plenty of space is available. Setting the following registry key may help:

   HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Services\LanmanServer\Parameters\IRPStackSize DWORD 0x0000000f (15) 

## todo
- parallelization
- remove unneeded html
-   "       "    css
- fix css to work with custom thumbnail sizes
- download gallery/album links and zip files
- prompt to remove extraneous files at the end
- password protected albums
- gapless gallery/album layout
- more themes

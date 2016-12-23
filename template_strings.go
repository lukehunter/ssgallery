package main

var albumTemplate string = `
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN" "http://www.w3.org/TR/html4/strict.dtd">
<html lang="en" dir="ltr">
<head>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<meta name="generator" content="ssgallery">

	<meta name="description" content="%SSG_ALBUM_NAME%">

	<title>%SSG_ALBUM_NAME% | %SSG_GALLERY_NAME%</title>

	<link rel="stylesheet" type="text/css" href="../data/ssgallery.css">
	<link rel="start" title="Home" href="/nielsenphotos/" >
</head>

<body id="theCategoryPage" class="  ntf   ats    ">
	<div id="the_page">
        <div class="titrePage" id="imageHeaderBar">
            <div class="browsePath">
                <a href="%SSG_HOME_URL%">%SSG_GALLERY_NAME%</a> / %SSG_ALBUM_NAME%
            </div>
        </div>
		<div id="content" >
			<div id="content_cell">
				<div id="subcontent">
					<ul class="thumbnailCategories">

						<!-- %SSG_IMAGE_LIST_ITEM_START% -->
						<li onclick="window.location='%SSG_IMAGE_URL%';">
						    <div class="thumbnailCategory">
						        <div class="illustration">
						            <a href="%SSG_IMAGE_PAGE_URL%">
						                <img src="%SSG_IMAGE_THUMBNAIL_URL%" width="%SSG_IMAGE_THUMBNAIL_WIDTH%" height="%SSG_IMAGE_THUMBNAIL_HEIGHT%" alt="%SSG_IMAGE_NAME%">
						            </a>
						        </div>
						        <div class="description">
						            <h3>
						                <a href="%SSG_IMAGE_PAGE_URL%">%SSG_IMAGE_NAME%</a>
						            </h3>
						        </div>
						    </div>
						</li>
						<!-- %SSG_IMAGE_LIST_ITEM_END% -->

					</ul>
				</div>
			</div>
		</div>
	</div>
</div>
</div>
</body>
</html>
`

var galleryTemplate string = `
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN" "http://www.w3.org/TR/html4/strict.dtd">
<html lang="en" dir="ltr">
<head>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<meta name="generator" content="ssgallery">

	<meta name="description" content="%SSG_GALLERY_NAME%">

	<title>%SSG_GALLERY_NAME%</title>

	<link rel="stylesheet" type="text/css" href="data/ssgallery.css">
	<link rel="start" title="Home" href="/nielsenphotos/" >
</head>

<body id="theCategoryPage" class="  ntf   ats    ">
	<div id="the_page">
        <div class="titrePage" id="imageHeaderBar">
            <div class="browsePath">
                %SSG_GALLERY_NAME%
            </div>
        </div>
		<div id="content" >
			<div id="content_cell">
				<div id="subcontent">
					<ul class="thumbnailCategories">

						<!-- %SSG_ALBUM_LIST_ITEM_START% -->
						<li onclick="window.location='%SSG_ALBUM_URL%';">
						    <div class="thumbnailCategory">
						        <div class="illustration">
						            <a href="%SSG_ALBUM_URL%">
						                <img src="%SSG_ALBUM_NAME%/thumbnail.jpg" width="%SSG_ALBUM_THUMBNAIL_WIDTH%" height="%SSG_ALBUM_THUMBNAIL_HEIGHT%" alt="%SSG_ALBUM_NAME%">
						            </a>
						        </div>
						        <div class="description">
						            <h3>
						                <a href="%SSG_ALBUM_URL%">%SSG_ALBUM_NAME%</a>
						            </h3>
						        </div>
						    </div>

						</li>
						<!-- %SSG_ALBUM_LIST_ITEM_END% -->

					</ul>
				</div>
			</div>
		</div>
	</div>
</div>
</div>
</body>
</html>
`

var imageTemplate string = `
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN" "http://www.w3.org/TR/html4/strict.dtd">
<html lang="en" dir="ltr">
<head>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<meta name="generator" content="ssgallery">

	<title>%SSG_IMAGE_NAME% | %SSG_GALLERY_NAME%</title>

	<link rel="start" title="Home" href="%SSG_HOME_URL%" >
	<link rel="prev" title="Prev" href="%SSG_PREV_IMAGE_PAGE_URL%" >
	<link rel="next" title="Next" href="%SSG_NEXT_IMAGE_PAGE_URL%" >
    %SSG_NEXT_IMAGE_LINK_REGION_START%
	<link rel="preload" href="%SSG_PRELOAD_URL%">
    %SSG_NEXT_IMAGE_LINK_REGION_END%

	<script type="text/javascript">
	// configuration options
	    var options = {
	        imageAutosize:true,
	        imageAutosizeMargin:60,
	        imageAutosizeMinHeight:200,
	        themeStyle:"white",
	        animatedTabs:true,
	        defaultTab:"none",
	        marginContainer:30,
	        paddingContainer:10,
	        defaultZoomSize:"full",
	        highResClickMode:"zoom",
	        navArrows:true

	    }
	</script>

	<script type="text/javascript">
	    document.documentElement.className = 'js';
	</script>

    <script type="text/javascript">
        function processingRoutine() {
            var swipedElement = document.getElementById(triggerElementID);
            if ( swipeDirection == 'left' ) {
                window.location.href = "%SSG_NEXT_IMAGE_PAGE_URL%";
            } else if (swipeDirection == 'right') {
                window.location.href = "%SSG_PREV_IMAGE_PAGE_URL%";
            }
        }
    </script>

    <link rel="stylesheet" type="text/css" href="../data/ssgallery.css">

    <script type="text/javascript" src="../data/js/3r5d4l.js"></script>
    <script type="text/javascript" src="../data/js/touchevents.js"></script>
</head>

<body id="thePicturePage" class="  ntf   ats    ">
	<div id="the_page">

			<script type="text/javascript">
			    var image_w = %SSG_IMAGE_WIDTH%
			        var image_h = %SSG_IMAGE_HEIGHT%
			</script>

			<div class="titrePage" id="imageHeaderBar">
				<div class="browsePath">
					<a href="%SSG_HOME_URL%">%SSG_GALLERY_NAME%</a> / <a href="%SSG_ALBUM_URL%">%SSG_ALBUM_NAME%</a> / %SSG_IMAGE_NAME%
				</div>
			</div>

			<div id="content">
			    <div id="theImageAndTitle" ontouchstart="touchStart(event, 'theImageAndTitle');" ontouchend="touchEnd(event);" ontouchmove="touchMove(event);" ontouchcancel="touchCancel(event);">
			        <div id="theImageBox" >
			            <div id="theImage">

			                <div id="theImg" class="img_frame">
			                    %SSG_PREV_IMAGE_LINK_REGION_START%
			                    <a href="%SSG_PREV_IMAGE_PAGE_URL%" class="img_nav img_prev">&nbsp;</a>
			                    %SSG_PREV_IMAGE_LINK_REGION_END%
                                %SSG_NEXT_IMAGE_LINK_REGION_START%
			                    <a href="%SSG_NEXT_IMAGE_PAGE_URL%" class="img_nav img_next">&nbsp;</a>
			                    %SSG_NEXT_IMAGE_LINK_REGION_END%
                                %SSG_NEXT_IMAGE_LINK_REGION_START%
			                    <a href="%SSG_NEXT_IMAGE_PAGE_URL%">
			                    %SSG_NEXT_IMAGE_LINK_REGION_END%
			                        <img src="%SSG_IMAGE_URL%" width="%SSG_IMAGE_WIDTH%" height="%SSG_IMAGE_HEIGHT%" alt="%SSG_IMAGE_NAME%" id="theMainImage" class="hideTabs">
                                %SSG_NEXT_IMAGE_LINK_REGION_START%
			                    </a>
                                %SSG_NEXT_IMAGE_LINK_REGION_END%
			                </div>
                        </div>
			        </div>
			        <div id="imageTitleContainer">
			            <div id="imageTitle">
			                %SSG_IMAGE_NAME%<br />
			                Full Size: <a href="%SSG_ORIG_IMAGE_URL%">View</a> | <a href="%SSG_ORIG_IMAGE_URL%" download>Download</a>
			            </div>
			        </div>

			    </div>
                <div id="disqus_thread"></div>
                <script type="text/javascript">
                    /**
                    *  RECOMMENDED CONFIGURATION VARIABLES: EDIT AND UNCOMMENT THE SECTION BELOW TO INSERT DYNAMIC VALUES FROM YOUR PLATFORM OR CMS.
                    *  LEARN WHY DEFINING THESE VARIABLES IS IMPORTANT: https://disqus.com/admin/universalcode/#configuration-variables*/
                    /*
                    var disqus_config = function () {
                    this.page.url = %SSG_IMAGE_PAGE_URL%;  // Replace PAGE_URL with your page's canonical URL variable
                    this.page.identifier = %SSG_IMAGE_ID%; // Replace PAGE_IDENTIFIER with your page's unique identifier variable
                    };
                    */
                    (function() { // DON'T EDIT BELOW THIS LINE
                        var d = document, s = d.createElement('script');
                        s.src = '%SSG_DISQUS_URL%';
                        s.setAttribute('data-timestamp', +new Date());
                        (d.head || d.body).appendChild(s);
                    })();
                </script>
			    <noscript>Please enable JavaScript to view the <a href="https://disqus.com/?ref_noscript">comments powered by Disqus.</a></noscript>

			</div>
		</dl>
	</div>
</body>
</html>
`
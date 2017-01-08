package main

var albumTemplateRaw string = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta name="generator" content="ssgallery">

    <meta name="description" content="%SSG_ALBUM_NAME%">

    <title>%SSG_ALBUM_NAME%</title>



    <!-- Bootstrap CSS -->
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-alpha.6/css/bootstrap.min.css" integrity="sha384-rwoIResjU2yc3z8GV/NPeZWAv56rSmLldC3R/AZzGRnGxQQKnKkoFVhFQhNUwEyJ" crossorigin="anonymous">

    <link rel="start" title="Home" href="%SSG_HOME_URL%" >

    <style>
        div.itemCaption {
            font-size: 1rem;
        }

        div.header {
            font-size: 2rem;
        }
    </style>
</head>
<body>

<div class="header">
            <!-- %SSG_BREADCRUMB_LIST_ITEM_START%
            <a href="%SSG_ALBUM_URL%">%SSG_ALBUM_NAME%</a>
            %SSG_BREADCRUMB_LIST_ITEM_END% -->
            %SSG_ALBUM_NAME%
</div>

<div class="container">

    <div class="row">

        <!-- %SSG_ALBUM_LIST_ITEM_START%
        <div class="col-lg-3 col-md-4 col-xs-6 thumb">
            <div>
                <div class="itemThumbnail">
                    <a href="%SSG_ALBUM_URL%">
                        <img src="%SSG_ALBUM_NAME%/thumbnail.jpg" width="%SSG_ALBUM_THUMBNAIL_WIDTH%" height="%SSG_ALBUM_THUMBNAIL_HEIGHT%" alt="%SSG_ALBUM_NAME%">
                    </a>
                </div>
                <div class="itemCaption">
                    <a href="%SSG_ALBUM_URL%">%SSG_ALBUM_NAME%</a>
                </div>
            </div>
        </div>
        %SSG_ALBUM_LIST_ITEM_END% -->

        <!-- %SSG_IMAGE_LIST_ITEM_START%
        <div class="col-lg-3 col-md-4 col-xs-6 thumb">
            <div>
                <div class="itemThumbnail">
                    <a href="%SSG_IMAGE_PAGE_URL%">
                        <img src="%SSG_IMAGE_THUMBNAIL_URL%" width="%SSG_IMAGE_THUMBNAIL_WIDTH%" height="%SSG_IMAGE_THUMBNAIL_HEIGHT%" alt="%SSG_IMAGE_NAME%">
                    </a>
                </div>
                <div class="itemCaption">
                    <a href="%SSG_IMAGE_PAGE_URL%">%SSG_IMAGE_NAME%</a>
                </div>
            </div>
        </div>
        %SSG_IMAGE_LIST_ITEM_END% -->

    </div>

    <!-- Footer -->
    <!--<footer>-->
        <!--<div class="row">-->
            <!--<div class="col-lg-12">-->
                <!--<p>%SSG_GALLERY_NAME%</p>-->
            <!--</div>-->
        <!--</div>-->
    <!--</footer>-->

</div>

<!-- jQuery first, then Tether, then Bootstrap JS. -->
<script src="https://code.jquery.com/jquery-3.1.1.slim.min.js" integrity="sha384-A7FZj7v+d/sdmMqp/nOQwliLvUsJfDHW+k9Omg/a/EheAdgtzNs3hpfag6Ed950n" crossorigin="anonymous"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/tether/1.4.0/js/tether.min.js" integrity="sha384-DztdAPBWPRXSA/3eYEEUWrWCy7G5KFbe8fFjk5JAIxUYHKkDx6Qin1DkWx51bBrb" crossorigin="anonymous"></script>
<script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-alpha.6/js/bootstrap.min.js" integrity="sha384-vBWWzlZJ8ea9aCX4pEW3rVHjgjt7zpkNpZk+02D9phzyeVkE+jo0ieGizqPLForn" crossorigin="anonymous"></script>
</body>
</html>
`

var imageTemplateRaw string = `
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN" "http://www.w3.org/TR/html4/strict.dtd">
<html lang="en" dir="ltr">
<head>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<meta name="generator" content="ssgallery">

	<title>%SSG_IMAGE_NAME% | %SSG_ALBUM_NAME%</title>

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

    <link rel="stylesheet" type="text/css" href="%SSG_HOME_URL%/data/ssgallery.css">

    <script type="text/javascript" src="%SSG_HOME_URL%/data/js/3r5d4l.js"></script>
    <script type="text/javascript" src="%SSG_HOME_URL%/data/js/touchevents.js"></script>
</head>

<body id="thePicturePage" class="  ntf   ats    ">
	<div id="the_page">

			<script type="text/javascript">
			    var image_w = %SSG_IMAGE_WIDTH%
			        var image_h = %SSG_IMAGE_HEIGHT%
			</script>

			<div class="titrePage" id="imageHeaderBar">
            <div class="browsePath">
                <!-- %SSG_BREADCRUMB_LIST_ITEM_START%
                <a href="%SSG_ALBUM_URL%">%SSG_ALBUM_NAME%</a> /
                %SSG_BREADCRUMB_LIST_ITEM_END% -->
                %SSG_IMAGE_NAME%
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
			    %SSG_DISQUS_REGION_START%
                <div id="disqus_thread"></div>
                <script type="text/javascript">
                    /**
                    *  RECOMMENDED CONFIGURATION VARIABLES: EDIT AND UNCOMMENT THE SECTION BELOW TO INSERT DYNAMIC VALUES FROM YOUR PLATFORM OR CMS.
                    *  LEARN WHY DEFINING THESE VARIABLES IS IMPORTANT: https://disqus.com/admin/universalcode/#configuration-variables*/

                    var disqus_config = function () {
                    this.page.identifier = '%SSG_IMAGE_DISQUS_ID%'; // Replace PAGE_IDENTIFIER with your page's unique identifier variable
                    };

                    (function() { // DON'T EDIT BELOW THIS LINE
                        var d = document, s = d.createElement('script');
                        s.src = '%SSG_DISQUS_URL%';
                        s.setAttribute('data-timestamp', +new Date());
                        (d.head || d.body).appendChild(s);
                    })();
                </script>
			    <noscript>Please enable JavaScript to view the <a href="https://disqus.com/?ref_noscript">comments powered by Disqus.</a></noscript>
			    %SSG_DISQUS_REGION_END%
			</div>
		</dl>
	</div>
</body>
</html>
`

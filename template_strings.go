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

	    <link rel="stylesheet" type="text/css" href="%SSG_CSS_URL%">

	    <style>
	        .thumbnailBox {
	            width: %SSG_GALLERY_THUMBNAIL_WIDTH%px;
	            height: %SSG_GALLERY_THUMBNAIL_HEIGHT%px;
	        }
	    </style>
	</head>
	<body>
		<div class="container-fluid">
			<div class="header">
			            <!-- %SSG_BREADCRUMB_LIST_ITEM_START%
			            <a href="%SSG_ALBUM_URL%">%SSG_ALBUM_NAME%</a> /
			            %SSG_BREADCRUMB_LIST_ITEM_END% -->
			            %SSG_ALBUM_NAME%
			</div>
			<hr style="margin-top: 0" />
		    <div class="row">

		        <!-- %SSG_ALBUM_LIST_ITEM_START%
		        <div class="col-lg-2 col-md-4 col-xs-6 thumb" style="text-align: center; margin-bottom: 20px">
		            <div class="wraptocenter" style="display: inline-block">
		                <div class="wraptocenter thumbnailBox">
		                    <a href="%SSG_ALBUM_URL%">
		                        <img src="%SSG_ALBUM_NAME%/thumbnail.jpg" width="%SSG_ALBUM_THUMBNAIL_WIDTH%" height="%SSG_ALBUM_THUMBNAIL_HEIGHT%" alt="%SSG_ALBUM_NAME%">
		                    </a>
		                </div>
		                <hr style="margin:5px" />
		                <div class="itemCaption" style="margin:0">
		                    <a href="%SSG_ALBUM_URL%">%SSG_ALBUM_NAME%</a>
		                </div>
		            </div>
		        </div>
		        %SSG_ALBUM_LIST_ITEM_END% -->

		        <!-- %SSG_IMAGE_LIST_ITEM_START%
		        <div class="col-lg-2 col-md-4 col-xs-6 thumb" style="text-align: center; margin-bottom: 20px">
		            <div class="wraptocenter" style="display: inline-block">
		                <div class="wraptocenter thumbnailBox">
		                    <a href="%SSG_IMAGE_PAGE_URL%">
		                        <img src="%SSG_IMAGE_THUMBNAIL_URL%" width="%SSG_IMAGE_THUMBNAIL_WIDTH%" height="%SSG_IMAGE_THUMBNAIL_HEIGHT%" alt="%SSG_IMAGE_NAME%">
		                    </a>
		                </div>
		                <hr style="margin:5px" />
		                <div class="itemCaption" style="margin:0">
		                    <a href="%SSG_IMAGE_PAGE_URL%">%SSG_IMAGE_NAME%</a>
		                </div>
		            </div>
		        </div>
		        %SSG_IMAGE_LIST_ITEM_END% -->

		    </div>
		</div>

		<!-- jQuery first, then Tether, then Bootstrap JS. -->
		<script src="https://code.jquery.com/jquery-3.1.1.slim.min.js" integrity="sha384-A7FZj7v+d/sdmMqp/nOQwliLvUsJfDHW+k9Omg/a/EheAdgtzNs3hpfag6Ed950n" crossorigin="anonymous"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/tether/1.4.0/js/tether.min.js" integrity="sha384-DztdAPBWPRXSA/3eYEEUWrWCy7G5KFbe8fFjk5JAIxUYHKkDx6Qin1DkWx51bBrb" crossorigin="anonymous"></script>
		<script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-alpha.6/js/bootstrap.min.js" integrity="sha384-vBWWzlZJ8ea9aCX4pEW3rVHjgjt7zpkNpZk+02D9phzyeVkE+jo0ieGizqPLForn" crossorigin="anonymous"></script>
	</body>
</html>
`

var imageTemplateRaw string = `
<!DOCTYPE html>
<html lang="en" dir="ltr">
	<head>
	    <meta charset="utf-8">
	    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
	    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	    <meta name="generator" content="ssgallery">

		<title>%SSG_IMAGE_NAME% | %SSG_ALBUM_NAME%</title>

		<!-- Bootstrap CSS -->
	    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-alpha.6/css/bootstrap.min.css" integrity="sha384-rwoIResjU2yc3z8GV/NPeZWAv56rSmLldC3R/AZzGRnGxQQKnKkoFVhFQhNUwEyJ" crossorigin="anonymous">

		<link rel="start" title="Home" href="%SSG_HOME_URL%" >
		<link rel="prev" title="Prev" href="%SSG_PREV_IMAGE_PAGE_URL%" >
		<link rel="next" title="Next" href="%SSG_NEXT_IMAGE_PAGE_URL%" >
	    %SSG_NEXT_IMAGE_LINK_REGION_START%
		<link rel="preload" href="%SSG_PRELOAD_URL%">
	    %SSG_NEXT_IMAGE_LINK_REGION_END%

		<link rel="stylesheet" type="text/css" href="%SSG_CSS_URL%">

	</head>
	<body>
		<div class="container-fluid">
			<div class="header">
		            <!-- %SSG_BREADCRUMB_LIST_ITEM_START%
		            <a href="%SSG_ALBUM_URL%">%SSG_ALBUM_NAME%</a> /
		            %SSG_BREADCRUMB_LIST_ITEM_END% -->
		            %SSG_IMAGE_NAME%
			</div>
			<hr style="margin-top: 0" />
			<div class="row">
			    <div class="col-1" style="padding: 0">
			        <a class="nav_link" href="%SSG_PREV_IMAGE_PAGE_URL%">
				        <div class="nav_prev">
				            <div class="wraptocenter" style="position: absolute; top: 50%; left: 50%; margin: -40px 0 0 -40px">
				                %SSG_PREV_IMAGE_LINK_REGION_START%
				                <div class="nav_prev_img" width="80" height="80"></div>
				                %SSG_PREV_IMAGE_LINK_REGION_END%
			                </div>
		                </div>
	                </a>
	            </div>
				<div class="col-10" style="padding: 0">
					<div class="wraptocenter" style="margin: auto; display: block">
		                <a href="%SSG_ORIG_IMAGE_URL%">
		                    <img class="viewer" src="%SSG_IMAGE_URL%" alt="%SSG_IMAGE_NAME%" id="theMainImage" class="hideTabs">
		                </a>
		            </div>
	            </div>
			    <div class="col-1" style="padding: 0">
			        <a class="nav_link" href="%SSG_NEXT_IMAGE_PAGE_URL%">
			            <div class="nav_next">
			                <div class="wraptocenter" style="position: absolute; top: 50%; right: 50%; margin:-40px -40px 0 0">
				                %SSG_NEXT_IMAGE_LINK_REGION_START%
				                <div class="nav_next_img" width="80" height="80"></div>
				                %SSG_NEXT_IMAGE_LINK_REGION_END%
				            </div>
				        </div>
	                </a>
	            </div>
		    </div>
		    <div class="row" style="margin-top: 10px">
		        <div class="col-1"></div>
	            <div class="col-10" style="text-align: center">
	                %SSG_IMAGE_NAME%<br />
		            Full Size: <a href="%SSG_ORIG_IMAGE_URL%">View</a> | <a href="%SSG_ORIG_IMAGE_URL%" download>Download</a>
	            </div>
	            <div class="col-1"></div>
		    </div>

		    %SSG_DISQUS_REGION_START%

			<div class="row">
				<div class="col-1"></div>
				<div class="col-10" style="text-align: center">
					<a class="show-comments" href="#">Show comments</a>
					<!-- The empty element required for Disqus to loads comments into -->
			        <div id="disqus_thread"></div>
			    </div>
			    <div class="col-1"></div>
			</div>

		    <noscript>Please enable JavaScript to view the <a href="https://disqus.com/?ref_noscript">comments powered by Disqus.</a></noscript>
		    %SSG_DISQUS_REGION_END%
		</div>

		<!-- jQuery first, then Tether, then Bootstrap JS. -->
		<script src="https://code.jquery.com/jquery-3.1.1.min.js" crossorigin="anonymous"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/tether/1.4.0/js/tether.min.js" crossorigin="anonymous"></script>
		<script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-alpha.6/js/bootstrap.min.js" crossorigin="anonymous"></script>

		<script>
			// Requires jQuery of course.
			$(document).ready(function() {
			    $('.show-comments').on('click', function(){
			          // ajax request to load the disqus javascript
			          $.ajax({
			                  type: "GET",
			                  url: '%SSG_DISQUS_URL%',
			                  dataType: "script",
			                  cache: true
			          });
			          // hide the button once comments load
			          $(this).fadeOut();
			    });
			});
		</script>
	</body>
</html>
`

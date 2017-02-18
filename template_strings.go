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

	    <link rel="stylesheet" type="text/css" href="%SSG_HOME_URL%data/ssgallery.css">
	    <link rel="stylesheet" type="text/css" href="%SSG_HOME_URL%data/default-skin/default-skin.css">

	    <link href="%SSG_HOME_URL%data/photoswipe.css" rel="stylesheet" />
		<script src="%SSG_HOME_URL%data/photoswipe.min.js"></script>
		<script src="%SSG_HOME_URL%data/photoswipe-ui-default.min.js"></script>

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
		    <div class="row gallery">

		        <!-- %SSG_ALBUM_LIST_ITEM_START%
		        <div class="col-xl-2 col-lg-3 col-md-4 col-xs-6 thumb" style="text-align: center; margin-bottom: 20px">
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
		        <div class="col-xl-2 col-lg-3 col-md-4 col-xs-6 thumb" data-size="%SSG_IMAGE_WIDTH%x%SSG_IMAGE_HEIGHT%" data-url="%SSG_IMAGE_URL%" data-thumb-url="%SSG_IMAGE_THUMBNAIL_URL%" data-title="%SSG_IMAGE_NAME%" data-index="%SSG_IMAGE_INDEX%" style="text-align: center; margin-bottom: 20px">
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

		<!-- Root element of PhotoSwipe. Must have class pswp. -->
		<div class="pswp" tabindex="-1" role="dialog" aria-hidden="true">

		    <!-- Background of PhotoSwipe.
		         It's a separate element as animating opacity is faster than rgba(). -->
		    <div class="pswp__bg"></div>

		    <!-- Slides wrapper with overflow:hidden. -->
		    <div class="pswp__scroll-wrap">

		        <!-- Container that holds slides.
		            PhotoSwipe keeps only 3 of them in the DOM to save memory.
		            Don't modify these 3 pswp__item elements, data is added later on. -->
		        <div class="pswp__container">
		            <div class="pswp__item"></div>
		            <div class="pswp__item"></div>
		            <div class="pswp__item"></div>
		        </div>

		        <!-- Default (PhotoSwipeUI_Default) interface on top of sliding area. Can be changed. -->
		        <div class="pswp__ui pswp__ui--hidden">

		            <div class="pswp__top-bar">

		                <!--  Controls are self-explanatory. Order can be changed. -->

		                <div class="pswp__counter"></div>

		                <button class="pswp__button pswp__button--close" title="Close (Esc)"></button>

		                <button class="pswp__button pswp__button--share" title="Share"></button>

		                <button class="pswp__button pswp__button--fs" title="Toggle fullscreen"></button>

		                <button class="pswp__button pswp__button--zoom" title="Zoom in/out"></button>

		                <!-- Preloader demo http://codepen.io/dimsemenov/pen/yyBWoR -->
		                <!-- element will get class pswp__preloader--active when preloader is running -->
		                <div class="pswp__preloader">
		                    <div class="pswp__preloader__icn">
		                      <div class="pswp__preloader__cut">
		                        <div class="pswp__preloader__donut"></div>
		                      </div>
		                    </div>
		                </div>
		            </div>

		            <div class="pswp__share-modal pswp__share-modal--hidden pswp__single-tap">
		                <div class="pswp__share-tooltip"></div>
		            </div>

		            <button class="pswp__button pswp__button--arrow--left" title="Previous (arrow left)">
		            </button>

		            <button class="pswp__button pswp__button--arrow--right" title="Next (arrow right)">
		            </button>

		            <div class="pswp__caption">
		                <div class="pswp__caption__center"></div>
		            </div>

		        </div>

		    </div>

		</div>

		<script>
    (function() {
		debugger;
		var initPhotoSwipeFromDOM = function(gallerySelector) {

			var parseThumbnailElements = function(el) {
			    var thumbElements = el.childNodes,
			        numNodes = thumbElements.length,
			        items = [],
			        el,
			        childElements,
			        thumbnailEl,
			        size,
			        item;

			    for(var i = 0; i < numNodes; i++) {
			        el = thumbElements[i];

			        // include only element nodes
			        if(el.nodeType !== 1 || !el.getAttribute('data-url')){
			          continue;
			        }

			        size = el.getAttribute('data-size').split('x');

			        // create slide object
			        item = {
						src: el.getAttribute('data-url'),
						w: parseInt(size[0], 10),
						h: parseInt(size[1], 10),
						msrc: el.getAttribute('data-thumb-url'),
						title: el.getAttribute('data-title')
			        };

			        item.el = el; // save link to element for getThumbBoundsFn

					var mediumSrc = el.getAttribute('data-med');
		          	if(mediumSrc) {
		            	size = el.getAttribute('data-med-size').split('x');
		            	// "medium-sized" image
		            	item.m = {
		              		src: mediumSrc,
		              		w: parseInt(size[0], 10),
		              		h: parseInt(size[1], 10)
		            	};
		          	}
		          	// original image
		          	item.o = {
		          		src: item.src,
		          		w: item.w,
		          		h: item.h
		          	};

			        items.push(item);
			    }

			    return items;
			};

			// find nearest parent element
			var closest = function closest(el, fn) {
			    return el && ( fn(el) ? el : closest(el.parentNode, fn) );
			};

			var onThumbnailsClick = function(e) {
			    e = e || window.event;
			    e.preventDefault ? e.preventDefault() : e.returnValue = false;

			    var eTarget = e.target || e.srcElement;

			    var clickedListItem = closest(eTarget, function(el) {
			        return el.tagName === 'A';
			    });

			    if(!clickedListItem) {
			        return;
			    }

			    var clickedGallery = clickedListItem.parentNode.parentNode.parentNode.parentNode;

			    var clickedDiv = clickedListItem.parentNode.parentNode.parentNode;

				var index = parseInt(clickedDiv.getAttribute('data-index'));

			    if(index >= 0) {
			        openPhotoSwipe( index, clickedGallery );
			    }
			    return false;
			};

			var photoswipeParseHash = function() {
				var hash = window.location.hash.substring(1),
			    params = {};

			    if(hash.length < 5) { // pid=1
			        return params;
			    }

			    var vars = hash.split('&');
			    for (var i = 0; i < vars.length; i++) {
			        if(!vars[i]) {
			            continue;
			        }
			        var pair = vars[i].split('=');
			        if(pair.length < 2) {
			            continue;
			        }
			        params[pair[0]] = pair[1];
			    }

			    if(params.gid) {
			    	params.gid = parseInt(params.gid, 10);
			    }

			    return params;
			};

			var openPhotoSwipe = function(index, galleryElement, disableAnimation, fromURL) {
			    var pswpElement = document.querySelectorAll('.pswp')[0],
			        gallery,
			        options,
			        items;

				items = parseThumbnailElements(galleryElement);

			    // define options (if needed)
			    options = {

			        galleryUID: galleryElement.getAttribute('data-pswp-uid'),

			        getThumbBoundsFn: function(index) {
			            // See Options->getThumbBoundsFn section of docs for more info
			            var thumbnail = items[index].el.children[0].children[0].children[0].children[0],
			                pageYScroll = window.pageYOffset || document.documentElement.scrollTop,
			                rect = thumbnail.getBoundingClientRect();

			            return {x:rect.left, y:rect.top + pageYScroll, w:rect.width};
			        },

			        addCaptionHTMLFn: function(item, captionEl, isFake) {
						if(!item.title) {
							captionEl.children[0].innerText = '';
							return false;
						}
						captionEl.children[0].innerHTML = item.title +  '<br/><small>Photo: ' + item.author + '</small>';
						return true;
			        },

			    };


			    if(fromURL) {
			    	if(options.galleryPIDs) {
			    		// parse real index when custom PIDs are used
			    		// http://photoswipe.com/documentation/faq.html#custom-pid-in-url
			    		for(var j = 0; j < items.length; j++) {
			    			if(items[j].pid == index) {
			    				options.index = j;
			    				break;
			    			}
			    		}
				    } else {
				    	options.index = parseInt(index, 10) - 1;
				    }
			    } else {
			    	options.index = parseInt(index, 10);
			    }

			    // exit if index not found
			    if( isNaN(options.index) ) {
			    	return;
			    }



				var radios = document.getElementsByName('gallery-style');
				for (var i = 0, length = radios.length; i < length; i++) {
				    if (radios[i].checked) {
				        if(radios[i].id == 'radio-all-controls') {

				        } else if(radios[i].id == 'radio-minimal-black') {
				        	options.mainClass = 'pswp--minimal--dark';
					        options.barsSize = {top:0,bottom:0};
							options.captionEl = false;
							options.fullscreenEl = false;
							options.shareEl = false;
							options.bgOpacity = 0.85;
							options.tapToClose = true;
							options.tapToToggleControls = false;
				        }
				        break;
				    }
				}

			    if(disableAnimation) {
			        options.showAnimationDuration = 0;
			    }

			    // Pass data to PhotoSwipe and initialize it
			    gallery = new PhotoSwipe( pswpElement, PhotoSwipeUI_Default, items, options);

			    // see: http://photoswipe.com/documentation/responsive-images.html
				var realViewportWidth,
				    useLargeImages = false,
				    firstResize = true,
				    imageSrcWillChange;

				gallery.listen('beforeResize', function() {

					var dpiRatio = window.devicePixelRatio ? window.devicePixelRatio : 1;
					dpiRatio = Math.min(dpiRatio, 2.5);
				    realViewportWidth = gallery.viewportSize.x * dpiRatio;


				    if(realViewportWidth >= 1200 || (!gallery.likelyTouchDevice && realViewportWidth > 800) || screen.width > 1200 ) {
				    	if(!useLargeImages) {
				    		useLargeImages = true;
				        	imageSrcWillChange = true;
				    	}

				    } else {
				    	if(useLargeImages) {
				    		useLargeImages = false;
				        	imageSrcWillChange = true;
				    	}
				    }

				    if(imageSrcWillChange && !firstResize) {
				        gallery.invalidateCurrItems();
				    }

				    if(firstResize) {
				        firstResize = false;
				    }

				    imageSrcWillChange = false;

				});

				gallery.listen('gettingData', function(index, item) {
				    if( useLargeImages ) {
				        item.src = item.o.src;
				        item.w = item.o.w;
				        item.h = item.o.h;
				    } else {
				        item.src = item.m.src;
				        item.w = item.m.w;
				        item.h = item.m.h;
				    }
				});

			    gallery.init();
			};

			// select all gallery elements
			var galleryElements = document.querySelectorAll( gallerySelector );
			for(var i = 0, l = galleryElements.length; i < l; i++) {
				galleryElements[i].setAttribute('data-pswp-uid', i+1);
				galleryElements[i].onclick = onThumbnailsClick;
			}

			// Parse URL and open gallery if it contains #&pid=3&gid=1
			var hashData = photoswipeParseHash();
			if(hashData.pid && hashData.gid) {
				openPhotoSwipe( hashData.pid,  galleryElements[ hashData.gid - 1 ], true, true );
			}
		};

		initPhotoSwipeFromDOM('.gallery');
		}());
		</script>

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

		<noscript><style> .jsonly { display: none } </style></noscript>
	</head>
	<body>
		<div class="container-fluid">
			<!-- LARGE DEVICE LAYOUT -->
			<div class="header hidden-md-down">
		            <!-- %SSG_BREADCRUMB_LIST_ITEM_START%
		            <a href="%SSG_ALBUM_URL%">%SSG_ALBUM_NAME%</a> /
		            %SSG_BREADCRUMB_LIST_ITEM_END% -->
		            %SSG_IMAGE_NAME%
			</div>
			<hr class="hidden-md-down" style="margin-top: 0" />
			<div class="row hidden-md-down">
			    <div class="col-1" style="padding: 0">
	                %SSG_PREV_IMAGE_LINK_REGION_START%
			        <a class="nav_link" href="%SSG_PREV_IMAGE_PAGE_URL%">
				        <div class="nav_prev">
				            <div class="wraptocenter" style="position: absolute; top: 50%; left: 50%; margin: -30px 0 0 -30px">
				                <div class="nav_prev_img"></div>
			                </div>
		                </div>
	                </a>
	                %SSG_PREV_IMAGE_LINK_REGION_END%
	            </div>
				<div class="col-10" style="padding: 0">
					<div class="wraptocenter" style="margin: auto; display: block">
		                <a href="%SSG_ORIG_IMAGE_URL%">
		                    <img class="viewer" src="%SSG_IMAGE_URL%" alt="%SSG_IMAGE_NAME%" class="hideTabs">
		                </a>
		            </div>
	            </div>
			    <div class="col-1" style="padding: 0">
			        %SSG_NEXT_IMAGE_LINK_REGION_START%
			        <a class="nav_link" href="%SSG_NEXT_IMAGE_PAGE_URL%">
			            <div class="nav_next">
			                <div class="wraptocenter" style="position: absolute; top: 50%; right: 50%; margin:-30px -30px 0 0">
				                <div class="nav_next_img"></div>
				            </div>
				        </div>
	                </a>
	                %SSG_NEXT_IMAGE_LINK_REGION_END%
	            </div>
		    </div>

		    <!-- MEDIUM AND SMALLER DEVICE LAYOUT -->
		    <div class="row hidden-lg-up">
		        <div class="col-12" style="padding: 0">
					<div class="wraptocenter" style="margin: auto; display: block">
		                <a href="%SSG_ORIG_IMAGE_URL%">
		                    <img class="viewer" src="%SSG_IMAGE_URL%" alt="%SSG_IMAGE_NAME%" class="hideTabs">
		                </a>
		            </div>
		        </div>
		    </div>

			<!-- NAV BAR (HYBRID DEVICE LAYOUT) -->
		    <div class="row" style="margin-top: 10px">
		        <div class="col-2">
		            %SSG_PREV_IMAGE_LINK_REGION_START%
			        <a class="nav_link hidden-lg-up" href="%SSG_PREV_IMAGE_PAGE_URL%">
				        <div class="nav_prev">
			                <div class="nav_prev_img left-align"></div>
		                </div>
	                </a>
	                %SSG_PREV_IMAGE_LINK_REGION_END%
		        </div>
	            <div class="col-8" style="text-align: center">
	                <a href="%SSG_ALBUM_URL%">%SSG_ALBUM_NAME%</a> : %SSG_IMAGE_NAME%<br />
		            Full Size: <a href="%SSG_ORIG_IMAGE_URL%">View</a> | <a href="%SSG_ORIG_IMAGE_URL%" download>Download</a>

                    %SSG_DISQUS_REGION_START%
                    <br />
                    <a class="show-comments jsonly" href="#">Show comments</a>
		            <noscript>Please enable JavaScript to view the <a href="https://disqus.com/?ref_noscript">comments powered by Disqus.</a></noscript>
		            %SSG_DISQUS_REGION_END%

	            </div>
	            <div class="col-2">
	                %SSG_NEXT_IMAGE_LINK_REGION_START%
			        <a class="nav_link hidden-lg-up" href="%SSG_NEXT_IMAGE_PAGE_URL%">
			            <div class="nav_next">
			                <div class="nav_next_img right-align"></div>
				        </div>
	                </a>
	                %SSG_NEXT_IMAGE_LINK_REGION_END%
                </div>
		    </div>

		    %SSG_DISQUS_REGION_START%
		    <div class="row" style="margin-top: 15px">
		        <div class="col-12">
		            <!-- The empty element required for Disqus to loads comments into -->
			        <div id="disqus_thread"></div>
			    </div>
		    </div>
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

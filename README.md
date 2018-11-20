# crawld
Image crawler

## Requirements

* Crawl url for images
* Actual images should not be downloaded, store all image urls including full size and thumbnails, remove any duplicate images
* Sets of images should be kept together
* URLs should be stored in a general location, there should be no duplicates
* Needs to scale across distributed hosts

## Design

URLs are collected and persisted such that multiple runs of the tool does not cause redundant crawling.  Image URLs are also stored.

### URLs

This essentially stores all the URLs and sub-URLs found by the crawler, this should be persisted over time so if this tool is run multiple times, state persists.  Lookups to see if this URL has been accessed already should be fast.  The schema for this store should be simple, basically store the URL and whether or not it has been crawled.     

`url, state, meta, tags, timestamp`

### Image URLs

This stores all the image URLs, there should be no duplicates, sets of
images should be grouped together.  Should differentiate thumbnail from
full size images.  Each image URL should allow a simple rating.

`url, rating (1-5), type (1-2), link，tags`

# crawld
Image crawler

## Requirements

* Crawl search URL for product URLs based on search criteria
* Actual images should not be downloaded, store all image URLs, remove any duplicate images
* Sets of images should be kept together
* URLs should be stored without duplicates
* Needs to scale across distributed hosts

## Design

There are 3 sets of URLs tracked:

1.  product page URL
1.  product image URL
1.  Desired content URL

URLs must be persisted such that multiple runs of crawld does not cause redundant crawling.

### product URLs

This essentially stores all the product  URLs found by the crawler, this should be persisted over time so if this tool is run multiple times, state can be resumed.  Lookups to see if this URL has been accessed already should be fast.  The schema for this store should be simple, basically store the URL and whether or not it has been crawled.

`url, state, meta, tags, timestamp`

### Image URLs

This stores all the image URLs, there should be no duplicates, sets of images should be grouped together.  Should differentiate thumbnail from full size images.  Each image URL should allow a simple rating.

`url, rating (1-5), type (1-2), link，tags`

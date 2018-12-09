# crawld
Image crawler

## Requirements

* Background service
* RESTful API
* Crawl Search URL for product URLs based on search criteria
* Actual images should not be downloaded, store all image URLs, remove any duplicate images
* Sets of images should be kept together
* URLs should be stored without duplicates
* Needs to scale across distributed hosts

## Setup

1.  `go get github.com/PuerkitoBio/goquery`
1.  `go get golang.org/x/net/html`
1.  `go get github.com/go-redis/redis`

`go build`

## Usage

`./crawd`

`Command line flags override crawld.json`

## Design

There are 2 sets of URLs tracked:

1.  product page URL
1.  product image URL

URLs must be persisted such that multiple runs of crawld does not cause redundant crawling.

### API

Endpoint | Parameters
--- | ---
GET /api/images | Retrieves image from unfiltered queue
POST /api/models | Creates a model record in database


### Product URLs

This essentially stores all the product URLs found by the crawler, this should be persisted over time so if this tool is run multiple times, state can be resumed.  Lookups to see if this URL has been accessed already should be fast.  The schema for this store should be simple, basically store the URL with no duplicates.

We use 2 datatypes for this, redis set and redis list.

Data Structure | Description | Name
--- | --- | ---
Redis Set | Guarantees uniqueness, fast lookup | crawld.urls.products 
Redis List | Work queue | crawld.queue.products

### Image URLs

This stores all the image URLs, there should be no duplicates, sets of images should be grouped together.  Should differentiate thumbnail from full size images.  Each image URL should allow a simple rating.

Data Structure | Description | Name
--- | --- | ---
Redis Set | Guarantees uniqueness | crawld.urls.images
Redis List | Work queue | crawld.queue.images
Sqlite3 | Query-able | Tags


## Queries

1.  Highest rating, based on categories
1.  Latest photos 

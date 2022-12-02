package main

var listingAPI = "https://tiki.vn/api/personalish/v1/blocks/listings"

// ?limit=${limit}&category=${category}&page=${page}
type categoryListing struct {
	page, limit, category int
}

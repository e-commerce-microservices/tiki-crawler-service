package main

// Product information
type Product struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	URLKey          string `json:"url_key"`
	BrandName       string `json:"brand_name"`
	Price           int    `json:"price"`
	Discount        int    `json:"discount"`
	DiscountRate    int    `json:"discount_rate"`
	RatingAverage   int    `json:"rating_average"`
	ReviewCount     int    `json:"review_count"`
	ThumbnailURL    string `json:"thumbnail_url"`
	ThumbnailWidth  int    `json:"thumbnail_width"`
	ThumbnailHeight int    `json:"thumbnail_height"`
	SellerProductID int    `json:"seller_product_id"`
}

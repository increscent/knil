package main

type link struct {
	linkId           uint
	PublicId         string `json:"link_id"`
	Title            string `json:"title"`
	Url              string `json:"url"`
	categoryId       uint
	CategoryPublicId string `json:"category_id"`
	IsValid          bool   `json:"is_valid"`
	PostedDate       date   `json:"posted_date"`
}

type date struct {
	Year  uint `json:"year"`
	Month uint `json:"month"`
	Day   uint `json:"day"`
}

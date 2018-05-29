package main

import (
	"time"
)

type link struct {
	linkId           uint32
	PublicId         string `json:"link_id"`
	Title            string `json:"title"`
	Url              string `json:"url"`
	categoryId       uint32
	CategoryPublicId string    `json:"category_id"`
	IsValid          bool      `json:"is_valid"`
	PostedDate       time.Time `json:"posted_date"`
}

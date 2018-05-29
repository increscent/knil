package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "robert:robert@/knil?parseTime=true")

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func queryLinks(db *sql.DB) ([]link, error) {
	result, err := db.Query("SELECT linkId, links.publicId, title, url, links.categoryId, categories.publicId, isValid, postedEpoch FROM links inner join categories on links.categoryId = categories.categoryId")
	if err != nil {
		return nil, err
	}

	var (
		linkId           uint32
		publicId         string
		title            string
		url              string
		categoryId       sql.NullInt64
		publicCategoryId string
		isValid          sql.RawBytes
		postedEpoch      uint64
		links            []link
	)

	links = *new([]link)

	for result.Next() {
		if err = result.Scan(&linkId, &publicId, &title, &url, &categoryId, &publicCategoryId, &isValid, &postedEpoch); err != nil {
			log.Fatal(err)
		}

		l := link{
			linkId,
			publicId,
			title,
			url,
			(uint32)(categoryId.Int64),
			publicCategoryId,
			(isValid[0] != 0x00),
			time.Unix((int64)(postedEpoch), 0),
		}

		links = append(links, l)
	}

	return links, nil
}

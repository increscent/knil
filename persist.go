package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
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

func insertLink(db *sql.DB, l link) (*link, error) {
	rows, err := db.Query(`SELECT categoryId
		FROM categories
		WHERE publicId = ?`,
		l.CategoryPublicId,
	)

	if err != nil {
		return nil, err
	}

	var categoryId uint
	if !rows.Next() {
		return nil, fmt.Errorf("Category not found: %s", l.CategoryPublicId)
	}
	if err = rows.Scan(&categoryId); err != nil {
		return nil, err
	}

	result, err := db.Exec(`INSERT INTO links (
		publicId,
		categoryId,
		title,
		url,
		isValid,
		postedYear,
		postedMonth,
		postedDay)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		sysuuid(),
		categoryId,
		l.Title,
		l.Url,
		true,
		time.Now().Year(),
		time.Now().Month(),
		time.Now().Day(),
	)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	ls, err := queryLinks(db, "WHERE linkId = "+strconv.FormatInt(id, 10))
	if err != nil {
		return nil, err
	}
	if len(ls) != 1 {
		return nil, fmt.Errorf("Something went wrong")
	}

	return &ls[0], nil
}

func queryLinks(db *sql.DB, whereClause string) ([]link, error) {

	query := `SELECT
		linkId,
		links.publicId,
		title,
		url,
		links.categoryId,
		categories.publicId,
		isValid,
		postedYear,
		postedMonth,
		postedDay
		FROM links
		inner join categories
		on links.categoryId = categories.categoryId`

	if whereClause != "" {
		query += " " + whereClause
	}

	result, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	links := *new([]link)

	for result.Next() {
		l, err := scanLink(result)

		if err != nil {
			return nil, err
		}

		links = append(links, *l)
	}

	return links, nil
}

func deleteLink(db *sql.DB, linkPublicId string) error {
	result, err := db.Exec("DELETE FROM links WHERE publicId = ?", linkPublicId)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil || rows == 0 {
		return fmt.Errorf("Link not found")
	}

	return nil
}

func scanLink(r *sql.Rows) (*link, error) {
	var (
		linkId           uint
		publicId         string
		title            string
		url              string
		categoryId       sql.NullInt64
		publicCategoryId string
		isValid          sql.RawBytes
		postedYear       uint
		postedMonth      uint
		postedDay        uint
	)

	err := r.Scan(
		&linkId,
		&publicId,
		&title,
		&url,
		&categoryId,
		&publicCategoryId,
		&isValid,
		&postedYear,
		&postedMonth,
		&postedDay,
	)

	if err != nil {
		return nil, err
	}

	l := link{
		linkId,
		publicId,
		title,
		url,
		uint(categoryId.Int64),
		publicCategoryId,
		isValid[0] != 0x00,
		date{
			postedYear,
			postedMonth,
			postedDay,
		},
	}

	return &l, nil
}

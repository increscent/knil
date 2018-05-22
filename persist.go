package main

import (
	"database/sql"
	"log"
)
import _ "github.com/go-sql-driver/mysql"

func connect() *sql.DB {
	db, err := sql.Open("mysql", "robert:robert@/knil")

	if err != nil {
		log.Fatal(err)
	}

	return db
}

package main

import (
	"fmt"
	"log"
	"net/http"
)
import _ "github.com/go-sql-driver/mysql"

func main() {
	fmt.Println("Hello world!")

	db := connect()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	result, err := tx.Exec("INSERT INTO links (title, url, isValid, postedDate) VALUES (?, ?, ?, ?)",
		"Ask HN: What is your favourite tech talk? | Hacker News",
		"https://news.ycombinator.com/item?id=16838460",
		true,
		"2018-05-17",
	)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	rows, err := result.RowsAffected()
	fmt.Println(rows)

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	db.Close()

	http.HandleFunc("/test", myHandler)

	http.ListenAndServe(":8080", nil)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	fmt.Println(r)
	fmt.Fprint(w, "Hello world! 2")
}

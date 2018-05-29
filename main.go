package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// fmt.Println("Hello world!")
	//
	// db := connect()
	//
	// if err := db.Ping(); err != nil {
	// 	log.Fatal(err)
	// }
	//
	// tx, err := db.Begin()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// result, err := tx.Exec("INSERT INTO links (publicId, categoryId, title, url, isValid, postedEpoch) VALUES (?, ?, ?, ?, ?, ?)",
	// 	uuid(),
	// 	1,
	// 	"Ask HN: What is your favourite tech talk? | Hacker News",
	// 	"https://news.ycombinator.com/item?id=16838460",
	// 	true,
	// 	time.Now().Unix(),
	// )
	// if err != nil {
	// 	tx.Rollback()
	// 	log.Fatal("Insert Error: ", err)
	// }
	//
	// rows, err := result.RowsAffected()
	// fmt.Println(rows)
	//
	// err = tx.Commit()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// links, err := queryLinks(db)
	//
	// fmt.Println(links)
	//
	// db.Close()

	http.HandleFunc("/links/", linksHandler(connect()))

	http.ListenAndServe(":8080", nil)
}

package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := connect()

	http.HandleFunc("/links", linksHandler(db))

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

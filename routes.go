package main

import (
	"net/http"
	"strings"
	"fmt"
	"database/sql"
	"encoding/json"
)

func linksHandler(db *sql.DB) (func (w http.ResponseWriter, r *http.Request)) {
	return func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.String()
		folders := strings.Split(url, "/")
		linkId := folders[len(folders) - 1]

		fmt.Println(linkId)

		switch r.Method {
		case "POST":
			var l link
			b := make([]byte, r.ContentLength)
			r.Body.Read(b)
			json.Unmarshal(b, &l)
			fmt.Println(b)
			fmt.Println(l)
		case "GET":
			links, _ := queryLinks(db)
			fmt.Println(links)
			b, _ := json.Marshal(links)
			w.Write(b)
		}
	}
}
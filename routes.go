package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type httpError struct {
	Error string
	code  int
}

func linksHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.String()
		folders := strings.Split(url, "/")
		linkId := folders[len(folders)-1]

		fmt.Println(linkId)

		var handler func(db *sql.DB, r *http.Request) (interface{}, *httpError)

		switch r.Method {
		case "POST":
			handler = linksPost
		case "GET":
			handler = linksGet
		default:
			handler = func(db *sql.DB, r *http.Request) (interface{}, *httpError) {
				return nil, &httpError{"Invalid request method. Use GET, POST, or DELETE", 405}
			}
		}

		d, err := handler(db, r)
		if err != nil {
			w.WriteHeader(err.code)
			d = err
		}
		b, _ := json.Marshal(d)
		w.Write(b)
	}
}

func linksPost(db *sql.DB, r *http.Request) (interface{}, *httpError) {
	b := make([]byte, r.ContentLength)
	r.Body.Read(b)
	var l link
	json.Unmarshal(b, &l)

	resp, err := http.Get(l.Url)
	if err != nil || resp.StatusCode < 200 || resp.StatusCode > 200 {
		return nil, &httpError{"Invalid link :(", 400}
	}
	if strings.TrimSpace(l.Title) == "" {
		return nil, &httpError{"Missing title :(", 400}
	}
	if strings.TrimSpace(l.CategoryPublicId) == "" {
		return nil, &httpError{"Missing category :(", 400}
	}

	newLink, err := insertLink(db, l)

	if err != nil {
		return nil, &httpError{err.Error(), 500}
	}

	return newLink, nil
}

func linksGet(db *sql.DB, r *http.Request) (interface{}, *httpError) {
	l, err := queryLinks(db, "")
	if err != nil {
		return nil, &httpError{err.Error(), 500}
	}
	return l, nil
}

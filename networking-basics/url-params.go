package main

import (
	"net/http"
	"io"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		key := "name"
		name := "Name: " + req.URL.Query().Get(key)
		io.WriteString(res, name)
	})
	http.ListenAndServe(":8080", nil)
}
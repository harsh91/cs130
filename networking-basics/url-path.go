package main

import (
	"net/http"
	"io"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "URL Path: "+req.URL.Path)
	})
	http.ListenAndServe(":8080", nil)
}
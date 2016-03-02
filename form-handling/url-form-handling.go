package main

import (
	"net/http"
	"io"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		key := "name"
		value := req.FormValue(key)
	 	res.Header().Set("Content-Type", "text/html; charset=utf-8")
		if value != "" {
			value = "You entered name as : " + value
			io.WriteString(res, value)
		} else {
			io.WriteString(res,"No Name for now try to put some name in url !!")
		}
	})
	http.ListenAndServe(":8080", nil)
}
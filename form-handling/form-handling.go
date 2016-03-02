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
			value = "Name we are getting : " + value;
			io.WriteString(res, value);
		} else {
			io.WriteString(res,
				`<form method = "GET">
				<input type = "text" name = "name"/>
				<input type = "submit"/>
			</form>`)
		}
	})
	http.ListenAndServe(":8080", nil)
}
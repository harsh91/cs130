package main

import(
	"fmt"
	"crypto/hmac"
	"crypto/sha256"
	"net/http"
	"html/template"
	"strings"
	"io"
)

var htmlCode *template.Template

func main(){
	http.HandleFunc("/", cookieStuff)
	http.HandleFunc("/check-hmac/", checkHmac)
	http.ListenAndServe(":8080", nil)
}

func cookieStuff(res http.ResponseWriter, req * http.Request){
	cookie, err := req.Cookie("user-info")
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name: "user-info",
			HttpOnly: true,
		}
	}
	if(req.FormValue("check") != "") {
		check := req.FormValue("check")
		cookie.Value = check + "|" + getHmac(check)
		http.SetCookie(res, cookie)
	}
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `
		<form method="POST" action="/">
			<input type="text" name="check"/>
			<input type="submit" value="submit">
		</form>
		`)
	if req.Method == "POST" {
		io.WriteString(res, `<form action="/check-hmac">
				<input type="submit" value="Check Hmac">
			</form>`)
	}
}

func checkHmac(res http.ResponseWriter, req * http.Request){
	cookie, err := req.Cookie("user-info")
	if(err == http.ErrNoCookie){
		io.WriteString(res, "No Cookie set !!")
		return
	}
	cookieValue := strings.Split((cookie.Value), "|")
	newHmac := getHmac(cookieValue[0])
	if(cookieValue[1] != newHmac){
		io.WriteString(res, "Cookie is not matching !!")
		return
	}
	io.WriteString(res, "Everything seems fine. Carry on..")
}

func getHmac(data string) string{
	h := hmac.New(sha256.New, []byte("key-aka-salt"))
	return fmt.Sprintf("%x", h.Sum(nil))
}
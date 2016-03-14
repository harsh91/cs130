package main

import(
	"net/http"
	"github.com/nu7hatch/gouuid"
	"io"
)

func main(){
	http.HandleFunc("/", UUIDCookie)
	http.ListenAndServe(":8080", nil)
}

func UUIDCookie(res http.ResponseWriter, req * http.Request){
	cookie, err := req.Cookie("UUID")
	if(err == http.ErrNoCookie){
		uuid, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name: "UUID",
			Value: uuid.String(),
			HttpOnly: true,
//			Secure: true,
		}
		http.SetCookie(res, cookie)
	}
	userCookie := "Cookie UUID : " + cookie.Value
	io.WriteString(res, userCookie)
}
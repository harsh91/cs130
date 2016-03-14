package main
import (
	"net/http"
	"html/template"
	"log"
	"github.com/nu7hatch/gouuid"
	"encoding/json"
	"encoding/base64"
	"strconv"
	"io"
)

var template1 *template.Template
type User struct {
	Name string
	Age int
}
func main() {
	var err error
	template1, err = template.ParseFiles("template.html")
	if(err != nil){
		log.Println("Error: ", err)
	}
	http.HandleFunc("/", home)
	http.HandleFunc("/user", user)
	http.ListenAndServe(":8080", nil)
}

func home(res http.ResponseWriter, req *http.Request) {
	cookie, err1 := req.Cookie("session-fino")
	if(err1 == http.ErrNoCookie){
		uuid, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name: "session-fino",
			Value: uuid.String(),
			HttpOnly: true,
//			Secure: true,
		}
		http.SetCookie(res, cookie)
	}

	var err error
	err = template1.Execute(res,nil)
	if(err != nil){
		log.Println("Error: ", err)
	}
}

func user(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		name := req.FormValue("name")
		age, errStringConvAge := strconv.Atoi(req.FormValue("age"))
		if errStringConvAge != nil {
			log.Println("Error: ", errStringConvAge)
		}
		user := User {
			Name: name,
			Age: age,
		}
		jsonUser, errJsonMarshalError := json.Marshal(user)
		if errJsonMarshalError != nil {
			log.Println("Error: ", errJsonMarshalError)
		}
		b64User := base64.StdEncoding.EncodeToString([]byte(jsonUser))
		cookie, err1 := req.Cookie("user-info")
		if(err1 == http.ErrNoCookie) {
			cookie = &http.Cookie{
				Name: "user-info",
				Value: b64User,
				HttpOnly: true,
				//			Secure: true,
			}
			http.SetCookie(res, cookie)
		}
	}
}
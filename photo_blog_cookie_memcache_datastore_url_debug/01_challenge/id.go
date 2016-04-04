package mem

import (
	"errors"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
)
/**
gets the id from cookie firstly,
if there is no cookie then searches in the url,
if no id whatsoever then generates a new id and redirects to logout page.
if there exists an id then store it in cookie or you can say in some cases updated the old one
 */
func getAvailableId(res http.ResponseWriter, req *http.Request) (string, error) {
	ctx := appengine.NewContext(req)
	var id, origin string
	var cookie *http.Cookie
	// getting id from cookie firstly
	origin = "COOKIE"
	cookie, err := req.Cookie("session-id")
	if err == http.ErrNoCookie {
		//no cookie set try to get id from url
		origin = "URL"
		id := req.FormValue("id")
		if id == "" {
			//no id whatsoever lets generate a new one and redirect to logout page
			//tracking our origin for id
			origin = "BRAND NEW VIA LOGOUT"
			log.Infof(ctx, "ID CAME FROM: %s", origin)
			//using the status code names is a good habit giving a readability of code
			http.Redirect(res, req, "/logout", http.StatusSeeOther)
			return id, errors.New("ERROR: redirect to /logout because no session id accessible")
		}
		// update the cookie storing id for later usages
		cookie = &http.Cookie{
			Name:  "session-id",
			Value: id,
			//for development usages secure has been turned off
			// Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
	}
	id = cookie.Value
	//checking out the origin from where our cookie is generated
	log.Infof(ctx, "ID CAME FROM %s", origin)
	return id, nil
}
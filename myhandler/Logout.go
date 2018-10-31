package myhandler

import (
	"net/http"
)

// Logout is Handler that remove session_id and redirect "/home"
func Logout(w http.ResponseWriter, r *http.Request) {
	endOfSession(r)
	http.Redirect(w, r, "/", 303)
}

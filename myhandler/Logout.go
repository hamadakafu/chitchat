package myhandler

import (
	"database/sql"
	"fmt"
	"net/http"
)

// Logout is Handler that remove session_id and redirect "/home"
func Logout(w http.ResponseWriter, r *http.Request) {
	endOfSession(r)
	http.Redirect(w, r, "/home", 303)
}

// endOfSession is Func that delete session_id and change session_state to false
func endOfSession(r *http.Request) {
	sessionID, err := r.Cookie("_cookie")
	db, err := sql.Open("postgres", sqlLoginWord)
	if err != nil {
		fmt.Println("cant open postgres in endOfSession!!!", err)
	}
	_, err = db.Exec(fmt.Sprintf("update userinfo set session_id = null, session_state = 'false' where session_id = '%s'", sessionID.Value))
	if err != nil {
		fmt.Println("cant exec in endOfSession!!!", err)
	}
}

package myhandler

import (
	"fmt"
	"html/template"
	"net/http"
)

// CreateChat is Handler that shows createchat.html
func CreateChat(w http.ResponseWriter, r *http.Request) {
	if isInSession, _ := sessionCheck(r); isInSession {

		t, err := template.ParseFiles("layout.html", "createChat.html")
		if err != nil {
			fmt.Println("template.ParseFiles error in createChat Handler!!!", err)
		}
		t.ExecuteTemplate(w, "layout", nil)
	} else {
		t, err := template.ParseFiles("lsyout.html", "login.html")
		if err != nil {
			fmt.Println("template ParseFiles error in CreateChat Handler!!!", err)
		}
		t.ExecuteTemplate(w, "layout", nil)
	}
}

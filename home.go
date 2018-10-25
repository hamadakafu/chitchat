package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if ok, name := isInSession(r); ok { // Do you have cookie?
		t, err := template.ParseFiles("layout.html", "chitchat.html")
		if err != nil {
			fmt.Println("template err in login func", err)
		}
		t.ExecuteTemplate(w, "layout", name)
		fmt.Println("execute isInSession")
	} else if existUser(username, password) {
		makeCookie(username, password, w)
		t, err := template.ParseFiles("layout.html", "chitchat.html")
		if err != nil {
			fmt.Println("template err in login func", err)
		}
		t.ExecuteTemplate(w, "layout", username)
		fmt.Println("excute existUser")
	} else {
		t, err := template.ParseFiles("layout.html", "login.html")
		if err != nil {
			fmt.Println("template err in login func")
			fmt.Println(err)
		}
		t.ExecuteTemplate(w, "layout", "")
		fmt.Println("excute nothing")
	}
}

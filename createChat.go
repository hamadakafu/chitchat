package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func createChat(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("layout.html", "createChat.html")
	if err != nil {
		fmt.Println("template.ParseFiles error in createChat Handler!!!", err)
	}
	t.ExecuteTemplate(w, "layout", nil)
}

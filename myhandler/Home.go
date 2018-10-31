package myhandler

import (
	"fmt"
	"html/template"
	"net/http"
)

// Home is Handler that show chitchat.html
func Home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("layout.html", "chitchat.html")
	if err != nil {
		fmt.Println("template err in login func", err)
	}

	data := Data{}
	makeData(&data, r)
	t.ExecuteTemplate(w, "layout", data)
}

package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func successOfCreateChat(w http.ResponseWriter, r *http.Request) {
	user := toUserFromRequest(r)
	err := addChatList(&user, r.FormValue("title"))
	if err != nil {
		fmt.Println(err)
	}
	t, err := template.ParseFiles("layout.html", "successOfCreateChat.html")
	if err != nil {
		fmt.Println("template.ParseFiles error in successOfCreateChat Handler!!!", err)
	}
	if err := t.ExecuteTemplate(w, "layout", nil); err != nil {
		fmt.Println("template.ExecuteTemplate error in successOfCreateChat", err)
	}
}

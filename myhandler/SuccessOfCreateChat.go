package myhandler

import (
	"fmt"
	"html/template"
	"net/http"
)

// SuccessOfCreateChat is Handler that create chat in DB and redirect home.html
func SuccessOfCreateChat(w http.ResponseWriter, r *http.Request) {
	if isInSession, _ := sessionCheck(r); isInSession {
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
	} else {
		t, err := template.ParseFiles("layout.html", "login.html")
		if err != nil {
			fmt.Println("Parsefiles error in successofcreatechat!!!", err)
		}
		t.ExecuteTemplate(w, "layout", nil)
	}
}

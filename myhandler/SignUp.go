package myhandler

import (
	"fmt"
	"html/template"
	"net/http"
)

// SignUp is Handler that show signup.html and dont need session
func SignUp(w http.ResponseWriter, r *http.Request) {
	if allNotEmptyForm(r, "username", "password") && !existUser(r.FormValue("username")) {
		registerUser(r)
		t, err := template.ParseFiles("layout.html", "login.html")
		if err != nil {
			fmt.Println("template ParseFiles error in SignUp!!!", err)
		}
		t.ExecuteTemplate(w, "layout", nil)
		fmt.Println("excute login")
	} else {
		t, err := template.ParseFiles("layout.html", "signup.html")
		if err != nil {
			fmt.Println("template ParseFiles error in SignUp!!!", err)
		}
		t.ExecuteTemplate(w, "layout", nil)
		fmt.Println("excute signup")
	}
}

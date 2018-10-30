package myhandler

import (
	"fmt"
	"html/template"
	"net/http"
)

// Login is Handler that return login.html
func Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		t, err := template.ParseFiles("layout.html", "login.html")
		if err != nil {
			fmt.Println("catn parseFiles in GET in Home Handler!!!", err)
		}
		t.ExecuteTemplate(w, "layout", nil)
	case "POST":
		usernameOfForm := r.FormValue("username")
		passwordOfForm := r.FormValue("password")
		if passwordCheck(usernameOfForm, passwordOfForm) {
			makeCookie(usernameOfForm, passwordOfForm, w)
			t, err := template.ParseFiles("layout.html", "chitchat.html")
			if err != nil {
				fmt.Println("template err in login func", err)
			}
			data := makeChatListData(usernameOfForm)
			t.ExecuteTemplate(w, "layout", chatListData)
			fmt.Println("excute existUser")
		}
	default:
	}

	// if isInsession, username := sessionCheck(r); isInsession { // Do you have cookie?
	// 	t, err := template.ParseFiles("layout.html", "chitchat.html")
	// 	if err != nil {
	// 		fmt.Println("template err in login func", err)
	// 	}
	// 	chatListData := makeChatListData(username)
	// 	if err := t.ExecuteTemplate(w, "layout", chatListData); err != nil {
	// 		fmt.Println("ExecuteTemplate error in home func", err)
	// 	}
	// 	fmt.Println("execute isInSession")
	// } else if passwordCheck(usernameOfForm, passwordOfForm) {
	// 	makeCookie(usernameOfForm, passwordOfForm, w)
	// 	t, err := template.ParseFiles("layout.html", "chitchat.html")
	// 	if err != nil {
	// 		fmt.Println("template err in login func", err)
	// 	}
	// 	chatListData := makeChatListData(username)
	// 	t.ExecuteTemplate(w, "layout", chatListData)
	// 	fmt.Println("excute existUser")
	// } else {
	// 	t, err := template.ParseFiles("layout.html", "login.html")
	// 	if err != nil {
	// 		fmt.Println("template err in login func")
	// 		fmt.Println(err)
	// 	}
	// 	t.ExecuteTemplate(w, "layout", "")
	// 	fmt.Println("excute nothing")
	// }
}

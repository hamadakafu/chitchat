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
		data := Data{}
		makeData(&data, "")
		t.ExecuteTemplate(w, "layout", data)
	case "POST":
		usernameOfForm := r.FormValue("username")
		passwordOfForm := r.FormValue("password")
		if passwordCheck(usernameOfForm, passwordOfForm) {
			makeCookie(usernameOfForm, passwordOfForm, w)
			http.Redirect(w, r, "/home", 303)
			fmt.Println("excute existUser")
		} else {
			t, err := template.ParseFiles("layout.html", "login.html")
			if err != nil {
				fmt.Println("parseFiles error in Login Handler!!!", err)
			}
			data := Data{}
			errorTitle := "invalidForm"
			errorMessage := "Username or password is invalid. Please confirm."
			makeErrorData(&data, errorTitle, errorMessage)
			t.ExecuteTemplate(w, "layout", data)
			fmt.Println("execute error formvalue")
			fmt.Println(data.UserError)
		}
	default:
	}
}

func makeErrorData(data *Data, errorTitle string, errorMessage string) {
	data.UserError = UserError{
		ErrorTitle:   errorTitle,
		ErrorMessage: errorMessage,
	}
}

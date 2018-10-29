package myhandler

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

// SignUp is Handler that show signup.html
func SignUp(w http.ResponseWriter, r *http.Request) {
	if allOkForm(r, "username", "password") && !existUser(r.FormValue("username")) {
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

func okForm(r *http.Request, formName string) bool {
	formValue := r.FormValue(formName)
	if formValue == "" {
		fmt.Println("This request is empty username!!!")
		return false
	} else {
		return true
	}
}

func allOkForm(r *http.Request, formNames ...string) bool {
	for _, formName := range formNames {
		if !okForm(r, formName) {
			return false
		}
	}
	return true
}

func existUser(username string) bool {
	db, err := sql.Open("postgres", sqlLoginWord)
	if err != nil {
		fmt.Println("cant open postgres in existUser!!!")
	}
	rows, err := db.Query(fmt.Sprintf(
		"select * from userinfo where user_name = '%s'", username))
	if err != nil {
		fmt.Println("Query is invalid in existUser!!!", err)
	}
	defer rows.Close()
	if rows.Next() {
		return true
	} else {
		return false
	}
}
func registerUser(r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	numUser := getNumberOfUser()
	date := time.Now().Format("2006-01-02")
	db, err := sql.Open("postgres", sqlLoginWord)
	if err != nil {
		fmt.Println("cant open postgres in registerInfo!!!", err)
	}
	_, err = db.Exec(fmt.Sprintf("begin transaction"))
	if err != nil {
		fmt.Println("cant Exec begin transaction", err)
	}
	_, err = db.Exec(fmt.Sprintf("insert into userinfo values(%d, '%s', '%s', '%s', 'false', null)", numUser, username, password, date))
	if err != nil {
		fmt.Println("cant Exec insert registerUser!!!", err)
	}
	_, err = db.Exec(fmt.Sprintf("update chitchatinfo set number_of_user = number_of_user + 1"))
	if err != nil {
		fmt.Println("cant update in resisterUser!!!")
	}
	_, err = db.Exec(fmt.Sprintf("commit"))
	if err != nil {
		fmt.Println("cant commit in registerUser!!!", err)
	}
}

func getNumberOfUser() (num int) {
	db, err := sql.Open("postgres", sqlLoginWord)
	if err != nil {
		fmt.Println("cant open postgres in getNumberOfUser!!!", err)
	}
	rows, err := db.Query(fmt.Sprintf("select number_of_user from chitchatinfo"))
	if err != nil {
		fmt.Println("Query is invalid in getNumberOfUser!!!", err)
	}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&num)
	} else {
		fmt.Println("cant read number of users, lists")
	}
	return
}

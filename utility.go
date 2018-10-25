package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"net/http"
	"os"
)

const rs = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = rs[rand.Intn(len(rs))]
	}
	return string(b)
}

func makeCookie(username string, password string, w http.ResponseWriter) {
	db, err := sql.Open("postgres",
		"user=kafuhamada password=wd dbname=chitchat sslmode=disable")
	if err != nil {
		fmt.Println("cant open postgres!!!", err)
	}
	sessionID := randString(255)
	_, err = db.Exec(fmt.Sprintf("update userinfo set session_state = 'true', session_id = '%s'", sessionID))
	if err != nil {
		fmt.Println("cant update session_state and session_id", err)
	}
	cookie := http.Cookie{
		Name:     "_cookie",
		Value:    sessionID,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
}
func existUser(username string, password string) bool {
	db, err := sql.Open("postgres",
		"user=kafuhamada password=wd dbname=chitchat sslmode=disable")
	if err != nil {
		fmt.Println("cant open postgres!!!", err)
	}
	rows, err := db.Query(fmt.Sprintf("select * from userinfo where user_name = '%s' and user_password = '%s'", username, password))

	if err != nil {
		fmt.Println(err)
	}
	if rows.Next() {
		return true
	}
	return false
}

func isInSession(r *http.Request) (bool, string) {
	cookie, err := r.Cookie("_cookie")
	if err == nil { // cookie exist
		return sessionCheck(cookie)
	}
	return false, ""
}

// User is user information...
type User struct {
	ID           int
	Name         string
	Password     string
	CreateDate   string
	SessionState bool
	SessionID    string
}

func sessionCheck(cookie *http.Cookie) (bool, string) {
	db, err := sql.Open("postgres",
		"user=kafuhamada password=wd dbname=chitchat sslmode=disable")
	if err != nil {
		fmt.Println("cant open db!!", err)
		os.Exit(1)
	}
	rows, err := db.Query(fmt.Sprintf("select user_id, user_name, user_password, create_date, session_state, session_id from userinfo where session_id = '%s'", cookie.Value))
	if err != nil {
		fmt.Println("This query is invalid!!", err)
	}
	defer rows.Close()

	user := User{}
	if rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Password,
			&user.CreateDate, &user.SessionState, &user.SessionID)
		if err != nil {
			fmt.Println("cant scan rows in db!!", err)
		}

		if cookie.Value == user.SessionID && user.SessionState {
			return true, user.Name
		}
		return false, ""
	}
	return false, ""
}

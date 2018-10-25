package main

import (
	"crypto/sha256"
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
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
		"user=root password=wd dbname=chitchat sslmode=disable")
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
		"user=root password=wd dbname=chitchat sslmode=disable")
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
		"user=root password=wd dbname=chitchat sslmode=disable")
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

func addChatList(user *User, title string) error {
	date := time.Now().Format("2006-01-02")
	hash := sha256.Sum256([]byte(title))
	db, err := sql.Open("postgres", "user=root password=wd dbname=chitchat sslmode=disable")
	if err != nil {
		fmt.Println("cant open postgres in addChatList!!!", err)
	}

	rows, err := db.Query(fmt.Sprintf("select * from chatlist where chat_name = '%s'", title))
	if err != nil {
		fmt.Println("query is invalid in addChatList!!!", err)
	}
	defer rows.Close()
	if rows.Next() {
		rows.Scan()
		return errors.New("chat title already exist")
	}
	_, err = db.Exec("begin transaction")
	if err != nil {
		fmt.Println("exec is invalid in addChatList!!!", err)
	}
	_, err = db.Exec(fmt.Sprintf("insert into chatlist values(%d, '%s', '%s', '%s', '%s')",
		user.ID,
		user.Name,
		date,
		fmt.Sprintf("%x", hash[:]),
		title))
	if err != nil {
		fmt.Println("exec is invalid in addChatList!!!", err)
	}
	_, err = db.Exec("commit")
	if err != nil {
		fmt.Println("exec is invalid in addChatList!!!", err)
	}
	return nil
}
func toUserFromRequest(r *http.Request) (user User) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		fmt.Println("cookie err!!! in toUserFromRequest", err)
	}
	db, err := sql.Open("postgres", "user=root password=wd dbname=chitchat sslmode=disable")
	if err != nil {
		fmt.Println("cant open postgres in toUserFromRequest!!!", err)
	}
	rows, err := db.Query(fmt.Sprintf("select * from userinfo where session_id = '%s'", cookie.Value))
	if err != nil {
		fmt.Println("Query is invalid in toUserFromRequest!!!", err)
	}
	if rows.Next() {
		rows.Scan(
			&user.ID,
			&user.Name,
			&user.Password,
			&user.CreateDate,
			&user.SessionState,
			&user.SessionID,
		)
	}
	defer rows.Close()
	return
}

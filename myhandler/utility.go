package myhandler

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

func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = rs[rand.Intn(len(rs))]
	}
	return string(b)
}

func makeCookie(username string, password string, w http.ResponseWriter) {
	db, err := sql.Open("postgres",
		"user=chitchatmanager password=wd dbname=chitchat sslmode=disable")
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
func passwordCheck(username string, password string) bool {
	db, err := sql.Open("postgres",
		"user=chitchatmanager password=wd dbname=chitchat sslmode=disable")
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

func sessionCheck(r *http.Request) (bool, string) {
	cookie, err := r.Cookie("_cookie")
	if err != nil { // cookie exist
		return false, ""
	}
	db, err := sql.Open("postgres",
		"user=chitchatmanager password=wd dbname=chitchat sslmode=disable")
	if err != nil {
		fmt.Println("cant open db!!", err)
		os.Exit(1)
	}
	rows, err := db.Query(fmt.Sprintf("select user_id, user_name, user_password, create_date, session_state, session_id from userinfo where session_id = '%s'", cookie.Value))
	if err != nil {
		fmt.Println("This query is invalid!!", err)
	}
	defer rows.Close()

	user := UserInfo{}
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

func addChatList(user *UserInfo, title string) error {
	date := time.Now().Format("2006-01-02")
	hash := sha256.Sum256([]byte(title))
	db, err := sql.Open("postgres", "user=chitchatmanager password=wd dbname=chitchat sslmode=disable")
	if err != nil {
		fmt.Println("cant open postgres in addChatList!!!", err)
	}

	rows, err := db.Query(fmt.Sprintf("select * from chatlist where chat_title = '%s'", title))
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
	_, err = db.Exec(
		fmt.Sprintf("insert into chatlist values(%d, '%s', '%s', '%s', '%s', 0)",
			user.ID,
			user.Name,
			date,
			fmt.Sprintf("%x", hash[:]),
			title),
	)
	if err != nil {
		fmt.Println("exec is invalid in addChatList!!!", err)
	}
	_, err = db.Exec("commit")
	if err != nil {
		fmt.Println("exec is invalid in addChatList!!!", err)
	}
	return nil
}
func toUserFromRequest(r *http.Request) (user UserInfo) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		fmt.Println("cookie err!!! in toUserFromRequest", err)
	}
	db, err := sql.Open("postgres", "user=chitchatmanager password=wd dbname=chitchat sslmode=disable")
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

// should be exist data in DB.
func makeData(data *Data, r *http.Request) {
	db, err := sql.Open("postgres", "user=kafuhamada password=pw dbname=chitchat sslmode=disable")
	if err != nil {
		fmt.Println("cant open postgres!!!", err)
	}
	rows, err := db.Query(fmt.Sprintf("select * from chatlist"))
	if err != nil {
		fmt.Println("this Query is invalid!!!", err)
	}
	defer rows.Close()
	var (
		user        UserInfo
		chatList    ChatList
		commentList CommentList
		userError   UserError
	)
	user.getData(r)
	chatList.getData()
	commentList.getData("")
	data.User = user
	data.ChatList = chatList
	data.CommentList = commentList
	data.UserError = userError
}

func notEmptyForm(r *http.Request, formName string) bool {
	formValue := r.FormValue(formName)
	if formValue == "" {
		fmt.Println("This request is empty username!!!")
		return false
	} else {
		return true
	}
}

func allNotEmptyForm(r *http.Request, formNames ...string) bool {
	for _, formName := range formNames {
		if !notEmptyForm(r, formName) {
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

// endOfSession is Func that delete session_id and change session_state to false
func endOfSession(r *http.Request) {
	sessionID, err := r.Cookie("_cookie")
	db, err := sql.Open("postgres", sqlLoginWord)
	if err != nil {
		fmt.Println("cant open postgres in endOfSession!!!", err)
	}
	_, err = db.Exec(fmt.Sprintf("update userinfo set session_id = null, session_state = 'false' where session_id = '%s'", sessionID.Value))
	if err != nil {
		fmt.Println("cant exec in endOfSession!!!", err)
	}
}

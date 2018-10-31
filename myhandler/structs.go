package myhandler

import (
	"database/sql"
	"fmt"
)

// DBData is interface that implement db access
type DBData interface {
	getData(string)
}

// Data is struct that is renderd data with template.
type Data struct {
	User        UserInfo
	ChatList    ChatList
	CommentList CommentList
	UserError   UserError
}

// ChatInfo is infomation of some chat
type ChatInfo struct {
	CreateUserID    int
	CreateUserName  string
	CreateDate      string
	ChatHash        string
	ChatTitle       string
	NumberOfComment int
}

// ChatList is []ChatInfo type alias
type ChatList []ChatInfo

// CommentInfo is struct that have comment infomation.
type CommentInfo struct {
	CommentID      int
	CommentText    string
	CreateUserID   int
	CreateUserName string
	CreateDate     string
	ChatTitle      string
}

// CommentList is []CommentInfo type alias
type CommentList []CommentInfo

// UserError is sturct that show error to user
type UserError struct {
	ErrorTitle   string
	ErrorMessage string
}

// UserInfo is user information...
type UserInfo struct {
	ID           int
	Name         string
	Password     string
	CreateDate   string
	SessionState bool
	SessionID    string
}

func (u *UserInfo) getData(username string) {
	if username == "" {
		return
	}
	db, err := sql.Open("postgres", sqlLoginWord)
	if err != nil {
		fmt.Println("cant open DB in getUserInfo!!!", err)
	}
	rows, err := db.Query(
		fmt.Sprintf("select * from userinfo where user_name = '%s'", username),
	)
	if err != nil {
		fmt.Println("Query is invalid in getUserInfo!!!", err)
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(
			&u.ID,
			&u.Name,
			&u.Password,
			&u.CreateDate,
			&u.SessionState,
			&u.SessionID,
		)
	} else {
		fmt.Println("cant scan rows in getUserInfo!!!")
	}
}

func (cl ChatList) getData() {
	db, err := sql.Open("postgres", sqlLoginWord)
	if err != nil {
		fmt.Println("catn open DB in getData of chatInfo!!!", err)
	}
	rows, err := db.Query(
		fmt.Sprintf("select * from chatlist"),
	)
	if err != nil {
		fmt.Println("Query is invalid in getData of chatInfo")
	}
	defer rows.Close()
	chatInfo := ChatInfo{}
	for rows.Next() {
		rows.Scan(
			&chatInfo.CreateUserID,
			&chatInfo.CreateUserName,
			&chatInfo.CreateDate,
			&chatInfo.ChatHash,
			&chatInfo.ChatTitle,
			&chatInfo.NumberOfComment,
		)
		cl = append(cl, chatInfo)
	}
}

func (cml CommentList) getData(chatTitle string) {
	if chatTitle == "" {
		return
	}
	db, err := sql.Open("postgrs", sqlLoginWord)
	if err != nil {
		fmt.Println("cant open DB in getData of CommentList!!!", err)
	}
	rows, err := db.Query(
		fmt.Sprintf("select * from comments where chat_title = '%s'", chatTitle),
	)
	if err != nil {
		fmt.Println("Query is invalid in getData of CommentList")
	}
	defer rows.Close()

	cm := CommentInfo{}
	for rows.Next() {
		rows.Scan(
			&cm.CommentID,
			&cm.CommentText,
			&cm.CreateUserID,
			&cm.CreateUserName,
			&cm.CreateDate,
			&cm.ChatTitle,
		)
		cml = append(cml, cm)
	}
}

const rs string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

const sqlLoginWord string = "user=chitchatmanager password=wd dbname=chitchat sslmode=disable"

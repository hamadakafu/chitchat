package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
)

func makeChatListData(username string) ChatListData {
	db, err := sql.Open("postgres", "user=kafuhamada password=pw dbname=chitchat sslmode=disable")
	if err != nil {
		fmt.Println("cant oepn postgres!!!", err)
	}
	rows, err := db.Query(fmt.Sprintf("select * from chatlist"))
	if err != nil {
		fmt.Println("this Query is invalid!!!", err)
	}
	defer rows.Close()
	chatInfo := ChatInfo{}
	chatListData := ChatListData{}
	for rows.Next() {
		rows.Scan(&chatInfo.CreateUserID, &chatInfo.CreateUserName, &chatInfo.CreateDate, &chatInfo.ChatHash, &chatInfo.ChatName)
		chatListData.ChatList = append(chatListData.ChatList, chatInfo)
	}
	chatListData.UserName = username
	fmt.Println(chatListData)
	return chatListData
}

func home(w http.ResponseWriter, r *http.Request) {
	usernameOfForm := r.FormValue("username")
	passwordOfForm := r.FormValue("password")

	if ok, username := isInSession(r); ok { // Do you have cookie?
		t, err := template.ParseFiles("layout.html", "chitchat.html")
		if err != nil {
			fmt.Println("template err in login func", err)
		}
		chatListData := makeChatListData(username)
		if err := t.ExecuteTemplate(w, "layout", chatListData); err != nil {
			fmt.Println("ExecuteTemplate error in home func", err)
		}
		fmt.Println("execute isInSession")
	} else if existUser(usernameOfForm, passwordOfForm) {
		makeCookie(usernameOfForm, passwordOfForm, w)
		t, err := template.ParseFiles("layout.html", "chitchat.html")
		if err != nil {
			fmt.Println("template err in login func", err)
		}
		chatListData := makeChatListData(username)
		t.ExecuteTemplate(w, "layout", chatListData)
		fmt.Println("excute existUser")
	} else {
		t, err := template.ParseFiles("layout.html", "login.html")
		if err != nil {
			fmt.Println("template err in login func")
			fmt.Println(err)
		}
		t.ExecuteTemplate(w, "layout", "")
		fmt.Println("excute nothing")
	}
}

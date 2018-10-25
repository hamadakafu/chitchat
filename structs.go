package main

type ChatInfo struct {
	CreateUserID   int
	CreateUserName string
	CreateDate     string
	ChatHash       string
	ChatName       string
}

type ChatListData struct {
	UserName string
	ChatList []ChatInfo
}

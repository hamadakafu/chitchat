package myhandler

// ChatInfo is infomation of some chat
type ChatInfo struct {
	CreateUserID    int
	CreateUserName  string
	CreateDate      string
	ChatHash        string
	ChatTitle       string
	NumberOfComment int
}

// ChatListData is Username and ChatList
type ChatListData struct {
	UserName string
	ChatList []ChatInfo
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

const rs = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

const sqlLoginWord string = "user=chitchatmanager password=wd dbname=chitchat sslmode=disable"

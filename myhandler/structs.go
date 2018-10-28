package myhandler

// ChatInfo is infomation of some chat
type ChatInfo struct {
	CreateUserID   int
	CreateUserName string
	CreateDate     string
	ChatHash       string
	ChatName       string
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

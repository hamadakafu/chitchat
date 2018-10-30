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

// Data is struct that is renderd data with template.
type Data struct {
	User        UserInfo
	ChatList    []ChatInfo
	CommentList []CommentInfo
	UserError   UserError
}

// CommentInfo is struct that have comment infomation.
type CommentInfo struct {
	CommentID      int
	CommentText    string
	CreateUserID   int
	CreateUserName string
	CreateDate     string
	ChatHash       string
}

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

const rs = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

const sqlLoginWord string = "user=chitchatmanager password=wd dbname=chitchat sslmode=disable"

# rule
comment text length must be less than 256.

user name length must be less than 256.

# page
## login.html
user must see login page at first.
- next
    - error
    - chatlist
- template
    - layout ... layout.html
    - content ... login.html
## chatlist.html
user success to login, then can see chatlist.
- next
    - some chat
    - error (when chat name too long)
- template
    - layout ... layout.html
    - content ... chatlist.html
## somechat.html
in some chat, user can comment.
- next
    - some chat
    - chatlist 
    - error (when comment too long)
- template
    - layout ... layout.html
    - content ... somechat.html
## error
when something cause error, handle shows error page.
- next
    - login
- template
    - layout ... layout.html
    - content ... error.html
# database 
use postgresql
## table
- chitchatinfo
    - number_of_user
        - int
        - not null
    - number_of_chat
        - int
        - not null

- chatlist
    - create_user_id 
        - int
        - not null
    - create_user_name
        - varchar[256] 
        - not null
    - create_date
        - date 
        - not null
        - **ex)** '2018-10-01'
    - chat_hash
        - varchar[256] 
        - not null
        - primary key
    - chat_title
        - varchat[256]
        - not null
    - number_of_comment
        - int
        - not null

- comments
    - comment_id
        - int
        - not null
        - primary key
    - comment_text
        - varchar[256] 
        - not null
    - create_user_id
        - int
        - not null
    - create_user_name
        - varchar[256]
        - not null
    - create_date
        - date
        - not null
    - chat_title
        - varchar[256] 
        - not null
- userinfo
    - user_id
        - int
        - not null
    - user_name
        - varchar[256]
        - not null
    - user_password
        - varchar[256]
        - not null
    - create_date
        - date
        - not null
    - session_state
        - bool
        - not null
    - session_id
        - varchar[256]

# function
## func sessionCheck() bool
sessionCheck check user session_state is valid, then return bool.

In Handler, 

If is valid, Handler check session_id and transfer user to chatlist. 

Else, Handler check session_id. If it exist as string, it means that in past session is interrupted.
## fucn makeCookie(username string, password string, w http.ResponseWriter) 
type http.Cookie struct {

    Name: "_cookie",

    Value: sessionID,

    HttpOnly: true,

}

# cookie and session
user have cookie of session_id as random string.

If cookie of session_id user have is same to session_id server have in postgres, user can login automaticaly.

If User don't has Cookie of session_id, there is something wrong, for expample User close browser without doing logout.

# Rendered Data
Variable name is generally "data".

```Go
type Data struct {
    User: User
    ChatList: []ChatInfo
    CommentList: []CommentInfo
    UserError: UserError 
}
```
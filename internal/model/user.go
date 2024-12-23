package model

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"` // 使用 "-" 避免在JSON中返回密码
}

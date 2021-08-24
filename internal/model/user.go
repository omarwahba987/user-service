package model

type User struct {
	UserName string `json:"user_name"`
	Age      int	`json:"age"`
}

var UserList []User

package model

import "time"

type ToDo struct {
	UserName string    `json:"user_name"`
	Task     string    `json:"task"`
	Date     time.Time `json:"date"`
}

var ToDoList []ToDo

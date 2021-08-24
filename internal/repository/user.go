package repository

import (
	"encoding/json"
	"log"
	"va_test_a/internal/model"
	"va_test_a/pkg/database"
)

type UserRepository interface {
	CreateUser(string, int) bool
	EditUser(string, int) bool
	IfExist(string) bool
}

type UserRepo struct {
}

func (u UserRepo) IfExist(name string) bool {
	db := database.DB{}
	userData := db.OpenUserFile()
	err := json.Unmarshal(userData, &model.UserList)
	if err != nil {
		log.Fatal(err)
	}
	for _, value := range model.UserList {
		if value.UserName == name {
			return true
		}
	}
	return false
}

func (u UserRepo) CreateUser(name string, age int) bool {
	db := database.DB{}
	userData := db.OpenUserFile()
	err := json.Unmarshal(userData, &model.UserList)
	if err != nil {
		log.Fatal(err)
	}

		if u.IfExist(name){
			return false
		}

	user := model.User{UserName: name, Age: age}
	model.UserList = append(model.UserList, user)
	userData, _ = json.MarshalIndent(model.UserList,"","")
	db.WriteUserFile(userData)
	return true
}

func (u UserRepo) EditUser(name string, age int) bool {
	db := database.DB{}
	userData := db.OpenUserFile()
	err := json.Unmarshal(userData, &model.UserList)
	if err != nil {
		log.Fatal(err)
	}
	for i, value := range model.UserList {
		if value.UserName == name {
			model.UserList[i].Age=age
			userData, _ = json.MarshalIndent(model.UserList,"","")
			db.WriteUserFile(userData)
			return true
		}
	}
	return false

}

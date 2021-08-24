package handler

import (
	"encoding/json"
	"net/http"
	"va_test_a/internal/model"
	"va_test_a/internal/repository"
	"va_test_a/pkg/authentication"
	"va_test_a/pkg/httpHelper"
)

func CreateUserApi(w http.ResponseWriter, req *http.Request) {
	var user model.User
		decoder := json.NewDecoder(req.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&user)
		if err != nil {
			authentication.SendError(w, "there is an error decoding")
			return
		}
		userRepo:=repository.UserRepo{}
		if userRepo.CreateUser(user.UserName,user.Age){
			helper:=httpHelper.HttpHandlerRepo{}
			helper.SendResponseMessage(w,"user added")
		}else {
			helper:=httpHelper.HttpHandlerRepo{}
			helper.SendResponseMessage(w,"user already exist")
		}
}
func EditUserApi(w http.ResponseWriter, req *http.Request) {
	var user model.User
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&user)
	if err != nil {
		authentication.SendError(w, "there is an error decoding")
		return
	}
	userRepo:=repository.UserRepo{}
	if userRepo.IfExist(user.UserName){
		if userRepo.EditUser(user.UserName,user.Age){
			helper:=httpHelper.HttpHandlerRepo{}
			helper.SendResponseMessage(w,"user modified")
		}else {
			helper:=httpHelper.HttpHandlerRepo{}
			helper.SendResponseMessage(w,"user not modified")
		}

	}else {
		helper:=httpHelper.HttpHandlerRepo{}
		helper.SendResponseMessage(w,"user not exist")
	}
}


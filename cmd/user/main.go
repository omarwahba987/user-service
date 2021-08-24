package main

import (
	"github.com/gorilla/mux"
	"va_test_a/internal/handler"
	"va_test_a/pkg/authentication"
	"va_test_a/pkg/httpHelper"
)

func main() {
	Router := mux.NewRouter()
	httpHelper:=httpHelper.HttpHandlerRepo{}
	defer httpHelper.OpenServer(Router,"8080")

	Router.Handle("/user/create",authentication.IsAuthorized(handler.CreateUserApi)).Methods("POST","OPTIONS")
	Router.Handle("/user/edit",authentication.IsAuthorized(handler.EditUserApi)).Methods("POST","OPTIONS")
}

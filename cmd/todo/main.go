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
	defer httpHelper.OpenServer(Router,"8081")
	Router.Handle("/todo/add",authentication.IsAuthorized(handler.AddTaskApi)).Methods("POST","OPTIONS")
	Router.Handle("/todo/get",authentication.IsAuthorized(handler.GetTaskApi)).Methods("POST","OPTIONS")

}

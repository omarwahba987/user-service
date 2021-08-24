package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"va_test_a/internal/model"
	"va_test_a/internal/repository"
	"va_test_a/pkg/authentication"
	"va_test_a/pkg/httpHelper"
)
type GetTasksRequest struct {
	UserName string `json:"user_name"`
}
func AddTaskApi(w http.ResponseWriter, req *http.Request) {
	var task model.ToDo
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&task)
	if err != nil {
		authentication.SendError(w, "there is an error decoding")
		fmt.Println(err)
		return
	}
	todoRepo:=repository.ToDoRepo{}

	if todoRepo.CreateTask(task.UserName,task.Task,task.Date){
		helper:=httpHelper.HttpHandlerRepo{}
		helper.SendResponseMessage(w,"task added")
	}else {
		helper:=httpHelper.HttpHandlerRepo{}
		helper.SendResponseMessage(w,"username not exist")
	}
}
func GetTaskApi(w http.ResponseWriter, req *http.Request) {
	var reqBody GetTasksRequest
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&reqBody)
	if err != nil {
		authentication.SendError(w, "there is an error decoding")

		return
	}
	todoRepo:=repository.ToDoRepo{}
	todoData:=todoRepo.GetTasks(reqBody.UserName)
	helper:=httpHelper.HttpHandlerRepo{}
	helper.SendResponseMessage(w,todoData)

}
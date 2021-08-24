package repository

import (
	"encoding/json"
	"log"
	"time"
	"va_test_a/internal/model"
	"va_test_a/pkg/database"
)

type ToDoRepository interface {
	CreateTask(string,string, time.Time) bool
	GetTasks(string) []model.ToDo
}

type ToDoRepo struct {
}

func (t ToDoRepo) CreateTask(userName string,task string, date time.Time) bool {
userRepo := UserRepo{}
if !userRepo.IfExist(userName){
	return false
}
	db := database.DB{}
	todoData := db.OpenTodoFile()
	err := json.Unmarshal(todoData, &model.ToDoList)
	if err != nil {
		log.Fatal(err)
	}
	todo := model.ToDo{UserName: userName,Task: task,Date: date}
	model.ToDoList = append(model.ToDoList, todo)
	todoData, _ = json.MarshalIndent(model.ToDoList,"","")
	db.WriteTodoFile(todoData)
	return true
}

// GetTasks  by username and its date does not come yet
func (t ToDoRepo) GetTasks(userName string)  []model.ToDo {
	var tasks []model.ToDo
	userRepo := UserRepo{}
	if !userRepo.IfExist(userName){
		return nil
	}
	db := database.DB{}
	todoData := db.OpenTodoFile()
	err := json.Unmarshal(todoData, &model.ToDoList)
	if err != nil {
		log.Fatal(err)
	}
	for _,value := range model.ToDoList{
		if value.UserName==userName&&value.Date.After(time.Now()) {
			tasks=append(tasks,value)
		}
	}
	return tasks
}


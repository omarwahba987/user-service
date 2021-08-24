package database

import (
	"fmt"
	"io/ioutil"
	"os"
)

type DB struct {
}

func (d DB) WriteUserFile(data []byte) {
	_ = ioutil.WriteFile("users.json", data, 0644)
}

func (d DB) OpenUserFile() []byte {
	jsonFile, err := os.Open("users.json")
	if err != nil {
		fmt.Println(err)
	}
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {

		}
	}(jsonFile)
	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue
}

func (d DB) WriteTodoFile(data []byte) {
	_ = ioutil.WriteFile("todo.json", data, 0644)
}
func (d DB) OpenTodoFile() []byte {
	jsonFile, err := os.Open("todo.json")
	if err != nil {
		fmt.Println(err)
	}
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {

		}
	}(jsonFile)
	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue
}

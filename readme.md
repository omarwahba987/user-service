this task number A
or run the program on command line
we have two services
first one for user run on port 8080 //cmd/user
second one for todo list run on port 8081//cmd/todo

	1- change directory to program file 
	2- open terminal
	3- type command "go mod tidy" 
	// for handle packges and import them
	4- type command "go mod vendor"
	// for run the program
	5- type command "go run main.go"
	
for Apis
all Apis work with authentication jwt token with secret key "User key"
you should in postman header section example
key :Jwt-Token
value : eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJPbmxpbmUgSldUIEJ1aWxkZXIiLCJpYXQiOjE2Mjc5MzQ1MjcsImV4cCI6MTY1OTQ3MDU2OCwiYXVkIjoid3d3LmV4YW1wbGUuY29tIiwic3ViIjoianJvY2tldEBleGFtcGxlLmNvbSIsIkdpdmVuTmFtZSI6IkpvaG5ueSIsIlN1cm5hbWUiOiJSb2NrZXQiLCJFbWFpbCI6Impyb2NrZXRAZXhhbXBsZS5jb20iLCJSb2xlIjpbIk1hbmFnZXIiLCJQcm9qZWN0IEFkbWluaXN0cmF0b3IiXX0.7ERAAuOuS4NoPaFnUO7DYm-3K--fpWYrO2dTVwLT5v8

1-create user api
url : localhost:8080/user/create
method : POST
body : {
    "user_name":"ahmed",
    "age":26
}

2-edit user api
url : localhost:8080/user/edit
method : POST
body : {
    "user_name":"ahmed",
    "age":30
}

3-add task for user todo list
url : localhost:8081/todo/add
method : POST
body : {
    "user_name":"ahmed",
    "task":"task3",
    "date":"2020-04-23T19:53:52.052717874Z"
}

4-get  user upcoming tasks todo list
url : localhost:8081/todo/get
method : POST
body : {
    "user_name":"ahmed"
}


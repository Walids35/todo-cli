package main

import (
	"github.com/walids35/todo-cli/todo"
)

const filename string = "data.csv"

func main(){
	var todos []todo.Todo
	todos, err := todo.LoadFromCSVFile(filename)
	if err != nil {
		todos = []todo.Todo{}
	}
	todo.ListTodos(todos)
}

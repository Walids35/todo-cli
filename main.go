package main

import "github.com/walids35/todo-cli/todo"

const filename string = "../data.csv"

func main(){
	todo.LoadFromCSVFile(filename)
}

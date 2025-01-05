/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/walids35/todo-cli/todo"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "This command will add a new todo",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		todos, _ := todo.LoadFromCSVFile("data.csv")
		name := args[0]
		todo.AddTodo(name, &todos)
		fmt.Println("Successfully added the todo item !")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

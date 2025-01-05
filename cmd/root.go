package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
)

var rootCmd = &cobra.Command{
  Use:   "todo-cli",
  Short: "Todo Cli is a simple and basic Todo List application with GO",
  Long: `A Fast and Flexible Todo List Application built with
                love by walids35`,
  Run: func(cmd *cobra.Command, args []string) {
    // Do Stuff Here
  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}

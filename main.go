package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	(&cli.App{
		Name:        "Task Tracker",
		Description: "Add, Delete, and Update your tasks",
		Usage:       "to track your tasks",
		Action: func(ctx *cli.Context) error {
			fmt.Println("Hello cliLearn")
			return nil
		},
	}).Run(os.Args)
}

// Package filter provides a function to filter commands that are used on the CLI
package filter

import (
	"fmt"
	"os"
	"todo/commands"
)

func delegateOrThrow(command *string, exec func(), eMsg string) {
	if command == nil {
		fmt.Println(eMsg)
		return
	}
	exec()
}

func FilterCommand() {
	var order *string

	if len(os.Args) < 2 {
		fmt.Println("Usage: todo <command> [argument]")
		return
	}

	if len(os.Args) > 2 {
		order = &os.Args[2]
	}

	command := os.Args[1]

	if len(command) <= 0 {
		fmt.Println("A command is required")
	}

	if command == "list" {
		commands.List()
	}

	if command == "version" {
		commands.List()
	}

	if command == "add" {
		delegateOrThrow(order, func() { commands.Add(order) }, "a description is required")
	}

	if command == "rm" {
		delegateOrThrow(order, func() { commands.Rm(order) }, "a task ID is required")
	}

	if command == "check" {
		delegateOrThrow(order, func() { commands.Check(order) }, "a task ID is required")
	}

}

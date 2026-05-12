// Package commands is the package that have all the commands for the CLI
package commands

import (
	"fmt"
	"strconv"
	"todo/store"
	"todo/tasks"
)

func Add(desc *string) {
	err := tasks.Save(*desc)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("task registered: ", *desc)

	List()
}

func Rm(id *string) {
	idToBeRemoved, err := strconv.Atoi(*id)
	if err != nil {
		fmt.Println("Error converting string to int")
	}
	removeByIDError := tasks.RemoveByID(idToBeRemoved)
	if removeByIDError != nil {
		fmt.Println("An error occurred")
	}
	fmt.Println("the task of id " + *id + " is removed")

	List()
}

func Check(id *string) {
	idToBeChecked, err := strconv.Atoi(*id)
	if err != nil {
		fmt.Println("Error converting string to int")
	}
	checkError := tasks.CheckByID(idToBeChecked)
	if checkError != nil {
		fmt.Println("An error occurred")
	}

	fmt.Println("the task of id " + *id + " is checked as done")

	List()
}

func List() {
	todoStore, err := store.GetStore()

	if err != nil {
		fmt.Println("An error occurred when accessing store")
	}

	todoTasks := todoStore.Tasks

	fmt.Println("Tasks: ")
	for _, task := range todoTasks {
		if task.Done {
			fmt.Printf("id: %d - [ x ] %s \n", task.ID, task.Text)
		} else {
			fmt.Printf("id: %d - [  ] %s \n", task.ID, task.Text)
		}
	}
}

func Version() {
	fmt.Println("TODO version 1.0.0")
}

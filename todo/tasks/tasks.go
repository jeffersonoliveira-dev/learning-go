// Package tasks is responsible for being a connection ( read and write ) tasks data into a JSON file
package tasks

import (
	"fmt"
	"todo/store"
)

func generateID(todoStore store.TodoStore) int {
	newID := 1

	if len(todoStore.Tasks) > 0 {
		newID = todoStore.Tasks[len(todoStore.Tasks)-1].ID + 1
	}

	return newID
}

func Save(text string) error {
	todo, err := store.GetStore()

	if err != nil {
		return nil
	}

	newTask := store.Task{
		ID:   generateID(todo),
		Text: text,
		Done: false,
	}

	todo.Tasks = append(todo.Tasks, newTask)
	todo.Metadata.Elements = len(todo.Tasks)

	return store.SaveStore(todo)
}

func getByID(id int) (*store.Task, error) {
	todo, _ := store.GetStore()
	tasks := todo.Tasks

	for _, task := range tasks {
		if task.ID == id {
			return &task, nil
		}
	}

	return nil, fmt.Errorf("task with ID %d not found", id)
}

func getAll() ([]store.Task, error) {
	store, err := store.GetStore()

	if err != nil {
		return nil, err
	}

	return store.Tasks, nil
}

func RemoveByID(id int) error {
	todo, err := store.GetStore()

	if err != nil {
		return err
	}

	var updatedTasks []store.Task

	for _, task := range todo.Tasks {
		if task.ID != id {
			updatedTasks = append(updatedTasks, task)
		}
	}

	todo.Tasks = updatedTasks
	todo.Metadata.Elements = len(updatedTasks)

	return store.SaveStore(todo)
}

func CheckByID(id int) error {
	todo, err := store.GetStore()

	if err != nil {
		return err
	}

	var updatedTasks []store.Task

	for _, task := range todo.Tasks {
		if task.ID == id {
			task.Done = !task.Done
		}
		updatedTasks = append(updatedTasks, task)
	}

	todo.Tasks = updatedTasks
	todo.Metadata.Elements = len(updatedTasks)

	return store.SaveStore(todo)
}

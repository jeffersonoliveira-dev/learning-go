// Package store is responsible to store tasks
package store

import (
	"encoding/json"
	"os"
)

type Metadata struct {
	Elements int `json:"elements"`
}

type Task struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

type TodoStore struct {
	Metadata Metadata `json:"metadata"`
	Tasks    []Task   `json:"tasks"`
}

func GetStore() (TodoStore, error) {
	fileData, err := os.ReadFile("tasks/tasks.json")
	var todo TodoStore

	if err == nil {
		json.Unmarshal(fileData, &todo)
	} else if !os.IsNotExist(err) {
		return todo, err
	}

	return todo, nil
}

func SaveStore(newStore TodoStore) error {
	updatedJSON, err := json.MarshalIndent(newStore, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile("tasks/tasks.json", updatedJSON, 0644)
}

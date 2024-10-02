package tasks

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const taskFile = "tasks.json"

// LoadTasks loads tasks from the tasks.json file
func LoadTasks() ([]Task, error) {
	if _, err := os.Stat(taskFile); os.IsNotExist(err) {
		return []Task{}, nil
	}
	data, err := os.ReadFile(taskFile)
	if err != nil {
		return nil, err
	}
	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

func SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(taskFile, data, 0644)
}

func AddTask(description string) {
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}
	newTask := Task{
		Description: description,
		Status:      StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	tasks = append(tasks, newTask)
	if err := SaveTasks(tasks); err != nil {
		fmt.Println("Error saving task:", err)
	} else {
		fmt.Printf("Task added successfully: %s\n", newTask.Description)
	}
}

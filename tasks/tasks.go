package tasks

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// TaskService manages task operations and ensures persistence.
type TaskService struct {
	Filename string
	tasks    []Task
}

// AddTask only modifies the in-memory task list, without saving to file.
func (ts *TaskService) AddTask(description string) {
	newTask := Task{
		Description: description,
		Status:      StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	ts.tasks = append(ts.tasks, newTask)
	fmt.Printf("Task added successfully: %s\n", newTask.Description)
}

func (ts *TaskService) SaveTasks() error {
	data, err := json.MarshalIndent(ts.tasks, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(ts.Filename, data, 0644)
}

// LoadTasks from file
func (ts *TaskService) LoadTasks() error {
	data, err := os.ReadFile(ts.Filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // If file doesn't exist, assume no tasks yet.
		}
		return err
	}
	return json.Unmarshal(data, &ts.tasks)
}

func (ts *TaskService) ListTasks() {
	if len(ts.tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}

	for i, task := range ts.tasks {
		fmt.Printf(
			"[%d] %s - %s (Created: %s, Updated: %s)\n",
			i,
			task.Description,
			task.Status,
			task.CreatedAt,
			task.UpdatedAt,
		)
	}
}

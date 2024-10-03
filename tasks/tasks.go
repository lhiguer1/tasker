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

func (ts *TaskService) MarkDone(i int) {
	if i >= len(ts.tasks) {
		fmt.Println("Index out of bounds.")
		return
	}
	ts.tasks[i].Status = StatusDone
}

func (ts *TaskService) ListDoneTasks() {
	if len(ts.tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}

	for i, task := range ts.tasks {
		if task.Status == StatusDone {
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
}

func (ts *TaskService) ListInProgressTasks() {
	if len(ts.tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}

	for i, task := range ts.tasks {
		if task.Status == StatusInProgress {
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
}

func (ts *TaskService) ListTodoTasks() {
	if len(ts.tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}

	for i, task := range ts.tasks {
		if task.Status == StatusTodo {
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
}

func (ts *TaskService) MarkInProgress(i int) {
	if i >= len(ts.tasks) {
		fmt.Println("Index out of bounds.")
		return
	}
	ts.tasks[i].Status = StatusInProgress
}

// DeleteTask only removes the task from in-memory list.
func (ts *TaskService) DeleteTask(i int) {
	if i >= len(ts.tasks) {
		fmt.Println("Index out of bounds.")
		return
	}
	ts.tasks = append(ts.tasks[:i], ts.tasks[i+1:]...)
	fmt.Printf("Task %d deleted successfully.\n", i)
}

func (ts *TaskService) UpdateTask(i int, newDescription string) {
	if i >= len(ts.tasks) {
		fmt.Println("Index out of bounds.")
		return
	}
	if i < len(ts.tasks) {
		ts.tasks[i].Description = newDescription
		ts.tasks[i].UpdatedAt = time.Now()
		fmt.Printf("Task %d updated successfully.\n", i)
		return
	}
	fmt.Printf("Task %d not found.\n", i)
}

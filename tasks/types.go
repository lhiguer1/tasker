package tasks

import (
	"time"
)

// TaskStatus is an enum-like type for the status of a task
type TaskStatus string

const (
	StatusTodo       TaskStatus = "todo"
	StatusInProgress TaskStatus = "in-progress"
	StatusDone       TaskStatus = "done"
)

type Task struct {
	Description string     // A short description of the task
	Status      TaskStatus // The status of the task (todo, in-progress, done)
	CreatedAt   time.Time  // The date and time when the task was created
	UpdatedAt   time.Time  // The date and time when the task was last updated
}

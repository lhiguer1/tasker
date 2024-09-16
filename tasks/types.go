package tasks

import "time"

type Task struct {
	description string     // A short description of the task
	status      string     // The status of the task (todo, in-progress, done)
	createdAt   time.Time  // The date and time when the task was created
	updatedAt   *time.Time // The date and time when the task was last updated
}

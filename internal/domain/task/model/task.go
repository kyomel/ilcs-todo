package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	DueDate     time.Time `json:"due_date"`
}

type TaskRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Status      string `json:"status"`
	DueDate     string `json:"due_date"`
}

type Pagination struct {
	CurrentPage int `json:"current_page"`
	TotalPages  int `json:"total_pages"`
	TotalTasks  int `json:"total_tasks"`
}

type AllTasks struct {
	Tasks      []*Task     `json:"tasks"`
	Pagination *Pagination `json:"pagination"`
}

func (r *TaskRequest) Validate() error {
	if r.Title == "" {
		return fmt.Errorf("title is required")
	}

	if r.Description == "" {
		return fmt.Errorf("description is required")
	}

	if r.DueDate == "" {
		return fmt.Errorf("due date is required")
	}

	dueDate, err := time.Parse("2006-01-02", r.DueDate)
	if err != nil {
		return fmt.Errorf("invalid due date format: must be YYYY-MM-DD")
	}

	if dueDate.Before(time.Now().Truncate(24 * time.Hour)) {
		return fmt.Errorf("due date cannot be in the past")
	}

	return nil
}

func (r *TaskRequest) ParseDueDate() (time.Time, error) {
	if r.DueDate == "" {
		return time.Time{}, nil
	}
	return time.Parse("2006-01-02", r.DueDate)
}

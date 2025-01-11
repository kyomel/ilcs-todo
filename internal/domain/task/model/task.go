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

type CreateTaskRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Status      string `json:"status"`
	DueDate     string `json:"due_date"`
}

func (r *CreateTaskRequest) Validate() error {
	if r.Title == "" {
		return fmt.Errorf("title is required")
	}

	if r.Status != "" && !IsValidTaskStatus(r.Status) {
		return fmt.Errorf("invalid status: must be either 'pending' or 'completed'")
	}

	if r.Status == "" {
		r.Status = string(TaskStatusPending)
	}

	return nil
}

func (r *CreateTaskRequest) ParseDueDate() (time.Time, error) {
	if r.DueDate == "" {
		return time.Time{}, nil
	}
	return time.Parse("2006-01-02", r.DueDate)
}

package task

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/kyomel/ilcs-todo/internal/model"
)

type taskRepo struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) Repository {
	return &taskRepo{
		db: db,
	}
}

func (r *taskRepo) PostTask(ctx context.Context, req *model.CreateTaskRequest) (*model.Task, error) {
	var result model.Task

	// Parse due date
	dueDate, err := req.ParseDueDate()
	if err != nil {
		return nil, err
	}

	// Generate UUID
	taskID := uuid.New()

	query := `
		INSERT INTO tasks (id, title, description, status, due_date)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, title, description, status, due_date, created_at, updated_at
	`

	err = r.db.QueryRowContext(ctx, query,
		taskID,
		req.Title,
		req.Description,
		req.Status,
		dueDate,
	).Scan(
		&result.ID,
		&result.Title,
		&result.Description,
		&result.Status,
		&result.DueDate,
		&result.CreatedAt,
		&result.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

package task

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/kyomel/ilcs-todo/internal/domain/task/entity"
)

type taskRepo struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) Repository {
	return &taskRepo{
		db: db,
	}
}

func (r *taskRepo) PostTask(ctx context.Context, req *entity.CreateTaskRequest) (*entity.Task, error) {
	var result entity.Task

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

func (r *taskRepo) GetAllTasks(ctx context.Context) ([]*entity.Task, error) {
	var tasks []*entity.Task

	query := `
		SELECT id, title, description, status, due_date, created_at, updated_at
		FROM tasks
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task entity.Task
		if err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.DueDate,
			&task.CreatedAt,
			&task.UpdatedAt,
		); err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

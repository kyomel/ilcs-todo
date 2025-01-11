package task

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/kyomel/ilcs-todo/internal/domain/task/model"
)

type taskRepo struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) Repository {
	return &taskRepo{
		db: db,
	}
}

func (r *taskRepo) PostTask(ctx context.Context, req *model.TaskRequest) (*model.Task, error) {
	var result model.Task

	dueDate, err := req.ParseDueDate()
	if err != nil {
		return nil, err
	}

	taskID := uuid.New()

	query := `
		INSERT INTO tasks (id, title, description, status, due_date)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, title, description, status, due_date
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
	)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *taskRepo) GetAllTasks(ctx context.Context) ([]*model.Task, error) {
	var tasks []*model.Task

	query := `
		SELECT id, title, description, status, due_date
		FROM tasks
		ORDER BY due_date ASC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task model.Task
		if err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.DueDate,
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

func (r *taskRepo) GetTaskByID(ctx context.Context, id uuid.UUID) (*model.Task, error) {
	var result model.Task

	query := `
		SELECT id, title, description, status, due_date
		FROM tasks
		WHERE id = $1
	`

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&result.ID,
		&result.Title,
		&result.Description,
		&result.Status,
		&result.DueDate,
	)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *taskRepo) UpdateTask(ctx context.Context, id uuid.UUID, req *model.TaskRequest) (*model.Task, error) {
	var result model.Task

	dueDate, err := req.ParseDueDate()
	if err != nil {
		return nil, err
	}

	query := `
		UPDATE tasks
		SET title = $2, description = $3, status = $4, due_date = $5
		WHERE id = $1
		RETURNING id, title, description, status, due_date
	`

	err = r.db.QueryRowContext(ctx, query,
		id,
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
	)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *taskRepo) DeleteTask(ctx context.Context, id uuid.UUID) error {
	query := `
		DELETE FROM tasks
		WHERE id = $1
	`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

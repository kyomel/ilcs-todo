package task

import (
	"context"
	"fmt"

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

func (r *taskRepo) GetTasksPaginated(ctx context.Context, page, limit int, status, search string) ([]*model.Task, error) {
	offset := (page - 1) * limit

	var query string
	var args []interface{}
	args = append(args, limit, offset)
	paramCount := 3

	baseQuery := `
		SELECT id, title, description, status, due_date
		FROM tasks
		WHERE 1=1
	`

	if status != "" {
		baseQuery += fmt.Sprintf(" AND status = $%d", paramCount)
		args = append(args, status)
		paramCount++
	}

	if search != "" {
		baseQuery += fmt.Sprintf(" AND (LOWER(title) LIKE LOWER($%d) OR LOWER(description) LIKE LOWER($%d))",
			paramCount, paramCount)
		args = append(args, "%"+search+"%")
	}

	query = baseQuery + `
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*model.Task
	for rows.Next() {
		task := &model.Task{}
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.DueDate,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *taskRepo) GetTotalTasksWithFilter(ctx context.Context, status, search string) (int, error) {
	var query string
	var args []interface{}
	paramCount := 1

	query = `
		SELECT COUNT(*)
		FROM tasks
		WHERE 1=1
	`

	if status != "" {
		query += fmt.Sprintf(" AND status = $%d", paramCount)
		args = append(args, status)
		paramCount++
	}

	if search != "" {
		query += fmt.Sprintf(" AND (LOWER(title) LIKE LOWER($%d) OR LOWER(description) LIKE LOWER($%d))",
			paramCount, paramCount)
		args = append(args, "%"+search+"%")
	}

	var total int
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&total)
	if err != nil {
		return 0, err
	}

	return total, nil
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

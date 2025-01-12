package task

import (
	"context"

	"github.com/google/uuid"
	"github.com/kyomel/ilcs-todo/internal/domain/task/model"
)

type Repository interface {
	PostTask(ctx context.Context, req *model.TaskRequest) (*model.Task, error)
	GetTasksPaginated(ctx context.Context, page, limit int, status string) ([]*model.Task, error)
	GetTotalTasksWithFilter(ctx context.Context, status string) (int, error)
	GetTaskByID(ctx context.Context, id uuid.UUID) (*model.Task, error)
	UpdateTask(ctx context.Context, id uuid.UUID, req *model.TaskRequest) (*model.Task, error)
	DeleteTask(ctx context.Context, id uuid.UUID) error
}

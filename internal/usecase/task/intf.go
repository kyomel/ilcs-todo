package task

import (
	"context"

	"github.com/google/uuid"
	"github.com/kyomel/ilcs-todo/internal/domain/task/model"
)

type Usecase interface {
	PostTask(ctx context.Context, req *model.TaskRequest) (*model.Task, error)
	GetAllTasks(ctx context.Context) ([]*model.Task, error)
	GetTaskByID(ctx context.Context, id uuid.UUID) (*model.Task, error)
}

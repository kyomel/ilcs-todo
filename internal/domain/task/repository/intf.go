package task

import (
	"context"

	"github.com/kyomel/ilcs-todo/internal/domain/task/model"
)

type Repository interface {
	PostTask(ctx context.Context, req *model.CreateTaskRequest) (*model.Task, error)
	GetAllTasks(ctx context.Context) ([]*model.Task, error)
}

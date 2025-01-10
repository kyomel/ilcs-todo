package task

import (
	"context"

	"github.com/kyomel/ilcs-todo/internal/domain/task/entity"
)

type Repository interface {
	PostTask(ctx context.Context, req *entity.CreateTaskRequest) (*entity.Task, error)
	GetAllTasks(ctx context.Context) ([]*entity.Task, error)
}

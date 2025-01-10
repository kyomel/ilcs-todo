package task

import (
	"context"

	"github.com/kyomel/ilcs-todo/internal/domain/task/entity"
)

type Usecase interface {
	PostTask(ctx context.Context, req *entity.CreateTaskRequest) (*entity.Task, error)
}

package task

import (
	"context"

	"github.com/kyomel/ilcs-todo/internal/model"
)

type Repository interface {
	PostTask(ctx context.Context, req *model.CreateTaskRequest) (*model.Task, error)
}

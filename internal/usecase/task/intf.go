package task

import (
	"context"

	"github.com/kyomel/ilcs-todo/internal/model"
)

type Usecase interface {
	PostTask(ctx context.Context, req *model.CreateTaskRequest) (*model.Task, error)
}

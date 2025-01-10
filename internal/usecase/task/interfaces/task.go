package task

import (
	"context"
	"time"

	"github.com/kyomel/ilcs-todo/internal/domain/task/entity"
	task "github.com/kyomel/ilcs-todo/internal/domain/task/repository"
)

type useCase struct {
	taskRepo   task.Repository
	ctxTimeout time.Duration
}

func NewUsecase(taskRepo task.Repository, ctxTimeout time.Duration) Usecase {
	return &useCase{
		taskRepo:   taskRepo,
		ctxTimeout: ctxTimeout,
	}
}

func (uc *useCase) PostTask(ctx context.Context, req *entity.CreateTaskRequest) (*entity.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.ctxTimeout)
	defer cancel()

	task, err := uc.taskRepo.PostTask(ctx, req)
	if err != nil {
		return nil, err
	}

	return task, err
}
package task

import (
	"context"
	"time"

	"github.com/kyomel/ilcs-todo/internal/domain/task/model"
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

func (uc *useCase) PostTask(ctx context.Context, req *model.CreateTaskRequest) (*model.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.ctxTimeout)
	defer cancel()

	task, err := uc.taskRepo.PostTask(ctx, req)
	if err != nil {
		return nil, err
	}

	return task, err
}

func (uc *useCase) GetAllTasks(ctx context.Context) ([]*model.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.ctxTimeout)
	defer cancel()

	tasks, err := uc.taskRepo.GetAllTasks(ctx)
	if err != nil {
		return nil, err
	}

	return tasks, err
}

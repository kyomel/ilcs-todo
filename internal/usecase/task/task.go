package task

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/kyomel/ilcs-todo/internal/domain/task/model"
	task "github.com/kyomel/ilcs-todo/internal/domain/task/repository"
)

type useCase struct {
	taskRepo   task.Repository
	ctxTimeout time.Duration
}

func NewTaskUseCase(repo task.Repository, timeout time.Duration) Usecase {
	return &useCase{
		taskRepo:   repo,
		ctxTimeout: timeout,
	}
}

func (uc *useCase) PostTask(ctx context.Context, req *model.TaskRequest) (*model.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.ctxTimeout)
	defer cancel()

	if err := req.Validate(); err != nil {
		return nil, err
	}

	task, err := uc.taskRepo.PostTask(ctx, req)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (uc *useCase) GetAllTasks(ctx context.Context, page, limit int, status, search string) (*model.AllTasks, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.ctxTimeout)
	defer cancel()

	var response model.AllTasks

	tasksChan := make(chan []*model.Task, 1)
	totalChan := make(chan int, 1)
	errorChan := make(chan error, 2)

	go func() {
		tasks, err := uc.taskRepo.GetTasksPaginated(ctx, page, limit, status, search)
		if err != nil {
			errorChan <- err
			return
		}
		tasksChan <- tasks
	}()

	go func() {
		total, err := uc.taskRepo.GetTotalTasksWithFilter(ctx, status, search)
		if err != nil {
			errorChan <- err
			return
		}
		totalChan <- total
	}()

	var tasks []*model.Task
	var total int

	for i := 0; i < 2; i++ {
		select {
		case err := <-errorChan:
			return nil, err
		case tasks = <-tasksChan:
		case total = <-totalChan:
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}

	if tasks == nil {
		tasks = []*model.Task{}
	}

	totalPages := (total + limit - 1) / limit
	response = model.AllTasks{
		Tasks: tasks,
		Pagination: &model.Pagination{
			CurrentPage: page,
			TotalPages:  totalPages,
			TotalTasks:  total,
		},
	}

	return &response, nil
}

func (uc *useCase) GetTaskByID(ctx context.Context, id uuid.UUID) (*model.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.ctxTimeout)
	defer cancel()

	task, err := uc.taskRepo.GetTaskByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return task, err
}

func (uc *useCase) UpdateTask(ctx context.Context, id uuid.UUID, req *model.TaskRequest) (*model.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.ctxTimeout)
	defer cancel()

	if err := req.Validate(); err != nil {
		return nil, err
	}

	task, err := uc.taskRepo.UpdateTask(ctx, id, req)
	if err != nil {
		return nil, err
	}

	return task, err
}

func (uc *useCase) DeleteTask(ctx context.Context, id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, uc.ctxTimeout)
	defer cancel()

	err := uc.taskRepo.DeleteTask(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

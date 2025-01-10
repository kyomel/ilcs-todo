package handlers

import (
	task "github.com/kyomel/ilcs-todo/internal/usecase/task/interfaces"
)

type Handlers struct {
	taskUsecase task.Usecase
}

func NewHandler(taskUsecase task.Usecase) *Handlers {
	return &Handlers{
		taskUsecase: taskUsecase,
	}
}

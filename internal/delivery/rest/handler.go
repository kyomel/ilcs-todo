package rest

import "github.com/kyomel/ilcs-todo/internal/usecase/task"

type handler struct {
	taskUsecase task.Usecase
}

func NewHandler(taskUsecase task.Usecase) *handler {
	return &handler{
		taskUsecase: taskUsecase,
	}
}

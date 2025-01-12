package handlers

import (
	task "github.com/kyomel/ilcs-todo/internal/usecase/task"
	"github.com/kyomel/ilcs-todo/internal/usecase/user"
)

type Handlers struct {
	taskUsecase task.Usecase
	userUsecase user.Usecase
}

func NewHandler(taskUsecase task.Usecase, userUsecase user.Usecase) *Handlers {
	return &Handlers{
		taskUsecase: taskUsecase,
		userUsecase: userUsecase,
	}
}

package handlers

import (
	"github.com/kyomel/ilcs-todo/internal/domain/task/entity"
	"github.com/kyomel/ilcs-todo/pkg/logger"
	"github.com/labstack/echo/v4"
)

func (h *Handlers) PostTask(c echo.Context) error {
	log := logger.GetLogger()
	log.Info("Received a request to post a task")

	var req entity.CreateTaskRequest
	ctx := c.Request().Context()
	if err := c.Bind(&req); err != nil {
		log.Errorf("Failed to bind request: %v", err)
		return c.JSON(400, map[string]interface{}{
			"error": "Invalid request format",
		})
	}

	if err := req.Validate(); err != nil {
		return c.JSON(400, map[string]interface{}{
			"error": err.Error(),
		})
	}

	task, err := h.taskUsecase.PostTask(ctx, &req)
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"error": err.Error(),
		})
	}

	log.Info("Task created successfully")
	return c.JSON(201, map[string]interface{}{
		"message": "Task created successfully",
		"task":    task,
	})
}

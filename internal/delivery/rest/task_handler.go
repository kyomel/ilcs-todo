package rest

import (
	"github.com/kyomel/ilcs-todo/internal/model"
	"github.com/labstack/echo/v4"
)

func (h *handler) PostTask(c echo.Context) error {
	var req model.CreateTaskRequest
	ctx := c.Request().Context()
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, map[string]interface{}{
			"error": "Invalid request format",
		})
	}

	// Validate request including status
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

	return c.JSON(201, map[string]interface{}{
		"message": "Task created successfully",
		"task":    task,
	})
}

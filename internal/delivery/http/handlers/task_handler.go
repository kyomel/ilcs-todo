package handlers

import (
	"github.com/google/uuid"
	"github.com/kyomel/ilcs-todo/internal/domain/task/model"
	"github.com/kyomel/ilcs-todo/pkg/logger"
	"github.com/labstack/echo/v4"
)

func (h *Handlers) PostTask(c echo.Context) error {
	log := logger.GetLogger()
	log.Info("Received a request to post a task")

	var req model.TaskRequest
	ctx := c.Request().Context()
	if err := c.Bind(&req); err != nil {
		log.Errorf("Failed to bind request: %v", err)
		return c.JSON(400, map[string]interface{}{
			"error": "Invalid request format",
		})
	}

	if err := req.Validate(); err != nil {
		log.Errorf("Failed to bind validate request: %v", err)
		return c.JSON(400, map[string]interface{}{
			"error": err.Error(),
		})
	}

	task, err := h.taskUsecase.PostTask(ctx, &req)
	if err != nil {
		log.Errorf("Failed to post task: %v", err)
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

func (h *Handlers) GetAllTasks(c echo.Context) error {
	log := logger.GetLogger()

	ctx := c.Request().Context()
	tasks, err := h.taskUsecase.GetAllTasks(ctx)
	if err != nil {
		log.Errorf("Failed to get all task: %v", err)
		return c.JSON(400, map[string]interface{}{
			"error": err.Error(),
		})
	}

	log.Info("Get all task successfully")
	return c.JSON(200, map[string]interface{}{
		"tasks": tasks,
	})
}

func (h *Handlers) GetTaskByID(c echo.Context) error {
	log := logger.GetLogger()
	log.Info("Received a request to get task by id")

	id := c.Param("id")
	ctx := c.Request().Context()
	idUUID, err := uuid.Parse(id)
	if err != nil {
		log.Errorf("Failed to parse id: %v", err)
		return c.JSON(400, map[string]interface{}{
			"error": "Invalid id format",
		})
	}

	task, err := h.taskUsecase.GetTaskByID(ctx, idUUID)
	if err != nil {
		log.Errorf("Failed to get task by id: %v", err)
		return c.JSON(400, map[string]interface{}{
			"error": err.Error(),
		})
	}

	log.Info("Get task by id successfully")
	return c.JSON(200, map[string]interface{}{
		"task": task,
	})
}

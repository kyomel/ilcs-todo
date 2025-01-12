package handlers

import (
	"strconv"

	"github.com/google/uuid"
	"github.com/kyomel/ilcs-todo/internal/domain/task/model"
	"github.com/kyomel/ilcs-todo/pkg/logger"
	"github.com/labstack/echo/v4"
)

const (
	parseIDError             = "Failed to parse id: %v"
	invalidIDFormat          = "Invalid id format"
	bindRequestError         = "Failed to bind request: %v"
	invalidRequestFormat     = "Invalid request format"
	bindValidateRequestError = "Failed to bind validate request: %v"
	postTaskError            = "Failed to post task: %v"
	getAllTaskError          = "Failed to get all task: %v"
	getTaskByIDError         = "Failed to get task by id: %v"
	updateTaskError          = "Failed to update task: %v"
	deleteTaskError          = "Failed to delete task: %v"
)

func (h *Handlers) PostTask(c echo.Context) error {
	log := logger.GetLogger()

	var req model.TaskRequest
	ctx := c.Request().Context()
	if err := c.Bind(&req); err != nil {
		log.Errorf(bindRequestError, err)
		return c.JSON(400, map[string]interface{}{
			"error": invalidRequestFormat,
		})
	}

	if err := req.Validate(); err != nil {
		log.Errorf(bindValidateRequestError, err)
		return c.JSON(400, map[string]interface{}{
			"error": err.Error(),
		})
	}

	task, err := h.taskUsecase.PostTask(ctx, &req)
	if err != nil {
		log.Errorf(postTaskError, err)
		return c.JSON(500, map[string]interface{}{
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

	page := 1
	limit := 10

	if pageStr := c.QueryParam("page"); pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	if limitStr := c.QueryParam("limit"); limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	response, err := h.taskUsecase.GetAllTasks(ctx, page, limit)
	if err != nil {
		log.Errorf(getAllTaskError, err)
		return c.JSON(500, map[string]interface{}{
			"error": err.Error(),
		})
	}

	log.Info("Get all task successfully")
	return c.JSON(200, response)
}

func (h *Handlers) GetTaskByID(c echo.Context) error {
	log := logger.GetLogger()
	log.Info("Received a request to get task by id")

	id := c.Param("id")
	ctx := c.Request().Context()
	idUUID, err := uuid.Parse(id)
	if err != nil {
		log.Errorf(parseIDError, err)
		return c.JSON(400, map[string]interface{}{
			"error": invalidIDFormat,
		})
	}

	task, err := h.taskUsecase.GetTaskByID(ctx, idUUID)
	if err != nil {
		log.Errorf(getTaskByIDError, err)
		return c.JSON(500, map[string]interface{}{
			"error": err.Error(),
		})
	}

	log.Info("Get task by id successfully")
	return c.JSON(200, map[string]interface{}{
		"task": task,
	})
}

func (h *Handlers) UpdateTask(c echo.Context) error {
	log := logger.GetLogger()

	var req model.TaskRequest
	ctx := c.Request().Context()
	if err := c.Bind(&req); err != nil {
		log.Errorf(bindRequestError, err)
		return c.JSON(400, map[string]interface{}{
			"error": invalidRequestFormat,
		})
	}

	if err := req.Validate(); err != nil {
		log.Errorf(bindValidateRequestError, err)
		return c.JSON(400, map[string]interface{}{
			"error": err.Error(),
		})
	}

	id := c.Param("id")
	idUUID, err := uuid.Parse(id)
	if err != nil {
		log.Errorf(parseIDError, err)
		return c.JSON(400, map[string]interface{}{
			"error": invalidIDFormat,
		})
	}

	task, err := h.taskUsecase.UpdateTask(ctx, idUUID, &req)
	if err != nil {
		log.Errorf(updateTaskError, err)
		return c.JSON(500, map[string]interface{}{
			"error": err.Error(),
		})
	}

	log.Info("Task updated successfully")
	return c.JSON(200, map[string]interface{}{
		"message": "Task updated successfully",
		"task":    task,
	})
}

func (h *Handlers) DeleteTask(c echo.Context) error {
	log := logger.GetLogger()

	id := c.Param("id")
	ctx := c.Request().Context()
	idUUID, err := uuid.Parse(id)
	if err != nil {
		log.Errorf(parseIDError, err)
		return c.JSON(400, map[string]interface{}{
			"error": invalidIDFormat,
		})
	}

	err = h.taskUsecase.DeleteTask(ctx, idUUID)
	if err != nil {
		log.Errorf(deleteTaskError, err)
		return c.JSON(500, map[string]interface{}{
			"error": err.Error(),
		})
	}

	log.Info("Task deleted successfully")
	return c.JSON(200, map[string]interface{}{
		"message": "Task deleted successfully",
	})
}

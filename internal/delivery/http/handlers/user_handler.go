package handlers

import (
	"github.com/kyomel/ilcs-todo/internal/domain/user/model"
	"github.com/kyomel/ilcs-todo/pkg/logger"
	"github.com/labstack/echo/v4"
)

const (
	registerUserError = "Failed to register user: %v"
	loginUserError    = "Failed to login user: %v"
)

func (h *Handlers) RegisterUser(c echo.Context) error {
	log := logger.GetLogger()

	var req model.UserRequest
	ctx := c.Request().Context()
	if err := c.Bind(&req); err != nil {
		log.Errorf(bindRequestError, err)
		return c.JSON(400, map[string]interface{}{
			"error": invalidRequestFormat,
		})
	}

	user, err := h.userUsecase.RegisterUser(ctx, &req)
	if err != nil {
		log.Errorf(registerUserError, err)
		return c.JSON(500, map[string]interface{}{
			"error": err.Error(),
		})
	}

	log.Info("User registered successfully")
	return c.JSON(201, map[string]interface{}{
		"user": user,
	})
}

func (h *Handlers) Login(c echo.Context) error {
	log := logger.GetLogger()

	var req model.Login
	ctx := c.Request().Context()
	if err := c.Bind(&req); err != nil {
		log.Errorf(bindRequestError, err)
		return c.JSON(400, map[string]interface{}{
			"error": invalidRequestFormat,
		})
	}

	token, err := h.userUsecase.Login(ctx, &req)
	if err != nil {
		log.Errorf(loginUserError, err)
		return c.JSON(500, map[string]interface{}{
			"error": err.Error(),
		})
	}

	log.Info("User logged in successfully")
	return c.JSON(200, map[string]interface{}{
		"token": token,
	})
}

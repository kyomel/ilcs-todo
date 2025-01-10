package rest

import (
	"github.com/labstack/echo/v4"
)

func LoadRoutes(e *echo.Echo, handler *handler) {
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "I'm Alive")
	})

	taskGroup := e.Group("/tasks")
	taskGroup.POST("", handler.PostTask)
}

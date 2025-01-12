package router

import (
	"github.com/kyomel/ilcs-todo/internal/delivery/http/handlers"
	"github.com/labstack/echo/v4"
)

func LoadRoutes(e *echo.Echo, handler *handlers.Handlers) {
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "I'm Alive")
	})

	taskGroup := e.Group("/tasks")
	taskGroup.POST("", handler.PostTask)
	taskGroup.GET("", handler.GetAllTasks)
	taskGroup.GET("/:id", handler.GetTaskByID)
	taskGroup.PUT("/:id", handler.UpdateTask)
	taskGroup.DELETE("/:id", handler.DeleteTask)

	userGroup := e.Group("/users")
	userGroup.POST("", handler.RegisterUser)
	userGroup.POST("/login", handler.Login)
}

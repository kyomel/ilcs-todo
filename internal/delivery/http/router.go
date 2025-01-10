package delivery

import "github.com/labstack/echo/v4"

func LoadRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "I'm Alive")
	})
}

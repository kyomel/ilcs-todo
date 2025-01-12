package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/kyomel/ilcs-todo/internal/utils/jwt"
	"github.com/kyomel/ilcs-todo/pkg/logger"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			log := logger.GetLogger()

			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"error": "Authorization header is required",
				})
			}

			splitToken := strings.Split(authHeader, "Bearer ")
			if len(splitToken) != 2 {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"error": "Invalid authorization format",
				})
			}

			tokenString := splitToken[1]
			claims, err := jwt.ValidateToken(tokenString, os.Getenv("JWT_SECRET"))
			if err != nil {
				log.Errorf("Invalid token: %v", err)
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"error": "Invalid token",
				})
			}

			c.Set("userID", claims.UserID)
			c.Set("email", claims.Email)
			c.Set("fullName", claims.FullName)

			return next(c)
		}
	}
}

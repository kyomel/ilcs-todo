package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/kyomel/ilcs-todo/internal/config"
	delivery "github.com/kyomel/ilcs-todo/internal/delivery/http"
	"github.com/kyomel/ilcs-todo/internal/logger"
	"github.com/labstack/echo/v4"
)

func main() {
	configApp := config.LoadConfig()

	logger.Init()
	e := echo.New()

	delivery.LoadRoutes(e)

	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(configApp.ContextTimeout)*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

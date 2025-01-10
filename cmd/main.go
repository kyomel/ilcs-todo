package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/kyomel/ilcs-todo/internal/config"
	rest "github.com/kyomel/ilcs-todo/internal/delivery/rest"
	"github.com/kyomel/ilcs-todo/internal/infrastructure/datastore"
	"github.com/kyomel/ilcs-todo/internal/logger"
	tRepo "github.com/kyomel/ilcs-todo/internal/repository/task"
	tUC "github.com/kyomel/ilcs-todo/internal/usecase/task"
	"github.com/labstack/echo/v4"
)

func main() {
	configApp := config.LoadConfig()
	logger.Init()
	e := echo.New()

	dbInstance, err := datastore.NewDatabase(configApp.DatabaseURL)
	if err != nil {
		panic(err)
	}

	taskRepo := tRepo.NewTaskRepository(dbInstance)

	ctxTimeout := time.Duration(configApp.ContextTimeout) * time.Second
	taskUC := tUC.NewUsecase(taskRepo, ctxTimeout)

	h := rest.NewHandler(taskUC)

	rest.LoadRoutes(e, h)

	go func() {
		if err := e.Start(":14045"); err != nil && err != http.ErrServerClosed {
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

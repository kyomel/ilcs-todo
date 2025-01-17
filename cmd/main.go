package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/kyomel/ilcs-todo/internal/delivery/http/handlers"
	"github.com/kyomel/ilcs-todo/internal/delivery/http/router"
	tRepo "github.com/kyomel/ilcs-todo/internal/domain/task/repository"
	uRepo "github.com/kyomel/ilcs-todo/internal/domain/user/repository"
	database "github.com/kyomel/ilcs-todo/internal/infrastructure/database"
	tUC "github.com/kyomel/ilcs-todo/internal/usecase/task"
	uUC "github.com/kyomel/ilcs-todo/internal/usecase/user"
	"github.com/kyomel/ilcs-todo/pkg/config"
	"github.com/labstack/echo/v4"
)

func main() {
	configApp := config.LoadConfig()
	e := echo.New()

	dbInstance, err := database.NewDatabase(configApp.DatabaseURL)
	if err != nil {
		panic(err)
	}

	taskRepo := tRepo.NewTaskRepository(dbInstance)
	userRepo := uRepo.NewUserRepository(dbInstance)

	ctxTimeout := time.Duration(configApp.ContextTimeout) * time.Second
	taskUC := tUC.NewTaskUseCase(taskRepo, ctxTimeout)
	userUC := uUC.NewUserUsecase(userRepo, ctxTimeout)

	h := handlers.NewHandler(taskUC, userUC)

	router.LoadRoutes(e, h)

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

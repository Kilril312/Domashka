package main

import (
	"domashka/internal/db"
	"domashka/internal/handlers"
	"domashka/internal/tasksService"
	"domashka/internal/web/tasks"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	e := echo.New()

	taskRepo := tasksService.NewTaskRepository(database)
	taskService := tasksService.NewTaskService(taskRepo)
	taskHandlers := handlers.NewTaskHandler(taskService)

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	strictHandler := tasks.NewStrictHandler(taskHandlers, nil) // тут будет ошибка
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start("localhost:8080"); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}

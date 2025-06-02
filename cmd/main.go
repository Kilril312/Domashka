package main

import (
	"domashka/internal/db"
	"domashka/internal/handlers"
	"domashka/internal/tasksService"
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
	e.GET("/tasks", taskHandlers.GetTaskHandler)
	e.POST("/tasks", taskHandlers.PostTaskHandler)
	e.PATCH("/tasks/:id", taskHandlers.PatchTaskHandler)
	e.DELETE("/tasks/:id", taskHandlers.DeleteTaskHandler)
	e.Start("localhost:8080")

}

package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
)

var tasks = make(map[int]requestBodyTask)
var id = 1

type requestBodyTask struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

type ResponseBodyTask struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func getTaskHandler(c echo.Context) error {
	var taskslice []requestBodyTask
	for _, tsk := range tasks {
		taskslice = append(taskslice, tsk)
	}
	return c.JSON(http.StatusOK, taskslice)
}

func postTaskHandler(c echo.Context) error {
	var task requestBodyTask
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseBodyTask{
			Status:  "error",
			Message: "Error binding task",
		})
	}
	task.ID = id
	id++
	tasks[task.ID] = task
	return c.JSON(http.StatusOK, task)
}

func patchTaskHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseBodyTask{
			Status:  "error",
			Message: "Error converting id to int",
		})
	}
	var updatedTask requestBodyTask
	if err := c.Bind(&updatedTask); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseBodyTask{
			Status:  "error",
			Message: "Error binding task",
		})
	}

	if _, exists := tasks[id]; !exists {
		return c.JSON(http.StatusBadRequest, ResponseBodyTask{
			Status:  "error",
			Message: "Task was not found",
		})
	}

	updatedTask.ID = id
	tasks[id] = updatedTask

	return c.JSON(http.StatusOK, updatedTask)
}

func deleteTaskHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseBodyTask{
			Status:  "error",
			Message: "Error converting id to int",
		})
	}

	if _, exists := tasks[id]; !exists {
		return c.JSON(http.StatusBadRequest, ResponseBodyTask{
			Status:  "error",
			Message: "Task was not found",
		})
	}
	delete(tasks, id)
	return c.NoContent(http.StatusNoContent)
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.GET("/task", getTaskHandler)
	e.POST("/task", postTaskHandler)
	e.PATCH("/task/:id", patchTaskHandler)
	e.DELETE("/task/:id", deleteTaskHandler)
	e.Start("localhost:8080")

}

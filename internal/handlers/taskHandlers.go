package handlers

import (
	"domashka/internal/tasksService"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TaskHandler struct {
	service tasksService.TaskService
}

func NewTaskHandler(s tasksService.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) GetTaskHandler(c echo.Context) error {
	tasks, err := h.service.GetAllTasks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not get tasks from database"})
	}
	return c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) PostTaskHandler(c echo.Context) error {
	var requestBody tasksService.RequestBodyTask
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
	}
	task, err := h.service.CreateTask(requestBody.Task)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not create task from database"})
	}

	return c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) PatchTaskHandler(c echo.Context) error {
	idParam := c.Param("id")

	var req tasksService.RequestBodyTask
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
	}

	updatedTask, err := h.service.UpdateTask(idParam, req.Task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not update task from database"})
	}

	return c.JSON(http.StatusOK, updatedTask)

}

func (h *TaskHandler) DeleteTaskHandler(c echo.Context) error {
	idParam := c.Param("id")

	if err := h.service.DeleteTaskByID(idParam); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Could not delete task",
		})
	}

	return c.NoContent(http.StatusNoContent)
}

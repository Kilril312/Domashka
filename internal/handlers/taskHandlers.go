package handlers

import (
	"domashka/internal/tasksService"
	"domashka/internal/web/tasks"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/context"
	"net/http"
)

type TaskHandler struct {
	service tasksService.TaskService
}

func (h *TaskHandler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	// Получение всех задач из сервиса
	allTasks, err := h.service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	// Создаем переменную респон типа 200джейсонРеспонс
	// Которую мы потом передадим в качестве ответа
	response := tasks.GetTasks200JSONResponse{}

	// Заполняем слайс response всеми задачами из БД
	for _, tsk := range allTasks {
		task := tasks.RequestBodyTask{
			Id:   &tsk.ID,
			Task: &tsk.Task,
		}
		response = append(response, task)
	}

	// САМОЕ ПРЕКРАСНОЕ. Возвращаем просто респонс и nil!
	return response, nil
}

func (h *TaskHandler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	taskRequest := request.Body
	// Обращаемся к сервису и создаем задачу
	taskToCreate := tasksService.RequestBodyTask{
		Task: *taskRequest.Task,
	}
	createdTask, err := h.service.CreateTask(taskToCreate)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := tasks.PostTasks201JSONResponse{
		Id:   &createdTask.ID,
		Task: &createdTask.Task,
	}
	// Просто возвращаем респонс!
	return response, nil
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

package handlers

import (
	"domashka/internal/tasksService"
	"domashka/internal/web/tasks"
	"golang.org/x/net/context"
	"strconv"
)

type TaskHandler struct {
	service tasksService.TaskService
}

func NewTaskHandler(s tasksService.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
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
	taskRequest := *request.Body.Task

	createdTask, err := h.service.CreateTask(taskRequest)

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

func (h *TaskHandler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	taskID := strconv.Itoa(request.Id)
	err := h.service.DeleteTaskByID(taskID)
	if err != nil {
		return nil, err
	}
	return tasks.DeleteTasksId204Response{}, nil
}

func (h *TaskHandler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	taskID := strconv.Itoa(request.Id)
	taskRequest := *request.Body.Task

	updatedTask, err := h.service.UpdateTask(taskID, taskRequest)
	if err != nil {
		return nil, err
	}
	response := tasks.PatchTasksId200JSONResponse{
		Id:   &updatedTask.ID,
		Task: &updatedTask.Task,
	}
	return response, nil

}

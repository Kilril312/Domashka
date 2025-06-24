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

func (h *TaskHandler) GetTasksUserId(ctx context.Context, request tasks.GetTasksUserIdRequestObject) (tasks.GetTasksUserIdResponseObject, error) {
	userTasks, err := h.service.GetTasksByUserId(request.UserId)
	if err != nil {
		return nil, err
	}
	response := tasks.GetTasksUserId200JSONResponse{}
	for _, tsk := range userTasks {
		task := tasks.RequestBodyTask{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			UserId: &tsk.User_id,
		}
		response = append(response, task)
	}
	return response, nil
}

func (h *TaskHandler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.RequestBodyTask{
			Id:   &tsk.ID,
			Task: &tsk.Task,
		}
		response = append(response, task)
	}

	return response, nil
}

func (h *TaskHandler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {

	taskRequest := *request.Body.Task
	taskUserId := *request.Body.UserId

	createdTask, err := h.service.CreateTask(taskRequest, taskUserId)

	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		UserId: &createdTask.User_id,
	}

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

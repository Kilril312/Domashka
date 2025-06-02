package tasksService

import (
	"errors"
	"fmt"
	"strconv"
)

type TaskService interface {
	CreateTask(task string) (RequestBodyTask, error)
	GetAllTasks() ([]RequestBodyTask, error)
	GetTaskByID(id string) (RequestBodyTask, error)
	UpdateTask(id, task string) (RequestBodyTask, error)
	DeleteTaskByID(id string) error
}

type taskService struct {
	repo TaskRepository
}

func NewTaskService(r TaskRepository) TaskService {
	return &taskService{repo: r}
}

func (s taskService) CreateTask(task string) (RequestBodyTask, error) {
	createTask := RequestBodyTask{Task: task}
	err := s.repo.CreateTask(createTask)
	if err != nil {
		return RequestBodyTask{}, err
	}
	return createTask, nil
}

func (s taskService) GetAllTasks() ([]RequestBodyTask, error) {
	return s.repo.GetAllTask()
}

func (s taskService) GetTaskByID(id string) (RequestBodyTask, error) {
	return s.repo.GetTaskByID(id)
}

func (s taskService) UpdateTask(id, task string) (RequestBodyTask, error) {
	id, err := strconv.Atoi(id)

	if err != nil {
		return RequestBodyTask{}, errors.New("invalid task ID format")
	}

	updatedTask, , err := s.repo.UpdateTask(id, task)
	if err != nil {
		return RequestBodyTask{}, fmt.Errorf("failed to update task: %v", err)
	}
	return updatedTask, nil
}

func (s taskService) DeleteTaskByID(id string) error {
	return s.repo.DeleteTaskByID(id)
}

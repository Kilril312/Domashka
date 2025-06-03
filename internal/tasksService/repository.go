package tasksService

import (
	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(task RequestBodyTask) (RequestBodyTask, error)
	GetAllTask() ([]RequestBodyTask, error)
	GetTaskByID(id string) (RequestBodyTask, error)
	UpdateTask(task RequestBodyTask) error
	DeleteTaskByID(id string) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task RequestBodyTask) (RequestBodyTask, error) {
	err := r.db.Create(&task).Error
	return task, err
}

func (r *taskRepository) GetAllTask() ([]RequestBodyTask, error) {
	var tasks []RequestBodyTask
	err := r.db.Find(&tasks).Error

	return tasks, err
}

func (r *taskRepository) GetTaskByID(id string) (RequestBodyTask, error) {
	var task RequestBodyTask
	var err = r.db.First(&task, "id = ?", id).Error
	return task, err
}

func (r *taskRepository) UpdateTask(task RequestBodyTask) error {
	return r.db.Save(&task).Error
}

func (r *taskRepository) DeleteTaskByID(id string) error {
	err := r.db.Delete(&RequestBodyTask{}, id).Error
	return err
}

package tasksService

type TaskService interface {
	CreateTask(task string, userid int) (RequestBodyTask, error)
	GetAllTasks() ([]RequestBodyTask, error)
	GetTaskByID(id string) (RequestBodyTask, error)
	GetTasksByUserId(userId int) ([]RequestBodyTask, error)
	UpdateTask(id string, task string) (RequestBodyTask, error)
	DeleteTaskByID(id string) error
}

type taskService struct {
	repo TaskRepository
}

func NewTaskService(r TaskRepository) TaskService {
	return &taskService{repo: r}
}

func (s taskService) CreateTask(task string, userid int) (RequestBodyTask, error) {
	createTask := RequestBodyTask{
		Task:    task,
		User_id: userid,
	}
	createdTask, err := s.repo.CreateTask(createTask)
	if err != nil {
		return RequestBodyTask{}, err
	}
	return createdTask, nil
}

func (s taskService) GetAllTasks() ([]RequestBodyTask, error) {
	return s.repo.GetAllTask()
}

func (s taskService) GetTaskByID(id string) (RequestBodyTask, error) {
	return s.repo.GetTaskByID(id)
}

func (s taskService) GetTasksByUserId(userId int) ([]RequestBodyTask, error) {
	return s.repo.GetTasksByUserId(userId)
}

func (s taskService) UpdateTask(id string, task string) (RequestBodyTask, error) {
	idTask, err := s.repo.GetTaskByID(id)
	if err != nil {
		return RequestBodyTask{}, err
	}
	updTask := RequestBodyTask{
		Task: task,
		ID:   idTask.ID,
	}
	if err := s.repo.UpdateTask(updTask); err != nil {
		return RequestBodyTask{}, err
	}

	return updTask, err

}

func (s taskService) DeleteTaskByID(id string) error {
	return s.repo.DeleteTaskByID(id)
}

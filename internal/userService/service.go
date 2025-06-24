package userService

type UserService interface {
	CreateUser(email string, password string) (Users, error)
	GetAllUsers() ([]Users, error)
	GetTasksForUser(userID int) ([]RequestBodyTask, error)
	UpdateUser(id string, email string, password string) (Users, error)
	DeleteUser(id string) error
}

type userService struct {
	repo UserRepository
}

func NewUserService(r UserRepository) UserService {
	return &userService{repo: r}
}

func (s userService) CreateUser(email string, password string) (Users, error) {
	createuser := Users{Email: email, Password: password}
	createduser, err := s.repo.CreateUser(createuser)
	if err != nil {
		return Users{}, err
	}
	return createduser, nil
}

func (s userService) GetAllUsers() ([]Users, error) {
	return s.repo.GetAllUsers()
}

func (s userService) GetTasksForUser(userID int) ([]RequestBodyTask, error) {
	tasks, err := s.repo.GetTasksForUser(userID)
	if err != nil {
		return []RequestBodyTask{}, err
	}
	return tasks, nil
}

func (s userService) UpdateUser(id string, email string, password string) (Users, error) {
	idUser, err := s.repo.GetUserById(id)
	if err != nil {
		return Users{}, err
	}

	updateuser := Users{
		Email:    email,
		Password: password,
		ID:       idUser.ID,
	}

	if err = s.repo.UpdateUser(updateuser); err != nil {
		return Users{}, err
	}
	return updateuser, nil
}

func (s userService) DeleteUser(id string) error {
	return s.repo.DeleteUser(id)
}

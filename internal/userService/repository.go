package userService

import (
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user Users) (Users, error)
	GetAllUsers() ([]Users, error)
	GetUserById(id string) (Users, error)
	GetTasksForUser(userID int) ([]RequestBodyTask, error)
	UpdateUser(user Users) error
	DeleteUser(id string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (s *userRepository) CreateUser(user Users) (Users, error) {
	err := s.db.Create(&user).Error
	return user, err
}

func (s *userRepository) GetAllUsers() ([]Users, error) {
	var us []Users
	err := s.db.Find(&us).Error
	return us, err
}

func (s *userRepository) GetUserById(id string) (Users, error) {
	var user Users
	var err = s.db.First(&user, "id = ?", id).Error
	return user, err
}

func (s *userRepository) GetTasksForUser(userID int) ([]RequestBodyTask, error) {
	var tasks []RequestBodyTask
	var err = s.db.First(&tasks, "user_id = ?", userID).Error
	return tasks, err

}

func (s *userRepository) UpdateUser(user Users) error {
	return s.db.Save(&user).Error
}

func (s *userRepository) DeleteUser(id string) error {
	err := s.db.Delete(&Users{}, id).Error
	return err
}

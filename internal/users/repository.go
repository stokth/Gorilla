package users

import (
	"Gorilla/internal/tasks"

	"gorm.io/gorm"
)

type UsersRepository interface {
	GetUsers() ([]User, error)
	GetUser(id int64) (*User, error)
	CreateUser(user *User) error
	UpdateUser(user *User) error
	DeleteUser(id int64) error
	GetTasksForUser(id int64) ([]tasks.Task, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) UsersRepository {
	return &userRepository{db: db}
}

func (u *userRepository) GetTasksForUser(id int64) ([]tasks.Task, error) {
	var tasks []tasks.Task
	err := u.db.Where("user_id = ?", id).Find(&tasks).Error
	return tasks, err

}

// CreateUser implements UsersRepository.
func (u *userRepository) CreateUser(user *User) error {
	return u.db.Create(&user).Error
}

// DeleteUser implements UsersRepository.
func (u *userRepository) DeleteUser(id int64) error {
	return u.db.Delete(&User{}, "id = ?", id).Error
}

// GetUser implements UsersRepository.
func (u *userRepository) GetUser(id int64) (*User, error) {
	var user User
	err := u.db.First(&user, "id = ?", int(id)).Error
	return &user, err
}

// GetUsers implements UsersRepository.
func (u *userRepository) GetUsers() ([]User, error) {
	var users []User
	err := u.db.Find(&users).Error
	return users, err
}

// UpdateUser implements UsersRepository.
func (u *userRepository) UpdateUser(user *User) error {
	err := u.db.Save(&user).Error
	return err
}

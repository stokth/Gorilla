package users

import "time"

type UsersService interface {
	GetUsers() ([]User, error)
	GetUser(id int64) (*User, error)
	CreateUser(user *User) (*User, error)
	UpdateUser(id int64, user *User) (*User, error)
	DeleteUser(id int64) error
}

type userService struct {
	repo UsersRepository
}

func NewUsersService(repo UsersRepository) UsersService {
	return &userService{repo: repo}
}

// CreateUser implements UsersService.
func (u *userService) CreateUser(user *User) (*User, error) {
	usr := User{
		Email:    user.Email,
		Password: user.Password,
	}

	if err := u.repo.CreateUser(&usr); err != nil {
		return &User{}, err
	}

	return &usr, nil
}

// UpdateUser implements UsersService.
func (u *userService) UpdateUser(id int64, user *User) (*User, error) {
	usr, err := u.repo.GetUser(id)
	if err != nil {
		return &User{}, err
	}

	usr.Email = user.Email
	usr.Password = user.Password
	usr.Updated_At = time.Now()

	if err := u.repo.UpdateUser(usr); err != nil {
		return &User{}, err
	}

	return usr, nil
}

// DeleteUser implements UsersService.
func (u *userService) DeleteUser(id int64) error {
	return u.repo.DeleteUser(id)
}

// GetUser implements UsersService.
func (u *userService) GetUser(id int64) (*User, error) {
	return u.repo.GetUser(id)
}

// GetUsers implements UsersService.
func (u *userService) GetUsers() ([]User, error) {
	return u.repo.GetUsers()
}

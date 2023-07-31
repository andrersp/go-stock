package user

import (
	domain "github.com/andrersp/go-stock/internal/domain/errors"
	"github.com/google/uuid"
)

type userService struct {
	repository UserRepository
}

func NewUserService(repository UserRepository) UserService {
	return &userService{
		repository: repository,
	}
}

func (us *userService) CreateUser(user *User) (*User, error) {

	if userDomain, _ := us.repository.GetUserByUserName(user.userName); userDomain != nil {
		err := domain.NewAppError("RESOURCE_EXISTS", "username exists.")
		return nil, err
	}

	user, err := us.repository.CreateUser(user)

	return user, err
}

func (us *userService) GetUserByID(userId uuid.UUID) (*User, error) {

	user, err := us.repository.GetUserByID(userId)

	if err != nil {
		err = domain.NewAppError("NOT_FOUND", "user not found by id")
		return nil, err
	}

	return user, nil

}

func (us *userService) GetUserByUserName(string) (*User, error) {

	return nil, nil
}

func (us *userService) ListUsers() ([]User, error) {

	return nil, nil
}

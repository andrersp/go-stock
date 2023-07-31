package user

import (
	domain "github.com/andrersp/go-stock/internal/domain/errors"
	"github.com/andrersp/go-stock/internal/utils/security"
	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(*User) (*User, error)
	GetUserByID(uuid.UUID) (*User, error)
	GetUserByUserName(string) (*User, error)
	GetUserByEmail(string) (*User, error)
	ListUsers() ([]User, error)
	Login(string, string) (*User, error)
}

type userService struct {
	repository UserRepository
}

func NewUserService(repository UserRepository) UserService {
	return &userService{
		repository: repository,
	}
}

func (us *userService) CreateUser(user *User) (*User, error) {

	if user, _ := us.repository.GetUserByUserName(user.userName); user != nil {
		err := domain.NewAppError("RESOURCE_EXISTS", "username exists.")
		return nil, err
	}

	if user, _ := us.repository.GetUserByEmail(user.email); user != nil {
		err := domain.NewAppError("RESOURCE_EXISTS", "email exists.")
		return nil, err
	}

	hashedPassword, _ := security.HashGenerator(user.GetPassword())
	user.SetPassword(hashedPassword)

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

func (us *userService) GetUserByUserName(userName string) (*User, error) {

	return us.repository.GetUserByUserName(userName)

}

func (us *userService) GetUserByEmail(email string) (*User, error) {
	return us.repository.GetUserByEmail(email)
}

func (us *userService) ListUsers() ([]User, error) {
	return us.repository.ListUsers()

}

func (us *userService) Login(userName, password string) (*User, error) {

	user, err := us.repository.GetUserByUserName(userName)

	if err != nil {
		err := domain.NewAppError("VALIDATION_ERROR", "invalid username or password")
		return nil, err
	}

	if err := security.CheckPasswordHash(user.GetPassword(), password); err != nil {
		err := domain.NewAppError("VALIDATION_ERROR", "invalid username or password")
		return nil, err
	}

	return user, nil
}

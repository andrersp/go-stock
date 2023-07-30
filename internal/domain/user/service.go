package user

import "github.com/google/uuid"

type userService struct {
	repository UserRepository
}

func NewUserService(repository UserRepository) UserService {
	return &userService{
		repository: repository,
	}
}

func (us *userService) CreateUser(user *User) (*User, error) {

	err := user.Validate()

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (us *userService) GetUserByID(uuid.UUID) (*User, error) {

	return nil, nil

}

func (us *userService) GetUserByUserName(string) (*User, error) {
	return nil, nil
}

func (us *userService) ListUsers() ([]User, error) {

	return nil, nil
}

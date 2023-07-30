package user

import "github.com/google/uuid"

type UserService interface {
	CreateUser(*User) (*User, error)
	GetUserByID(uuid.UUID) (*User, error)
	GetUserByUserName(string) (*User, error)
	ListUsers() ([]User, error)
}

type UserRepository interface {
	CreateUser(*User) (*User, error)
	GetUserByID(uuid.UUID) (*User, error)
	GetUserByUserName(string) (*User, error)
	ListUsers() ([]User, error)
}

package user

import "github.com/google/uuid"

type UserRepository interface {
	CreateUser(*User) (*User, error)
	GetUserByID(uuid.UUID) (*User, error)
	GetUserByUserName(string) (*User, error)
	GetUserByEmail(string) (*User, error)
	ListUsers() ([]User, error)
}

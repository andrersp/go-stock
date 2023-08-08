package schema

import "github.com/google/uuid"

type CreateUserRequest struct {
	UserName        string `json:"userName" validate:"required"`
	Password        string `json:"password" validate:"required,min=6"`
	ConfirmPassword string `json:"confirmPassword" validate:"eqfield=Password"`
	Email           string `json:"email" validate:"required,email"`
} //@name CreateUserRequest

type UserResponse struct {
	ID       uuid.UUID `json:"id"`
	UserName string    `json:"userName"`
	Email    string    `json:"email"`
	Enable   bool      `json:"enable"`
} //@name UserResponse

type CreateUserResponse struct {
	ID uuid.UUID `json:"id"`
} //@name CreateUserResponse

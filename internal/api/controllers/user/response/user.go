package response

import "github.com/google/uuid"

type UserResponse struct {
	ID       uuid.UUID `json:"id"`
	UserName string    `json:"userName"`
	Email    string    `json:"email"`
	Enable   bool      `json:"enable"`
}

type CreateUserResponse struct {
	ID uuid.UUID `json:"id"`
}

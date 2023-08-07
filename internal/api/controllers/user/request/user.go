package request

type CreateUserRequest struct {
	UserName        string `json:"userName" validate:"required"`
	Password        string `json:"password" validate:"required,min=6"`
	ConfirmPassword string `json:"confirmPassword" validate:"eqfield=Password"`
	Email           string `json:"email" validate:"required,email"`
}

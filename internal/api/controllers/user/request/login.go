package request

type LoginRequest struct {
	UserName string `json:"userName" form:"userName" validate:"required"`
	Password string `json:"password" validate:"required"`
}

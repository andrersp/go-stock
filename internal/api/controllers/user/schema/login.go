package schema

type LoginRequest struct {
	UserName string `json:"userName" form:"userName" validate:"required"`
	Password string `json:"password" validate:"required"`
} //@name LoginRequest

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	TokenType    string `json:"tokenType"`
} //@name LoginResponse

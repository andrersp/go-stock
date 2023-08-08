package controler

import (
	"net/http"

	"github.com/andrersp/go-stock/internal/api/controllers/user/schema"
	jwttoken "github.com/andrersp/go-stock/internal/pkg/jwt-token"

	"github.com/andrersp/go-stock/internal/domain/user"
	"github.com/labstack/echo/v4"
)

type userControler struct {
	userService user.UserService
}

func NewUserControler(userService user.UserService) *userControler {

	return &userControler{userService: userService}
}

// @summary Login
// @description "Endpoint login"
// @tags Login
// @param payload body LoginRequest true "Payload"
// @success 200 {object} LoginResponse
// @router /login [post]
func (uc userControler) Login(c echo.Context) error {

	var payload schema.LoginRequest

	if err := c.Bind(&payload); err != nil {
		return c.JSON(400, err)
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(400, err)

	}

	user, err := uc.userService.Login(payload.UserName, payload.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	accessToken, err := jwttoken.CreateAccessToken(user)
	if err != nil {
		return c.String(400, err.Error())
	}

	refreshToken, err := jwttoken.CreateRefreshToken(user)
	if err != nil {
		return c.String(400, err.Error())
	}

	return c.JSON(200, schema.LoginResponse{
		TokenType:    "Bearer",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

// @summary Get Users
// @description Get List of users
// @tags Users
// @security ApiKey
// @success 200 {array} UserResponse
// @failure 400 {object} AppError
// @router /users [get]
func (uc userControler) GetUsers(c echo.Context) error {

	responseDTO := make([]schema.UserResponse, 0)

	users, err := uc.userService.ListUsers()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	for _, u := range users {
		userDTO := schema.UserResponse{
			ID:       u.GetId(),
			UserName: u.GetUserName(),
			Email:    u.GetEmail(),
			Enable:   u.IsEnable(),
		}
		responseDTO = append(responseDTO,
			userDTO,
		)
	}

	return c.JSON(http.StatusOK, responseDTO)

}

// @summary Create User
// @description Create a new user
// @tags Users
// @security ApiKey
// @param payload body CreateUserRequest true "Payload"
// @success 201 {object} CreateUserResponse
// @failure 400 {object} AppError
// @router /users [post]
func (uc userControler) CreateUser(c echo.Context) error {

	var payload schema.CreateUserRequest

	if err := c.Bind(&payload); err != nil {
		return c.JSON(400, err)
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(400, err)
	}

	user, err := user.NewUser(payload.UserName, payload.Password, payload.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user, err = uc.userService.CreateUser(user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	responseDTO := schema.CreateUserResponse{ID: user.GetId()}

	return c.JSON(http.StatusCreated, responseDTO)
}

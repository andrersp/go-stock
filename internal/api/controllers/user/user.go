package controler

import (
	"net/http"

	"github.com/andrersp/go-stock/internal/api/controllers/user/request"
	"github.com/andrersp/go-stock/internal/api/controllers/user/response"
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
// @param payload body request.LoginRequest true "Payload"
// @success 200 {object} response.LoginResponse
// @router /login [post]
func (uc userControler) Login(c echo.Context) error {

	var payload request.LoginRequest

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

	return c.JSON(200, user)
}

// @summary Get Users
// @description Get List of users
// @tags Users
// @security ApiKey
// @success 200 {array} response.UserResponse
// @failure 400 {object} domain.AppError
// @router /users [get]
func (uc userControler) GetUsers(c echo.Context) error {

	responseDTO := make([]response.UserResponse, 0)

	users, err := uc.userService.ListUsers()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	for _, u := range users {
		userDTO := response.UserResponse{
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
// @param payload body request.CreateUserRequest true "Payload"
// @success 201 {object} response.CreateUserResponse
// @failure 400 {object} domain.AppError
// @router /users [post]
func (uc userControler) CreateUser(c echo.Context) error {

	var payload request.CreateUserRequest

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

	responseDTO := response.CreateUserResponse{ID: user.GetId()}

	return c.JSON(http.StatusCreated, responseDTO)
}

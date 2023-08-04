package user

import (
	"github.com/andrersp/go-stock/internal/api/v1/user/request"
	"github.com/andrersp/go-stock/internal/domain/user"
	"github.com/labstack/echo/v4"
)

type userControler struct {
	userService user.UserService
}

func NewUserControler() *userControler {

	return &userControler{}
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

	return c.JSON(200, payload)
}

package routers

import (
	"net/http"

	controler "github.com/andrersp/go-stock/internal/api/controllers/user"
	"github.com/andrersp/go-stock/internal/domain/user"
)

func UserRouters(userService user.UserService) []RouterModel {

	controler := controler.NewUserControler(userService)

	routers := []RouterModel{
		{
			Path:         "/login",
			Func:         controler.Login,
			Method:       http.MethodPost,
			AuthRequired: false,
		},
		{
			Path:         "/users",
			Method:       http.MethodPost,
			Func:         controler.CreateUser,
			AuthRequired: true,
		},
		{
			Path:         "/users",
			Method:       http.MethodGet,
			Func:         controler.GetUsers,
			AuthRequired: true,
		},
	}

	return routers
}

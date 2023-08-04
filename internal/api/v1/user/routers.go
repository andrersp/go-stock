package user

import (
	"net/http"

	"github.com/andrersp/go-stock/internal/utils"
)

func GetUserRouters() []utils.RouterModel {

	controler := NewUserControler()

	routers := []utils.RouterModel{
		{
			Path:         "/login",
			Func:         controler.Login,
			Method:       http.MethodPost,
			AuthRequired: false,
		},
	}

	return routers
}

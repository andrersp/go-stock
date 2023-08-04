package api

import (
	"github.com/andrersp/go-stock/internal/api/v1/user"
	"github.com/andrersp/go-stock/internal/utils"
	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo) {

	v1 := e.Group("/v1")

	routers := make([]utils.RouterModel, 0)

	userRoutres := user.GetUserRouters()

	routers = append(routers, userRoutres...)

	for _, router := range routers {

		v1.Add(router.Method, router.Path, router.Func)

	}

}

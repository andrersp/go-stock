package routers

import (
	"github.com/andrersp/go-stock/internal/config"
	"github.com/andrersp/go-stock/internal/domain/user"
	"github.com/andrersp/go-stock/internal/repository"
	"github.com/labstack/echo/v4"
)

type RouterModel struct {
	Path         string
	Method       string
	Func         func(echo.Context) error
	AuthRequired bool
}

func RegisterRouters(e *echo.Echo) {
	v1 := e.Group("/v1")

	db, err := config.ConnectDB()

	if err != nil {
		panic(err)
	}

	userRepository := repository.NewUserRepository(db)
	userService := user.NewUserService(userRepository)

	routers := make([]RouterModel, 0)
	userRouters := UserRouters(userService)
	routers = append(routers, userRouters...)

	for _, router := range routers {
		v1.Add(router.Method, router.Path, router.Func)
	}

}

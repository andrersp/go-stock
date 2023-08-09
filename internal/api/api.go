package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/andrersp/go-stock/internal/api/routers"
	"github.com/andrersp/go-stock/internal/config"
	"github.com/go-playground/validator"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type RouterModel struct {
	Path         string
	Method       string
	Func         func(echo.Context) error
	AuthRequired bool
}

func StartServer() *http.Server {

	e := echo.New()
	e.Binder = &customBinder{}
	e.Validator = &customValidator{validator: validator.New()}
	e.HideBanner = true

	e.GET("/docs/*", echoSwagger.WrapHandler)

	routers.RegisterRouters(e)

	server := http.Server{
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Addr:           fmt.Sprintf(":%s", config.API_PORT),
		Handler:        e,
	}

	return &server

}

package api

import (
	"log"
	"strings"
	"time"

	"github.com/andrersp/go-stock/internal/api/routers"
	"github.com/go-playground/validator"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/labstack/echo-contrib/echoprometheus"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type RouterModel struct {
	Path         string
	Method       string
	Func         func(echo.Context) error
	AuthRequired bool
}

func StartServer() *echo.Echo {

	e := echo.New()
	e.Binder = &customBinder{}
	e.Validator = &customValidator{validator: validator.New()}
	e.Server.ReadTimeout = 10 * time.Second
	e.Server.WriteTimeout = 10 * time.Second
	e.Server.MaxHeaderBytes = 1 << 20

	customCounter := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "custom_requests_total",
		Help: "How many HTTP requests processed, partitioned by status code and HTTP method.",
	})

	if err := prometheus.Register(customCounter); err != nil {
		log.Fatal(e)
	}

	mwConfig := echoprometheus.MiddlewareConfig{
		Namespace: "stock",
		Skipper: func(c echo.Context) bool {
			return strings.HasPrefix(c.Path(), "/doc") || strings.HasPrefix(c.Path(), "/metrics")
		},
		AfterNext: func(c echo.Context, err error) {
			customCounter.Inc()
		},
	}

	e.Use(echoprometheus.NewMiddlewareWithConfig(mwConfig))
	e.GET("/metrics", echoprometheus.NewHandler())

	e.GET("/docs/*", echoSwagger.WrapHandler)

	routers.RegisterRouters(e)

	return e

}

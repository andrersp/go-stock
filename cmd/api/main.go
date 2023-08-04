package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	_ "github.com/andrersp/go-stock/docs"
	"github.com/andrersp/go-stock/internal/api"
	"github.com/andrersp/go-stock/internal/config"
	"github.com/andrersp/go-stock/internal/utils"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"

	echoSwagger "github.com/swaggo/echo-swagger"
)

func init() {
	config.LoadConfig()
}

// @title Api Stock Azul e Rosa
// @version 1.0

// @accept json
// @produce json

// @BasePath /v1

// @securityDefinitions.apiKey ApiKey
// @in header
// @name Authorization
func main() {

	e := echo.New()
	e.Binder = &utils.CustomBinder{}
	e.Validator = &utils.CustomValidator{Validator: validator.New()}
	e.GET("/docs/*", echoSwagger.WrapHandler)

	api.RegisterPath(e)

	// run server
	go func() {
		address := fmt.Sprintf(":%s", config.API_PORT)
		if err := e.Start(address); err != nil {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// a timeout of 10 seconds to shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	// api.StartServer()
}

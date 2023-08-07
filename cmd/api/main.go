package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	_ "github.com/andrersp/go-stock/docs"
	"github.com/andrersp/go-stock/internal/api"
	"github.com/andrersp/go-stock/internal/config"
	"github.com/andrersp/go-stock/internal/repository"
)

func init() {
	config.LoadConfig()

	err := config.CreateDatabaseConnection()

	if err != nil {
		log.Fatal(err)
	}

	err = repository.Migrate()
	if err != nil {
		log.Fatal(err)
	}

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

	server := api.StartServer()

	// run server
	go func() {
		if err := server.ListenAndServe(); err != nil {

			log.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// a timeout of 10 seconds to shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	// api.StartServer()
}

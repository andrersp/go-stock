package repository

import (
	"github.com/andrersp/go-stock/internal/config"
	userrepository "github.com/andrersp/go-stock/internal/repository/user"
)

func Migrate() error {

	db, err := config.ConnectDB()
	if err != nil {
		return err
	}
	return db.AutoMigrate(&userrepository.User{})
}

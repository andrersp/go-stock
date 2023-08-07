package repository

import (
	"github.com/andrersp/go-stock/internal/config"
)

func Migrate() error {

	db, err := config.ConnectDB()
	if err != nil {
		return err
	}
	return db.AutoMigrate(&UserModel{})
}

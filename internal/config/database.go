package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dbConn *gorm.DB

func CreateDatabaseConnection() error {
	if dbConn != nil {
		closeDbConnection(dbConn)
	}

	if STAGE == DEV_STAGE {
		return createSQLiteConn()
	}

	return createConnectionPostgres()
}

func createSQLiteConn() error {
	db, err := gorm.Open(sqlite.Open("database.sqlite"), &gorm.Config{})
	if err != nil {
		return err
	}
	dbConn = db
	return nil

}

func createConnectionPostgres() (err error) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=America/Sao_Paulo",
		DB_HOST, DB_USER, DB_PASSWD, DB_NAME, DB_PORT, DB_REQUIRED_SSL)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	if err != nil {
		return err
	}

	dbConn = db
	return nil
}

func closeDbConnection(conn *gorm.DB) {

	db, err := conn.DB()

	if err != nil {
		return
	}

	defer db.Close()
}

func ConnectDB() (db *gorm.DB, err error) {
	sqlDb, err := dbConn.DB()

	if sqlDb == nil {
		return
	}

	if err = sqlDb.Ping(); err != nil {
		log.Fatal(err)
	}

	db = dbConn
	return
}

package config

import (
	"fmt"

	"github.com/kamilanindita/go-grpc/server/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(configuration Config) *gorm.DB {

	dbUsername := configuration.Get("DB_USERNAME")
	dbPassword := configuration.Get("DB_PASSWORD")
	dbHost := configuration.Get("DB_HOST")
	dbPort := configuration.Get("DB_PORT")
	dbName := configuration.Get("DB_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Failed on connecting to the database server")
	}

	migrateDDL(db)

	return db
}

func migrateDDL(db *gorm.DB) {
	db.AutoMigrate(
		model.BookDB{},
	)
}

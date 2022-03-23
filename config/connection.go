package config

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"main.go/models"
	"main.go/models/entity/product"
	"main.go/models/entity/user"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), //IO Writter
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL ThreshOld
			LogLevel:                  logger.Info, //Log Level
			IgnoreRecordNotFoundError: true,        // Ignore ErrorRecordNotFound Error for logger
			Colorful:                  false,
		},
	)

	dsn := "root:@tcp(127.0.0.1:3306)/user-backend?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("Failed To Connect Database")
	}
	DB.AutoMigrate(
		&user.User{},
		&product.Product{},
		&models.Role{},
	)
	return DB
}

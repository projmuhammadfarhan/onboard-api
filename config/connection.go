package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"main.go/models"
	"main.go/models/entity/product"
	"main.go/models/entity/user"
)

var (
	DB  *gorm.DB
	err error
)

type db struct {
	DB_User     string
	DB_Password string
	DB_Port     string
	DB_Name     string
}

func Connect() *gorm.DB {
	// Definisi Koneksi Database
	var mysqlConfig = &db{
		DB_User:     "root",
		DB_Password: "",
		DB_Port:     "3306",
		DB_Name:     "user-backend",
	}

	// Log info Koneksi
	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags), //IO Writter
	// 	logger.Config{
	// 		SlowThreshold:             time.Second, // Slow SQL ThreshOld
	// 		LogLevel:                  logger.Info, //Log Level
	// 		IgnoreRecordNotFoundError: true,        // Ignore ErrorRecordNotFound Error for logger
	// 		Colorful:                  false,
	// 	},
	// )

	dsn := mysqlConfig.DB_User + ":" + mysqlConfig.DB_Password + "@tcp(127.0.0.1:" + mysqlConfig.DB_Port + ")/" + mysqlConfig.DB_Name + "?charset=utf8mb4&parseTime=True&loc=Local"

	// DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("Failed To Connect Database")
	}
	// Auto Migrate Table
	DB.AutoMigrate(
		&user.User{},
		&product.Product{},
		&models.Role{},
	)
	return DB
}

package config

import (
	"fmt"

	"github.com/Jiran03/agmc/task/day3/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     int
	DB_HOST     string
	DB_NAME     string
}

func InitDB() {
	config := Config{
		DB_USERNAME: "root",
		DB_PASSWORD: "mysecretpw",
		DB_PORT:     3300,
		DB_HOST:     "localhost",
		DB_NAME:     "task_day3",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString))
	if err != nil {
		panic(err)
	}
}

func InitMigration() {
	DB.AutoMigrate(&models.User{})
}

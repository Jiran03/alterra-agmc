package config

import (
	"fmt"
	"os"
	"strconv"

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
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	config := Config{
		DB_USERNAME: os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASS"),
		DB_PORT:     dbPort,
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_NAME:     os.Getenv("DB_NAME"),
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

package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Jiran03/agmc/task/day4/models"
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

	DB_USERNAME_TEST string
	DB_PASSWORD_TEST string
	DB_HOST_TEST     string
	DB_PORT_TEST     int
	DB_NAME_TEST     string
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

//testing
func InitDBTest() {
	// config := Config{
	// 	DB_USERNAME_TEST: "root",
	// 	DB_PASSWORD_TEST: "mysecretpw",
	// 	DB_HOST_TEST:     "localhost",
	// 	DB_PORT_TEST:     3300,
	// 	DB_NAME_TEST:     "task_day4_test",
	// }
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT_TEST"))
	config := Config{
		DB_USERNAME_TEST: os.Getenv("DB_USER_TEST"),
		DB_PASSWORD_TEST: os.Getenv("DB_PASS_TEST"),
		DB_PORT_TEST:     dbPort,
		DB_HOST_TEST:     os.Getenv("DB_HOST_TEST"),
		DB_NAME_TEST:     os.Getenv("DB_NAME_TEST"),
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_USERNAME_TEST,
		config.DB_PASSWORD_TEST,
		config.DB_HOST_TEST,
		config.DB_PORT_TEST,
		config.DB_NAME_TEST,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString))
	if err != nil {
		panic(err)
	}
}

func InitMigrationTest() {
	DB.Migrator().DropTable(&models.User{})
	DB.AutoMigrate(&models.User{})
}

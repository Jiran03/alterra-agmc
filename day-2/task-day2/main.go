package main

import (
	"log"

	"github.com/Jiran03/agmc/task/day2/config"
	"github.com/Jiran03/agmc/task/day2/routes"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.InitDB()
	config.InitMigration()
}

func main() {
	e := routes.New()

	e.Start(":8080")
	e.Logger.Fatal(e.Start(":8080"))
}

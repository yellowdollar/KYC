package main

import (
	"KYC/iternals/configs"
	"KYC/iternals/controller"
	"KYC/iternals/db"
	"KYC/logger"
	"log"
)

func main() {

	// Reading app configs
	if err := configs.ReadProjectSettings(); err != nil {
		log.Fatalf("error while reading settings: %s", err)
	}

	// Init logger
	if err := logger.Init(); err != nil {
		log.Fatalf("logger init error: %s", err)
	}
	logger.Info.Println("Logger established succesfully")

	// Database connection
	if _, err := db.ConnDB(); err != nil {
		logger.Error.Println("error while connecting to database")
	}
	logger.Info.Println("Connection to database established successfully")

	// Migrate table
	if err := db.InitMigrations(); err != nil {
		logger.Error.Println("error while making migrations")
	}
	logger.Info.Println("Migrations completed successfully")

	// running server
	err := controller.RunServer()
	if err != nil {
		logger.Error.Println("error while running server")
	}
	logger.Info.Println("Server started successfully")
}

package main

import (
	"gin/application/errorsCodes"
	"gin/domain/entities"
	"gin/infrastructure/database"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	loadEnvironmentVariables()

	ensureDatabaseExists()

	database := database.ConnectToDB()
	dbContext := database.GetDBContext()

	dbContext.AutoMigrate(&entities.User{})
	log.Println("Migration Complete.")
}

func loadEnvironmentVariables() {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func ensureDatabaseExists() {

	serverName := os.Getenv("SERVER_STRING")
	databaseName := os.Getenv("DATABASE_NAME")

	server, err := gorm.Open(postgres.Open(serverName), &gorm.Config{})
	if err != nil {
		log.Fatal(errorsCodes.DATABASE_CONNECTION_ERROR, err)
	}

	var result int64
	server.Raw("SELECT 1 FROM pg_database WHERE datname = ?", databaseName).Count(&result)

	if result == 0 {
		createDatabase(server, databaseName)
	}
}

func createDatabase(server *gorm.DB, databaseName string) {

	err := server.Exec("CREATE DATABASE " + databaseName).Error

	if err != nil {
		log.Fatal(errorsCodes.DATABASE_NOT_CREATED, err)
	}
}

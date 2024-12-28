package integration

import (
	"gin/application/utility"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func EnsureDatabaseExists(t *testing.T) *gorm.DB {

	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	serverName := os.Getenv("SERVER_STRING_TEST")
	databaseName := os.Getenv("DATABASE_NAME_TEST")

	server, err := gorm.Open(postgres.Open(serverName), &gorm.Config{})
	if err != nil {
		t.Fatal(*utility.DatabaseConnectionError.WithDescription(err.Error()))
	}

	var result int64
	server.Raw("SELECT 1 FROM pg_database WHERE datname = ?", databaseName).Count(&result)

	if result == 0 {
		createDatabase(server, databaseName, t)
	}

	return server
}

func createDatabase(server *gorm.DB, databaseName string, t *testing.T) {

	err := server.Exec("CREATE DATABASE " + databaseName).Error

	if err != nil {
		t.Fatal(*utility.DatabaseCreationError.WithDescription(err.Error()))
	}
}

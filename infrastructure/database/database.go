package database

import (
	"gin/application/errorsCodes"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	dbContext *gorm.DB
}

func (database *Database) GetDBContext() *gorm.DB {
	return database.dbContext
}

func ConnectToDB() *Database {
	dsn := os.Getenv("CONNECTION_STRING")

	dbContext, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil || dbContext == nil {
		log.Fatal(errorsCodes.DATABASE_CONNECTION_ERROR, err)
	}

	return &Database{dbContext: dbContext}
}

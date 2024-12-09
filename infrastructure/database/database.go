package database

import (
	"gin/application/utility"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DBContext *gorm.DB
}

func NewDatabase() (*Database, *utility.ErrorCode) {

	dsn := os.Getenv("CONNECTION_STRING")

	context, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil || context == nil {
		return nil, utility.DatabaseConnectionError.WithDescription(err.Error())
	}

	return &Database{DBContext: context}, nil
}

func (database *Database) GetDBContext() *gorm.DB {
	return database.DBContext
}

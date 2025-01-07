package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DBContext *gorm.DB
}

// NewDatabase establishes a database connection
func NewDatabase() (*Database, error) {

	connection := os.Getenv("CONNECTION_STRING")

	context, err := gorm.Open(postgres.Open(connection), &gorm.Config{})

	if err != nil || context == nil {
		return nil, err
	}

	return &Database{DBContext: context}, nil
}

// GetDBContext returns the current database context.
func (database *Database) GetDBContext() *gorm.DB {
	return database.DBContext
}

package database

import "gorm.io/gorm"

type IDatabase interface {
	GetDBContext() *gorm.DB
}

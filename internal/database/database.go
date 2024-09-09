package database

import (
	"go-boilerplate/configs"

	"gorm.io/gorm"
)

type Database interface {
	Connect(cfg configs.Database) error
	Close() error
	GetDB() *gorm.DB
}

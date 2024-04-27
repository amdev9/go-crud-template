package models

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dbName string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	DB = db
	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Book{})
}

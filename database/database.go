package database

import (
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"

	"Login-api/models"
)
var (
	DB  *gorm.DB
)

func ConectaNoBD() {
	db, err := gorm.Open(sqlite.Open("login.db"), &gorm.Config{})
  	if err != nil {
    	panic("failed to connect database")
  	}

	DB = db

	db.AutoMigrate(&models.User{})
}

func GetDataBase() *gorm.DB {
	return DB
}
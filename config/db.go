package config

import (
	"go-rest-api/models"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func InitDB() {
	var err error
	db, err = gorm.Open("postgres", os.Getenv("DATABASE_CONNECTION"))

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Article{})
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	db.Close()
}

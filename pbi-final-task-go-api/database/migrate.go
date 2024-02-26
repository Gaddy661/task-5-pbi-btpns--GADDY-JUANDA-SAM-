package database

import (
	"pbi-final-task-go-api/models"
)

func Migrate() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Photo{})
}

package database

import (
	"gorm.io/gorm"
	"scoresystem/app/models"
)

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Student{},
	)
}

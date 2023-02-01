package database

import (
	"gorm.io/gorm"
	"scoresystem/app/models"
	"scoresystem/app/models/modules"
)

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&modules.Pe{},
		&modules.Art{},
		&modules.GPA{},
		&modules.Innovate{},
		&modules.Labour{},
		&modules.Moral{},
		&models.Application{},
	)
}

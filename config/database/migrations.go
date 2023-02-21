package database

import (
	"gorm.io/gorm"
	"scoresystem/app/models"
	"scoresystem/app/models/modules"
)

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Score{},
		&modules.Pe{},
		&modules.Art{},
		&modules.GPA{},
		&modules.Innovate{},
		&modules.Labour{},
		&modules.Moral{},
		&models.Application{},
		&models.Reason{},
		&models.SquareTopic{},
		&models.Admin{},
		&models.Advice{},
		&models.Appeal{},
	)
}

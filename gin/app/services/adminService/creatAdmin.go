package adminService

import (
	"scoresystem/app/models"
	"scoresystem/config/database"
)

func CreateAdmin(user models.User) error {
	name := user.Name
	ID := user.ID
	admin1 := &models.Admin{
		AdminName: name,
		AdminID:   ID,
	}

	result := database.DB.Create(admin1)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

package userService

import (
	"scoresystem/app/models"
	"scoresystem/app/utility"
	"scoresystem/config/database"
)

func UpdatePasswordByID(user *models.User, password string) error {
	pass := utility.Encryrpt(password)
	user.Password = pass
	err := database.DB.Model(models.User{}).Where(
		models.User{
			ID: user.ID,
		}).Updates(user).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}

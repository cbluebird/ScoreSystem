package userService

import (
	"scoresystem/app/models"
	"scoresystem/config/database"
)

func GetUserById(id int) (*models.User, error) {
	student := models.User{}
	result := database.DB.Where(
		&models.User{
			ID: id,
		},
	).First(&student)
	if result.Error != nil {
		return nil, result.Error
	}
	return &student, nil
}

func GetUserByAccount(account string) (*models.User, error) {
	user := models.User{}
	result := database.DB.Where(
		&models.User{
			Account: account,
		},
	).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

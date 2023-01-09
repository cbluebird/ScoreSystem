package studentService

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

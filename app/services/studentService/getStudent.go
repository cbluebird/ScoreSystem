package studentService

import (
	"scoresystem/app/models"
	"scoresystem/config/database"
)

func GetStudentById(id int) (*models.Student, error) {
	student := models.Student{}
	result := database.DB.Where(
		&models.Student{
			ID: id,
		},
	).First(&student)
	if result.Error != nil {
		return nil, result.Error
	}
	return &student, nil
}

package userService

import (
	"scoresystem/app/models"
	"scoresystem/app/utility"
	"scoresystem/config/database"
)

func CreateUser(user models.Reg) error {
	pass := utility.Encryrpt(user.Password)
	user1 := &models.User{
		Age:        user.Age,
		Password:   pass,
		Account:    user.Account,
		Admin:      user.Admin,
		Profession: user.Profession,
		Name:       user.Name,
	}
	result := database.DB.Create(user1)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

package userService

import (
	"scoresystem/app/models"
	"scoresystem/app/utility"
)

func CheckUserByPassword(user *models.User, password string) bool {
	pass := utility.Encryrpt(user.Password)
	if user.Password != pass {
		return false
	} else {
		return true
	}
}

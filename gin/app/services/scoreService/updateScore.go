package scoreService

import (
	"scoresystem/app/models"
	"scoresystem/config/database"
)

func UpdateScore(userid, year int) {
	var s models.Score
	result := database.DB.Where(models.Score{
		Year:   year,
		Userid: userid,
	}).First(&s)
	var score float32
	list, _ := GetDetail(userid, year)
	for _, value := range list {
		score += value.Score
	}
	if result.RowsAffected == 0 {
		var newScore models.Score
		newScore.Year = year
		newScore.Userid = userid
		newScore.Grade = score
		result = database.DB.Create(&newScore)
	} else {
		s.Grade = score
		result = database.DB.Where(models.Score{
			Year:   year,
			Userid: userid,
		}).Updates(&s)
	}
}

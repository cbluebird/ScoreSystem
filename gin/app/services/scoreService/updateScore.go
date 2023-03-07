package scoreService

import (
	"fmt"
	"log"
	"scoresystem/app/engine"
	"scoresystem/app/models"
	"scoresystem/config/database"
	"strconv"
)

func UpdateScore(userid, year int) {
	var s models.Score
	result := database.DB.Where(models.Score{
		Year:   year,
		Userid: userid,
	}).First(&s)
	//var score float64
	list, _ := GetDetail(userid, year)
	user, _ := GetUserById(userid)
	var rule models.Rule
	_ = database.DB.Where(
		&models.Rule{Collage: user.Profession},
	).First(&rule)
	m := make(map[int]float64, 6)
	for _, v := range list {
		m[v.Module] += v.Score
	}
	value := make(map[string]interface{})
	for i := 1; i <= 6; i++ {
		value[models.GetModule[i]] = m[i]
	}
	//for _, value := range list {
	//	score += value.Score
	//}
	engine := engine.NewEngine(rule.Rule)
	err := engine.Calculate(value)
	if err != nil {
		log.Println(err)
		return
	}
	score, _ := engine.GetVal()
	log.Println(engine.GetVal())
	if result.RowsAffected == 0 {
		var newScore models.Score
		newScore.Year = year
		newScore.Userid = userid
		newScore.Grade = Decimal(score.(float64))
		result = database.DB.Create(&newScore)
	} else {
		s.Grade = Decimal(score.(float64))
		result = database.DB.Where(models.Score{
			Year:   year,
			Userid: userid,
		}).Updates(&s)
	}
}

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

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

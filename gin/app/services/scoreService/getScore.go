package scoreService

import (
	"log"
	"scoresystem/app/models"
	"scoresystem/app/models/modules"
	"scoresystem/config/database"
)

func GetAllScore(userid, year int) (models.Score, error) {
	var score models.Score
	result := database.DB.Where(
		models.Score{
			Userid: userid,
			Year:   year,
		}).First(&score)
	log.Println(score)
	if result.Error != nil {
		return score, result.Error
	}
	return score, nil
}

func GetDetail(userid, year int) ([]modules.Detail, error) {
	var list []modules.Detail
	var list1 []modules.Detail
	list1, _ = GetArt(userid, year)
	list = append(list, list1...)
	list1, _ = GetGpa(userid, year)
	list = append(list, list1...)
	list1, _ = GetPe(userid, year)
	list = append(list, list1...)
	list1, _ = GetInnovate(userid, year)
	list = append(list, list1...)
	list1, _ = GetLabour(userid, year)
	list = append(list, list1...)
	list1, _ = GetMoral(userid, year)
	list = append(list, list1...)
	return list, nil
}

func GetArt(userid, year int) ([]modules.Detail, error) {
	var origin []modules.Art
	result := database.DB.Where(
		modules.Art{
			Userid: userid,
			Age:    year,
		}).Find(&origin)
	if result.Error != nil {
		return nil, result.Error
	}
	list := make([]modules.Detail, len(origin))
	for i, value := range origin {
		list[i].Score = value.Score
		list[i].Class = value.Class
		list[i].Module = 1
		list[i].Description = value.Description
	}
	return list, nil
}
func GetGpa(userid, year int) ([]modules.Detail, error) {
	var origin []modules.GPA
	result := database.DB.Where(
		modules.GPA{
			Userid: userid,
			Age:    year,
		}).Find(&origin)
	if result.Error != nil {
		return nil, result.Error
	}
	list := make([]modules.Detail, len(origin))
	for i, value := range origin {
		list[i].Score = value.Score
		list[i].Class = value.Class
		list[i].Module = 2
		list[i].Description = value.Description
	}
	return list, nil
}
func GetInnovate(userid, year int) ([]modules.Detail, error) {
	var origin []modules.Innovate
	result := database.DB.Where(
		modules.Innovate{
			Userid: userid,
			Age:    year,
		}).Find(&origin)
	if result.Error != nil {
		return nil, result.Error
	}
	list := make([]modules.Detail, len(origin))
	for i, value := range origin {
		list[i].Score = value.Score
		list[i].Class = value.Class
		list[i].Module = 3
		list[i].Description = value.Description
	}
	return list, nil
}
func GetLabour(userid, year int) ([]modules.Detail, error) {
	var origin []modules.Labour
	result := database.DB.Where(
		modules.Labour{
			Userid: userid,
			Age:    year,
		}).Find(&origin)
	if result.Error != nil {
		return nil, result.Error
	}
	list := make([]modules.Detail, len(origin))
	for i, value := range origin {
		list[i].Score = value.Score
		list[i].Class = value.Class
		list[i].Module = 4
		list[i].Description = value.Description
	}
	return list, nil
}
func GetMoral(userid, year int) ([]modules.Detail, error) {
	var origin []modules.Moral
	result := database.DB.Where(
		modules.Moral{
			Userid: userid,
			Age:    year,
		}).Find(&origin)
	if result.Error != nil {
		return nil, result.Error
	}
	list := make([]modules.Detail, len(origin))
	for i, value := range origin {
		list[i].Score = value.Score
		list[i].Class = value.Class
		list[i].Module = 5
		list[i].Description = value.Description
	}
	return list, nil
}
func GetPe(userid, year int) ([]modules.Detail, error) {
	var origin []modules.Pe
	result := database.DB.Where(
		modules.Pe{
			Userid: userid,
			Age:    year,
		}).Find(&origin)
	if result.Error != nil {
		return nil, result.Error
	}
	list := make([]modules.Detail, len(origin))
	for i, value := range origin {
		list[i].Score = value.Score
		list[i].Class = value.Class
		list[i].Module = 6
		list[i].Description = value.Description
	}
	return list, nil
}

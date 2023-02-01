package models

import "time"

// 存放学生提交的申请
type Application struct {
	ID          int `gorm:"primary_key;AUTO_INCREMENT"`
	Age         int
	Module      int
	Class       int
	CreateTime  time.Time
	Userid      int
	Score       float32
	Description string
	Status      int
	AdminId     int
}
type ApplicationData struct {
	Module      int     `json:"module"`
	Age         int     `json:"age"`
	Class       int     `json:"class"`
	Description string  `json:"description"`
	Score       float32 `json:"score"`
	Admin       int     `json:"admin"`
}

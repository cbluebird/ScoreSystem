package models

import "time"

type SquareTopic struct {
	SquareId int `gorm:"primary_key;AUTO_INCREMENT"`
	UserId   int
	Word     string
	Time     time.Time
}

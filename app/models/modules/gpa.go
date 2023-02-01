package modules

import "time"

type GPA struct {
	ID          int `gorm:"primary_key;AUTO_INCREMENT"`
	Age         int
	Class       int
	CreateTime  time.Time
	Userid      int
	Score       float32
	Description string
	Type        int
}

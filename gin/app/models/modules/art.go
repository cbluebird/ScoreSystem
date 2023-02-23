package modules

import (
	"time"
)

//美育素质的储存

type Art struct {
	ID            int `gorm:"primary_key;AUTO_INCREMENT"`
	Age           int
	Class         int
	CreateTime    time.Time
	Userid        int
	Score         float32
	Description   string
	Type          int
	ApplicationID int
}

//type表示是不是唯一的
//class表示具体的哪一项

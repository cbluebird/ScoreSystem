package modules

import "time"

type Moral struct {
	ID            int `gorm:"primary_key;AUTO_INCREMENT"`
	Age           int
	Class         int
	CreateTime    time.Time
	Userid        int `gorm:"uniqueIndex"`
	Score         float32
	Description   string
	Type          int
	ApplicationID int
}

package modules

import "time"

type Innovate struct {
	ID            int `gorm:"primary_key;AUTO_INCREMENT"`
	Age           int `gorm:"index:idx_member"`
	Class         int
	CreateTime    time.Time
	Userid        int `gorm:"index:idx_member"`
	Score         float64
	Description   string
	Type          int
	ApplicationID int
}

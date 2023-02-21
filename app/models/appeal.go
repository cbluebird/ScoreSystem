package models

import "time"

type Appeal struct {
	ID            int    `gorm:"primary_key;AUTO_INCREMENT"`
	ApplicationID int    `json:"applicationid"`
	AppealReason  string `json:"appealreason"`
	CreatTime     time.Time
}

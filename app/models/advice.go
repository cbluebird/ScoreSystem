package models

import "time"

type Advice struct {
	Id         int
	Suggestion string
	AdminId    int
	UserName   string
	Time       time.Time
	Anon       string
}

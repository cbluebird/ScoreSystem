package models

import "time"

type Rule struct {
	Rule    	string
	CreateTime 	time.Time
	ID         	int `gorm:"primary_key;AUTO_INCREMENT"`
	Collage 	string
}
var GetModule map[int]string = map[int]string{1: "art", 2: "gpa", 3: "innovate", 4: "labour", 5: "moral", 6: "pe"}
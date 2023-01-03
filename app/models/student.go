package models

type Student struct {
	StudentId string
	ID        int `gorm:"primary_key;AUTO_INCREMENT"`
	Password  string
}

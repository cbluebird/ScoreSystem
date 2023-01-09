package models

type User struct {
	StudentId  string
	ID         int `gorm:"primary_key;AUTO_INCREMENT"`
	Password   string
	Admin      int //判断是不是辅导员，值为0就是学生，值为1就是辅导员
	Age        int
	Profession string
}

//用于登录

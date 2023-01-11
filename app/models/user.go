package models

type User struct {
	Account    string
	ID         int `gorm:"primary_key;AUTO_INCREMENT"`
	Password   string
	Admin      int //判断是不是辅导员，值为0就是学生，值为1就是辅导员
	Age        int
	Profession string
	Name       string
}

//用于登录

type Reg struct {
	Account    string `json:"account"`
	Password   string `json:"password"`
	Admin      int    `json:"admin"` // 判断是不是辅导员，值为0就是学生，值为1就是辅导员
	Age        int    `json:"age"`
	Profession string `json:"profession"`
	Name       string `json:"name"`
}

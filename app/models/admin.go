package models

//辅导员的管理权限
//8个字段，6个模块加上姓名和id

type Admin struct {
	Art       int
	Gpa       int
	Innovate  int
	Labour    int
	Moral     int
	PE        int
	AdminName string
	AdminID   int
}

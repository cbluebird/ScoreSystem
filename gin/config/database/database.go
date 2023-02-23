package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"scoresystem/config/config"
)

var DB *gorm.DB

func Init() { // 初始化数据库
	user := config.Config.GetString("db.user")
	pass := config.Config.GetString("db.password")
	port := config.Config.GetString("db.address")
	name := config.Config.GetString("db.name")
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		pass,
		port,
		name,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Panicln("Database Error: ", err)
	} else {
		fmt.Printf("database start")
	}
	err = autoMigrate(db)
	if err != nil {
		log.Fatal("DatabaseMigrateFailed", err)
	}
	DB = db
}

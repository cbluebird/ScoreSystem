package config

import (
	"github.com/spf13/viper"
	"log"
	"scoresystem/app/models"
)

var Config = viper.New()

func InitConfig() {
	Config.SetConfigName("config")
	Config.SetConfigType("yaml")
	Config.AddConfigPath(".")
	Config.WatchConfig() // 自动将配置读入Config变量
	err := Config.ReadInConfig()
	if err != nil {
		log.Fatal("Config not find", err)
	}
	models.Authority[1][1] = 0
	models.Authority[1][2] = 0
	models.Authority[2][1] = 1
	models.Authority[3][1] = 0
	models.Authority[3][2] = 0
	models.Authority[3][3] = 0
	models.Authority[3][4] = 0
	models.Authority[4][1] = 1
	models.Authority[4][2] = 0
	models.Authority[4][3] = 0
	models.Authority[4][4] = 0
	models.Authority[4][5] = 0
	models.Authority[5][1] = 1
	models.Authority[5][2] = 0
	models.Authority[5][3] = 0
	models.Authority[5][4] = 0
	models.Authority[5][5] = 0
	models.Authority[5][6] = 0
	models.Authority[6][1] = 0
	models.Authority[6][2] = 0
	models.Authority[6][3] = 0
	models.Authority[6][4] = 0
}

package startConfig

import (
	"scoresystem/config/config"
	"scoresystem/config/database"
	"scoresystem/config/redis"
)

func Init() {
	config.InitConfig()
	database.Init()
	redis.Init()
}

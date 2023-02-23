package session

import (
	"github.com/gin-contrib/sessions"
	sessionRedis "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"scoresystem/config/redis"
)

func setRedis(r *gin.Engine, name string) {
	Info := redis.RedisInfo
	store, _ := sessionRedis.NewStore(10, "tcp", Info.Host+":"+Info.Port, "", []byte("secret"))
	r.Use(sessions.Sessions(name, store))
}

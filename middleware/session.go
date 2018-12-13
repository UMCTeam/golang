package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-contrib/sessions"
	"myapp/conf"
)

func SessionStore() gin.HandlerFunc{
	//设置session存储
	store, _ := redis.NewStore(10, "tcp",
		conf.RedisSetting.String(),
		conf.RedisSetting.Password,
		[]byte(conf.MyAppSetting.Tag))

	return sessions.Sessions(conf.MyAppSetting.Tag + "_session", store)
}

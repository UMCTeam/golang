package main

import (
	_ "github.com/go-sql-driver/mysql"
	"myapp/conf"
	"myapp/lib"
	"myapp/model"
	"myapp/redis"
	"myapp/router"
)

func main() {
	//应用设置
	conf.Step()

	//初始化数据库
	model.InitDB("mysql", conf.MySqlSetting.String())

	//初始化Redis
	redis.InitRedis(conf.RedisSetting.String(), conf.RedisSetting.Password, conf.RedisSetting.Database)

	//初始化日志
	lib.InitLogger()

	//初始化路由
	r, _ := router.InitRouter()

	//启动服务
	r.Run(conf.MyAppSetting.Host + ":" + conf.MyAppSetting.Port)
}
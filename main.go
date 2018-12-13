package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
	"myapp/conf"
	"myapp/model"
	"myapp/redis"
	"myapp/router"
	"os"
)

func main() {
	//应用设置
	conf.Step()

	//初始化数据库
	_, err := model.InitDB("mysql", conf.MySqlSetting.String())
	if err != nil {
		log.Fatal(err)
	}

	//初始化缓存服务
	redis.InitRedis(conf.RedisSetting.String(), conf.RedisSetting.Password, conf.RedisSetting.Database)

	//配置日志
	if !gin.IsDebugging() {
		gin.DisableConsoleColor()
		f, _ := os.Create(conf.MyAppSetting.Tag + ".log")
		gin.DefaultWriter = io.MultiWriter(f)
	}

	//初始化路由
	r, _ := router.InitRouter()
	//启动服务
	r.Run(conf.MyAppSetting.Host + ":" + conf.MyAppSetting.Port)
}
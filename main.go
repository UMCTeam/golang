package main

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"myapp/conf"
	"myapp/model"
	"myapp/router"
)

func main() {
	//应用设置
	conf.Step()

	//初始化数据库
	conn := conf.MySqlSetting.User +":" + conf.MySqlSetting.Password  + "@tcp(" + conf.MySqlSetting.Host + ":" + conf.MySqlSetting.Port +")/myapp"
	err := model.InitDB("mysql", conn)
	if err != nil {
		log.Fatal(err)
	}

	//初始化路由
	r, _ := router.InitRouter()
	r.Run(conf.MyAppSetting.Host + ":" + conf.MyAppSetting.Port)
}
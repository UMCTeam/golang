package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"myapp/model"
	"net/http"
)

type MyApp struct {
	Host string
	Port string
}

var MyAppSetting = &MyApp{}

type MySql struct {
	Host     string
	Port     string
	User     string
	Password string
}

var MySqlSetting = &MySql{}

var cfg *ini.File

func Step()  {
	var err error
	cfg, err = ini.Load("conf/default.ini")

	if err != nil {
		log.Print(err)
	}

	cfg.Section("myapp").MapTo(MyAppSetting)
	cfg.Section("mysql").MapTo(MySqlSetting)
}

func main() {
	//应用设置
	Step()

	//初始化数据库
	conn := MySqlSetting.User +":" + MySqlSetting.Password  + "@tcp(" + MySqlSetting.Host + ":" + MySqlSetting.Port +")/myapp"
	err := model.InitDB("mysql", conn)
	if err != nil {
		log.Fatal(err)
	}

	//服务配置
	router:= gin.Default()
	//配置模板目录
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "./index")
	})

	router.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.tmpl", gin.H{
			"notice": "暂不开放注册",
		})
	})

	router.Run(MyAppSetting.Host + ":" + MyAppSetting.Port)
}
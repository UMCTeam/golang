package conf

import (
	"github.com/go-ini/ini"
	"log"
	"os"
)

type MyApp struct {
	Host string
	Port string
	Storage string
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

	//配置存储目录
	if _, err := os.Stat(MyAppSetting.Storage); err != nil {
		os.MkdirAll("/data", 755)
	}
}
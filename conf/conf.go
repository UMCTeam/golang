package conf

import (
	"github.com/go-ini/ini"
	"log"
	"os"
)

type MyApp struct {
	Tag     string
	Host    string
	Port    string
	Storage string
}

var MyAppSetting = &MyApp{}

type MySql struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func (c *MySql) String() string {
	mySqlAddr := c.User
	mySqlAddr += ":" + c.Password
	mySqlAddr += "@tcp(" + c.Host + ":" + c.Port +")/"
	mySqlAddr += c.Database

	return mySqlAddr
}

var MySqlSetting = &MySql{}

type Redis struct {
	Host     string
	Password string
	Port     string
	Database string
}

func (c *Redis) String() string {
	redisAddr := c.Host + ":" + c.Port
	return redisAddr
}

var RedisSetting = &Redis{}

var cfg *ini.File

func Step()  {
	var err error

	cfg, err = ini.Load("conf/default.ini")
	if err != nil {
		log.Print(err)
	}

	//获取配置参数
	cfg.Section("myapp").MapTo(MyAppSetting)
	cfg.Section("mysql").MapTo(MySqlSetting)
	cfg.Section("redis").MapTo(RedisSetting)

	//配置存储目录
	if _, err := os.Stat(MyAppSetting.Storage); err != nil {
		os.MkdirAll("/data", 755)
	}
}


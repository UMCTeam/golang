package lib

import (
	"github.com/gin-gonic/gin"
	"io"
	"myapp/conf"
	"os"
)

func InitLogger()  {
	//配置日志
	if !gin.IsDebugging() {
		gin.DisableConsoleColor()
		f, _ := os.Create(conf.MyAppSetting.Tag + ".log")
		gin.DefaultWriter = io.MultiWriter(f)
	}
}
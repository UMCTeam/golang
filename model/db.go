package model

import (
	"github.com/go-xorm/xorm"
	"log"
	"time"
)

var db *xorm.Engine

func InitDB(driver string, addr string){
	var err error
	//连接数据库
	db, err = xorm.NewEngine(driver, addr)

	//改变时区
	db.TZLocation, _ = time.LoadLocation("Asia/shanghai")

	//打印执行的sql命令
	db.ShowSQL(true)

	//同步数据库结构
	err = db.Sync2(new(User))

	if err != nil {
		log.Fatal(err)
	}
}
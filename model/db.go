package model

import (
	"github.com/go-xorm/xorm"
	"time"
)

var db *xorm.Engine

func InitDB(driver string, conn string) error{
	var err error
	//连接数据库
	db, err = xorm.NewEngine(driver, conn)
	//改变时区
	db.TZLocation, _ = time.LoadLocation("Asia/shanghai")

	if err != nil {
		return  err
	}

	db.ShowSQL(true)
	return db.Sync2(new(User))
}
package model

import (
	"github.com/go-xorm/xorm"
	"time"
)

var db *xorm.Engine

func InitDB(driver string, addr string) (*xorm.Engine, error){
	var err error
	//连接数据库
	db, err = xorm.NewEngine(driver, addr)
	if err != nil {
		return nil, err
	}

	//改变时区
	db.TZLocation, _ = time.LoadLocation("Asia/shanghai")

	//打印执行的sql命令
	db.ShowSQL(true)

	//同步数据库结构
	if err := db.Sync2(new(User)); err != nil {
		return nil, err
	}

	return db, nil
}
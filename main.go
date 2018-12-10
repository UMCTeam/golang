package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_"github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	//数据库配置
	db, err := sql.Open("mysql", "root:lx85572134@tcp(192.168.51.183:3500)/myapp")
	if err != nil{
		log.Fatal(err)
	}
    //main 函数结束时调用
	defer db.Close()

	//设置数据库连接数量
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)

	//检查数据连接
	if err := db.Ping(); err != nil {
		log.Print(err)
	}

	//服务配置
	r:= gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "hello world",
			"result" : gin.H{
				"username": "chendaole",
			},
		})
	})

	r.Run("0.0.0.0:8081")
}
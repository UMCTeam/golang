package router

import (
	"github.com/gin-gonic/gin"
	"myapp/middleware"
	"myapp/router/rAdmin"
	"myapp/router/rGeneral"
	"myapp/router/rMyApp"
	"net/http"
)

func InitRouter() (*gin.Engine, error)  {
	//服务配置
	router:= gin.Default()

	//中间件配置
	router.Use(middleware.SessionStore())

	//配置模板目录
	router.LoadHTMLGlob("templates/*")

	//重定向到登录页
	router.GET("/", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "./index")
	})

	//登录页
	router.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.tmpl", gin.H{
			"notice": "暂不开放注册",
		})
	})

	//进行功能区分
	//登录功能
	routerAdmin := router.Group("/admin")
	{
		rAdmin.AcceptRouter(routerAdmin)
	}

	//项目功能性接口
	routerMyApp := router.Group("/myApp", middleware.AuthRequired())
	{
		rMyApp.AcceptRouter(routerMyApp)
	}

	//通用功能接口
	routerGeneral := router.Group("/general", middleware.AuthRequired())
	{
		rGeneral.AcceptRouter(routerGeneral)
	}

	return router, nil
}
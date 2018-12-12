package router

import (
	"github.com/gin-gonic/gin"
	"myapp/router/rAdmin"
	"myapp/router/rGeneral"
	"myapp/router/rMyApp"
)

func initRouter() *gin.Engine  {
	//服务配置
	router:= gin.Default()
	//配置模板目录
	router.LoadHTMLGlob("templates/*")

	//进行功能区分
	//登录功能
	routerAdmin := router.Group("/admin")
	{
		rAdmin.AcceptRouter(routerAdmin)
	}

	//项目功能性接口
	routerMyApp := router.Group("/myApp")
	{
		rMyApp.AcceptRouter(routerMyApp)
	}

	//通用功能接口
	routerGeneral := router.Group("/general")
	{
		rGeneral.AcceptRouter(routerGeneral)
	}

	return router
}
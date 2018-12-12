package rGeneral

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AcceptRouter(rGroup *gin.RouterGroup) *gin.RouterGroup {
	rGroup.GET("/file/upload", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "success",
			"result":  "test",
		})
	})

	rGroup.GET("/file/download", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "success",
			"result":  "test",
		})
	})

	return rGroup
}
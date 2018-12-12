package rAdmin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AcceptRouter(rGroup *gin.RouterGroup) *gin.RouterGroup {
	rGroup.GET("/oauth2", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "success",
			"result":  "test",
		})
	})
	return rGroup
}
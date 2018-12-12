package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(context *gin.Context)  {
	context.JSON(http.StatusOK, gin.H{
		"code": 100,
		"msg" : "success",
		"result": gin.H{

		},
	})
}

func Register(context *gin.Context)  {
	context.JSON(http.StatusOK, gin.H{
		"code": 100,
	})
}
package rGeneral

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AcceptRouter(rGroup *gin.RouterGroup) *gin.RouterGroup {
	rGroup.POST("/file/upload", func(context *gin.Context) {
		file, _ := context.FormFile("file")

		err := context.SaveUploadedFile(file, "/data/" + file.Filename)

		if err != nil {
			log.Fatal(err)
			context.JSON(http.StatusOK, gin.H {
				"code": 99,
				"msg": "save file error",
				"result": gin.H {

				},
			})
			return
		}

		context.JSON(http.StatusOK, gin.H {
			"code": 100,
			"msg": "upload file success",
			"result": gin.H {

			},
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
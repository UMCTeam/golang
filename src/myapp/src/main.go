package main

import "github.com/gin-gonic/gin"

func main() {
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
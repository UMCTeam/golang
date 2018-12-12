package rAdmin

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"myapp/model"
	"net/http"
	"strings"
)

func AcceptRouter(rGroup *gin.RouterGroup) *gin.RouterGroup {
	rGroup.GET("/oauth2", func(context *gin.Context) {

	})

	rGroup.POST("/login", func(context *gin.Context) {
		session  := sessions.Default(context)
		username := context.PostForm("username")
		password := context.PostForm("password")

		if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
			context.JSON(http.StatusUnauthorized, gin.H{
				"error": "Parameters can't be empty",
			})
			return
		}

		err := model.User{}.FindUser(username, password)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authentication failed",
			})
		} else {
			session.Set("username", username)
			err := session.Save()

			if err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{
					"error": "Failed to generate session token",
				})
			} else {
				context.JSON(http.StatusOK, gin.H{
					"message": "Successfully authenticated user",
				})
			}
		}
	})

	return rGroup
}
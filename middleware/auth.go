package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthRequired() gin.HandlerFunc  {
	return func(context *gin.Context) {
		session := sessions.Default(context)
		user    := session.Get("user")

		if user != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid session token",
			})
		} else {
			context.Next()
		}
	}
}
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func RequireAdmin() gin.HandlerFunc {
	return func(context *gin.Context) {
		role, exists := context.Get("role")
		if !exists || role != "admin" {
			context.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "forbidden",})
			return
		}
	context.Next()
	}
}


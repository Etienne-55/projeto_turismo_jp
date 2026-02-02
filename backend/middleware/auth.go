package middleware

import (
	"net/http"
	"projeto_turismo_jp/utils"
	"strings"

	"github.com/gin-gonic/gin"
)


func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "not authorized1"})
		return
	}
	token = strings.TrimPrefix(token, "Bearer ")

	touristID, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	context.Set("touristID", touristID)
	context.Next()
}


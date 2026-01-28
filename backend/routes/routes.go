package routes

import (
	"projeto_turismo_jp/controllers"

	"github.com/gin-gonic/gin"
)


func AppRoutes(server *gin.Engine) {
	server.POST("/signup", controllers.Signup)
}


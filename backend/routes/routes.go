package routes

import (
	"projeto_turismo_jp/controllers"

	"github.com/gin-gonic/gin"
)


type Dependencies struct {
	TouristController *controllers.TouristController
	//add more when needed
}

func AppRoutes(server *gin.Engine, deps *Dependencies){
	//public routes
	server.POST("/signup", deps.TouristController.Signup)

	//test route
	server.GET("/test", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// protected := server.Group("/api")
	// {
	// 	//protected route for auth required functions
	// }

}


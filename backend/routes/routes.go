package routes

import (
	"projeto_turismo_jp/controllers"
	"projeto_turismo_jp/middleware"

	"github.com/gin-gonic/gin"
)


type Dependencies struct {
	TouristController *controllers.TouristController
	TripController *controllers.TripController
	//add more when needed
}

func AppRoutes(server *gin.Engine, deps *Dependencies){
	//public routes
	server.POST("/signup", deps.TouristController.Signup)
	server.POST("/login", deps.TouristController.Login)
	server.GET("/get_trip_by_id/:id", deps.TripController.GetTripByID)

	//test route
	server.GET("/test", func(c *gin.Context) {
		c.String(200, "pong")
	})

	//protected routes
	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("/trip", deps.TripController.CreateTrip)
	authenticated.DELETE("/delete_trip/:id", deps.TripController.DeleteTrip)


	server.GET("/get_all_trips", middleware.Authenticate, middleware.RequireAdmin(), deps.TripController.GetAllTrips)
}


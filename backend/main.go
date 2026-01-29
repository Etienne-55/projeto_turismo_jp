package main

import (
	"projeto_turismo_jp/controllers"
	"projeto_turismo_jp/db"
	"projeto_turismo_jp/repositories"
	"projeto_turismo_jp/routes"
	"projeto_turismo_jp/server"
)


func main() {
	db.InitDB()
	server := server.SetupServer()

	//Data layer
	touristRepo := repositories.NewTouristRepository(db.DB)

	//http layer
	touristController := controllers.NewTouristController(touristRepo)

	deps := &routes.Dependencies{
		TouristController: touristController,
	}
	routes.AppRoutes(server, deps)

	server.Run(":8080")
}


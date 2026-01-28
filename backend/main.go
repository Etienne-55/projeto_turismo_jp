package main

import (
	"projeto_turismo_jp/db"
	"projeto_turismo_jp/routes"

	"github.com/gin-gonic/gin"
)


func main() {
	db.InitDB()
	server := gin.Default()

	routes.AppRoutes(server)

	server.Run(":8080")
}


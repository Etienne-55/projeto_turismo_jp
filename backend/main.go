package main

import (
	"projeto_turismo_jp/db"
	"projeto_turismo_jp/routes"

	"github.com/gin-gonic/gin"
)

func setupServer() *gin.Engine {
	server := gin.Default()
	server.GET("/test_server", func(c *gin.Context){
		c.String(200, "server working on port 8080")
	})
	return server
}

func main() {
	db.InitDB()
	server := setupServer()

	routes.AppRoutes(server)

	server.Run(":8080")
}


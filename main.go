package main

import (
	"github.com/gin-gonic/gin"
	"sanyathecreator.com/eb-rest/db"
	"sanyathecreator.com/eb-rest/routes"
)

func main() {
	db.InitDB()

	// creates a router with Logger and Recovery middleware attached
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") // listen on localhost:8080
}

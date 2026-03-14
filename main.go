package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// setup http server
	server := gin.Default()

	// set a handler for GET request
	server.GET("/events", getEvents)

	server.Run(":8080") // localhost:8080
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello!"})
}

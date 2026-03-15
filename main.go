package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sanyathecreator.com/eb-rest/db"
	"sanyathecreator.com/eb-rest/models"
)

func main() {
	db.InitDB()

	// creates a router with Logger and Recovery middleware attached
	server := gin.Default()

	// Register route handlers
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080") // listen on localhost:8080
}

// getEvents returns all events currently stored in memory as a JSON array
func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

// createEvent parses a JSON request body into an Event and returns it.
// TODO: assign a real ID and UserID (e.g. from auth), and call event.Save()
func createEvent(context *gin.Context) {
	var event models.Event
	// ShouldBindJSON decodes the request body and enforces `binding:"required"` tags
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	// Placeholder values — real implementation should derive these from context/auth
	event.ID = 1
	event.UserID = 1

	event.Save()

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

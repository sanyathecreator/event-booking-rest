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
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later."})
		return
	}

	context.JSON(http.StatusOK, events)
}

// createEvent parses a JSON request body into an Event and returns it.
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

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save events. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

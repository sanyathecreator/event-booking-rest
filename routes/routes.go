package routes

import (
	"github.com/gin-gonic/gin"
	"sanyathecreator.com/eb-rest/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	// Register route handlers
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}

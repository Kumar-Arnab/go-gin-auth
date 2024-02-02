package routes

import "github.com/gin-gonic/gin"

func RoutePath(server *gin.Engine) {
	// points to the server initialized at main.go
	// setting up an handler for an incoming GET request,
	// server can provide support for GET, POST, PUT, PATCH, DELETE
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent)
}

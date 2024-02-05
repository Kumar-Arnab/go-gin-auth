package routes

import (
	"github.com/Kumar-Arnab/events-rests-auth/middleware"
	"github.com/gin-gonic/gin"
)

func RoutePath(server *gin.Engine) {
	// points to the server initialized at main.go
	// setting up an handler for an incoming GET request,
	// server can provide support for GET, POST, PUT, PATCH, DELETE
	server.GET("/events", GetEvents)
	server.GET("/events/:id", GetEvent)

	// creating a group of routes
	authenticatedGoup := server.Group("/")
	authenticatedGoup.Use(middleware.Authenticate)
	authenticatedGoup.POST("/events", CreateEvent)
	authenticatedGoup.PUT("/events/:id", UpdateEvent)
	authenticatedGoup.DELETE("/events/:id", DeleteEvent)
	authenticatedGoup.POST("/events/:id/register", RegisterForEvents)
	authenticatedGoup.DELETE("/events/:id/register", CancelRegistration)

	// users routes
	server.POST("/signup", SignUp)

	// login routes
	server.POST("/login", Login)
}

package routes

import "github.com/gin-gonic/gin"

func RoutePath(server *gin.Engine) {
	// points to the server initialized at main.go
	// setting up an handler for an incoming GET request,
	// server can provide support for GET, POST, PUT, PATCH, DELETE
	server.GET("/events", GetEvents)
	server.GET("/events/:id", GetEvent)
	server.POST("/events", CreateEvent)
	server.PUT("/events/:id", UpdateEvent)
	server.DELETE("/events/:id", DeleteEvent)

	// users routes
	server.POST("/signup", SignUp)
}

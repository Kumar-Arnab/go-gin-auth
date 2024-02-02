package main

import (
	"github.com/Kumar-Arnab/events-rests-auth/db"
	"github.com/Kumar-Arnab/events-rests-auth/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	// setting up a default http server behind the scenes and returns a pointer to the server
	server := gin.Default()

	routes.RoutePath(server)

	server.Run(":8081") // localhost:8081

}

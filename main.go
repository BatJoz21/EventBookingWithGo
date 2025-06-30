package main

import (
	"github.com/gin-gonic/gin"
	"practice.batjoz/event-booking-with-go/db"
	"practice.batjoz/event-booking-with-go/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}

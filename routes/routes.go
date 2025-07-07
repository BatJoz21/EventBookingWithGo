package routes

import (
	"github.com/gin-gonic/gin"
	"practice.batjoz/event-booking-with-go/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	// User Routes
	server.POST("/signup", signup)
	server.POST("/login", login)

	// Events Routes
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.GET("/events/registrations", seeAllRegistration)

	authenticatedRoute := server.Group("/")
	authenticatedRoute.Use(middlewares.Authenticate)
	authenticatedRoute.POST("/events", createEvent)
	authenticatedRoute.POST("/events/:id/register", newRegistration)
	authenticatedRoute.PUT("/events/:id", updateEvent)
	authenticatedRoute.DELETE("/events/:id", deleteEvent)
	authenticatedRoute.DELETE("/events/:id/register", cancelRegistration)
}

package routes

import (
	"rest-api/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	authenticated := server.Group("/")
	authenticated.Use(middleware.Autenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", updateEvent)
	authenticated.DELETE("/events/:id/register", updateEvent)
	// user routes
	server.POST("/signup", signup)
	server.POST("/login", login)

}

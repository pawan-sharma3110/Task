package main

import (
	"net/http"
	"rest/api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/events", getEvent)
	server.POST("/events", createEvent)
	server.Run(":8080")
}

func getEvent(context *gin.Context) {
	events := models.GetAllEvent()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse json data."})
		return
	}
	event.ID = 1
	event.UserID = 1
	context.JSON(http.StatusCreated, gin.H{"massage": "Event created.", "event": event})
}

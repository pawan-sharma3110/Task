package routes

import (
	"fmt"
	"net/http"
	"rest-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could fetch events4: %v", err)})
		return
	}
	context.JSON(http.StatusOK, events)
}
func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse event id."})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch event."})
		return
	}
	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Not authorized "})
		return
	}
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse json data."})
		return
	}
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not save event: %v", err)})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created.", "event": event})
}
func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse event id."})
		return
	}
	_, err = models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch event."})
		return
	}
	var updatedEvent models.Event
	err = context.ShouldBind(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse json data."})
		return
	}
	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not update event: %v", err)})
		return
	}
	context.JSON(http.StatusOK, gin.H{"massage": " update event."})
}
func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse event id."})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch event."})
		return
	}
	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not delete event: %v", err)})
		return
	}
	context.JSON(http.StatusOK, gin.H{"massage": "event deleteed."})
}

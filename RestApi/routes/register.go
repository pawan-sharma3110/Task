package routes

import (
	"net/http"
	"rest-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse event id."})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not fatch event."})
		return
	}
	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register user for this event."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"massag": "Registered"})
}

func cancelForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse event id."})
		return
	}
	var event models.Event
	event.ID = eventId
	err = event.CancleRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register user for this event."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"massag": "Canceled"})
}

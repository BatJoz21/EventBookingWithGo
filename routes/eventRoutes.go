package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"practice.batjoz/event-booking-with-go/models"
)

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the data"})
		return
	}

	user_id := context.GetInt64("user_id")
	event.UserID = user_id
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}

func updateEvent(context *gin.Context) {
	stringId := context.Param("id")
	intID, err := strconv.ParseInt(stringId, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse string to int64"})
		return
	}

	theEvent, err := models.GetEventByID(intID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch the event"})
		return
	}

	currentUserID := context.GetInt64("user_id")
	if theEvent.UserID != currentUserID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "user unauthorized"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the data"})
		return
	}
	updatedEvent.ID = intID

	err = updatedEvent.UpdateEventByID()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not update the event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully", "Updated Event\n": updatedEvent})
}

func getEvent(context *gin.Context) {
	stringId := context.Param("id")
	intID, err := strconv.ParseInt(stringId, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse string to int64"})
		return
	}

	var event *models.Event
	event, err = models.GetEventByID(intID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	context.JSON(http.StatusOK, event)
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
	}

	context.JSON(http.StatusOK, events)
}

func deleteEvent(context *gin.Context) {
	intID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse string to int64"})
		return
	}

	event, err := models.GetEventByID(intID)
	if err != nil {
		context.JSON(http.StatusNoContent, gin.H{"message": "could not find the event"})
	}

	currentUserID := context.GetInt64("user_id")
	if event.UserID != currentUserID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "user is unauthorized"})
		return
	}

	err = event.DeleteEventByID()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete the event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}

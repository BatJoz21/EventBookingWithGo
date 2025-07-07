package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"practice.batjoz/event-booking-with-go/models"
)

func newRegistration(context *gin.Context) {
	var newRegistration models.Registration

	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "failed to parse data"})
		return
	}

	_, err = models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "events doesn't exists"})
		return
	}

	newRegistration.EventID = eventID
	currentUserID := context.GetInt64("user_id")
	newRegistration.UserID = currentUserID

	err = newRegistration.SaveRegistration()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not register new registration"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "registration successfull"})
}

func cancelRegistration(context *gin.Context) {
	regisID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "failed to parse data"})
	}

	var regis *models.Registration
	regis, err = models.GetRegistrationByID(regisID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "failed to get registration data"})
	}

	err = regis.DeleteRegistration(context.GetInt64("user_id"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "failed to cancel registration data"})
	}

	context.JSON(http.StatusOK, gin.H{"message": "registration canceled"})
}

func seeAllRegistration(context *gin.Context) {
	registrations, err := models.GetAllRegistration()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "failed to obtain all registrations"})
		return
	}

	context.JSON(http.StatusOK, registrations)
}

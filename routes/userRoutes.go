package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"practice.batjoz/event-booking-with-go/models"
)

func signup(context *gin.Context) {
	var user models.Users
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the data"})
		return
	}

	err = user.SaveUser()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not save user to database"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "user saved successful"})
}

func login(context *gin.Context) {
	var user models.Users

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the data"})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"practice.batjoz/event-booking-with-go/utils"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
		return
	}

	user_id, err := utils.VerifiedToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
		return
	}

	context.Set("user_id", user_id)
	context.Next()
}
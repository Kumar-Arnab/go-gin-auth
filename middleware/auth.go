package middleware

import (
	"net/http"

	"github.com/Kumar-Arnab/events-rests-auth/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	// extracting the jwt token from the request header(Authorization)
	token := context.Request.Header.Get("Authorization")

	// if token is passed
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
		return
	}

	userId, err := utils.ValidateToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
		return
	}

	// set the user id for the next request handler in queue
	context.Set("userId", userId)
	// call the next request handler in queue
	context.Next()
}

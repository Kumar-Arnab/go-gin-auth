package routes

import (
	"net/http"

	"github.com/Kumar-Arnab/events-rests-auth/models"
	"github.com/gin-gonic/gin"
)

func SignUp(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data." + err.Error()})
		return
	}

	savedUser, err := user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user"})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User was created successfuly!", "user": savedUser})
}

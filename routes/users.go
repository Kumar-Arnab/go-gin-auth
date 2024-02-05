package routes

import (
	"net/http"

	"github.com/Kumar-Arnab/events-rests-auth/models"
	"github.com/Kumar-Arnab/events-rests-auth/utils"
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

func Login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data." + err.Error()})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login Successful!", "access_token": token})
}

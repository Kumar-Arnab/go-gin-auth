package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Kumar-Arnab/events-rests-auth/models"
	"github.com/gin-gonic/gin"
)

// we can get access to the incoming request and response functions using the context parameter
func getEvents(context *gin.Context) {
	// this data stored in the events will be automatically converted to JSON by the gin package
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	// while returning any response from getEvents function we need to use the JSON func to return in json format
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event

	// we pass a pointer of the variable that has to be populated with data
	// internally gin will take a look from the json body and store the body into a variable
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data." + err.Error()})
		return
	}

	event.UserID = 1

	savedEvent, er := event.Save()

	if er != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event", "event": er})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event was created successfuly!", "event": savedEvent})
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Could not parse event Id. %+v", err)})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Could not parse event Id. %+v", err)})
		return
	}
	context.JSON(http.StatusOK, event)
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Could not parse event Id. %+v", err)})
		return
	}

	_, err = models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Could not parse event Id. %+v", err)})
		return
	}

	var updatedEvent models.Event

	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data." + err.Error()})
		return
	}

	updatedEvent.ID = eventId
}

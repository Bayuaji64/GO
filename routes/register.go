package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {

	userId := context.GetInt64("userId")

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id", "error": err.Error()})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch  event ", "error": err.Error()})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not register  ", "error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "succes register"})

}

func cancelForEvent(context *gin.Context) {

	userId := context.GetInt64("userId")

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id", "error": err.Error()})
		return
	}

	var event models.Event
	event.ID = eventId

	err = event.CancelRegister(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not cancel", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Cancelled!"})

}

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	model "facade-golang/api-eventos/models"
	serv "facade-golang/api-eventos/services"
)

func GetEvents(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, model.Events)
}

func PostEvent(c *gin.Context) {
	var newEvent model.Event

	// Call BindJSON to bind the received JSON to
	// newEvent.
	if err := c.BindJSON(&newEvent); err != nil {
		return
	}

	// Add the new event to the slice.
	serv.InsertEventOnTopic(newEvent)
	c.IndentedJSON(http.StatusCreated, newEvent)
}

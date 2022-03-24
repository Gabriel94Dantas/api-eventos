package main

import (
	control "facade-golang/api-eventos/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/events", control.GetEvents)
	router.POST("/events", control.PostEvent)
	router.Run("localhost:7777")
}

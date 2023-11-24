package routes

import (
	"api/handlers"

	"github.com/gin-gonic/gin"
)

func NewRoutes() *gin.Engine {
	router := gin.Default()
	router.POST("/tasks", handlers.NewTaskHandler)
	return router
}

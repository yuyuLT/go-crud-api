package routes

import (
	"api/handlers"

	"github.com/gin-gonic/gin"
)

func NewRoutes() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("views/*.html")
	handlers.DbInit()

	router.GET("/", func(c *gin.Context) {
		tasks := handlers.GetTasks()
		c.HTML(200, "index.html", gin.H{"tasks": tasks})
	})

	router.POST("/new", handlers.TaskInsert)

	router.GET("/delete/:id", handlers.TaskDelete)
	return router
}

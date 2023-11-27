package handlers

import (
	"api/models"

	"github.com/gin-gonic/gin"
)

func TaskInsert(c *gin.Context) {
	body := c.PostForm("body")
	db := sqlConnect()
	defer db.Close()

	db.Create(&models.Task{Body: body})
	c.Redirect(302, "/")

}

package handlers

import (
	"api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func TaskDelete(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic("ERROR")
	}
	db := sqlConnect()
	var task models.Task
	db.First(&task, id)
	db.Delete(&task)
	db.Close()
	c.Redirect(302, "/")
}

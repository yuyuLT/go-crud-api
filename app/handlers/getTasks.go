package handlers

import (
	"api/models"
)

func GetTasks() []models.Task {
	db := sqlConnect()
	defer db.Close()

	var tasks []models.Task
	db.Order("created_at desc").Find(&tasks)
	return tasks
}

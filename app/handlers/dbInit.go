package handlers

import (
	"api/models"
	"fmt"
)

func DbInit() {
	db := sqlConnect()
	defer db.Close()
	// db.DropTable(&models.Task{})
	db.AutoMigrate(&models.Task{})
	fmt.Println("DB Init")
}

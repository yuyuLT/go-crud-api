package models

import (
	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	Body string `json:"body"`
}

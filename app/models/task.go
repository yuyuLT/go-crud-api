package task

import "time"

type Recipe struct {
	ID        string    `json:"id"`
	Body      string    `json:"body"`
	Tags      []string  `json:"tags"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdateAt  time.Time `json:"UpdateAt"`
}

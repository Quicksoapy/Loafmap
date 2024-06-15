package database

import "time"

type Marker struct {
	ID          uint      `json:"id"`
	Description string    `json:"description"`
	Datetime    time.Time `json:"datetime"`
	ImageUrl    string    `json:"image"`
	UserId      uint      `json:"userid"`
}

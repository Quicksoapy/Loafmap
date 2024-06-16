package database

import "time"

type Marker struct {
	ID          uint      `json:"id"`
	Description string    `json:"description"`
	Datetime    time.Time `json:"datetime"`
	ImageId     string    `json:"image"`
	UserId      uint      `json:"userid"`
	Latitude    string    `json:"latitude"`
	Longitude   string    `jsone:"longitude"`
}

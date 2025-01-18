package database

import (
	"time"
)

type Marker struct {
	ID                  uint      `json:"id"`
	Description         string    `json:"description"`
	Datetime            time.Time `json:"datetime"`
	Base64EncodedString string    `json:"image"`
	UserId              uint      `json:"userid"`
	Latitude            string    `json:"latitude"`
	Longitude           string    `json:"longitude"`
}

func (marker Marker) Add() error {
	_, err := database.Exec("INSERT INTO markers (userid, description, latitude, longitude, datetime, imageid)"+
		" VALUES ($1, $2, $3, $4, $5, $6); ", marker.UserId, marker.Description, marker.Latitude, marker.Longitude, time.Now(), marker.Base64EncodedString)
	return err
}

func (marker Marker) Delete() error {
	_, err := database.Exec("DELETE FROM markers WHERE id = $1", marker.ID)
	return err
}

func GetAllMarkers() (markers []Marker, err error) {
	rows, err := database.Query("SELECT * FROM markers")

	if err != nil {
		return
	}

	//TODO this could be faster
	for rows.Next() {
		var marker Marker
		err = rows.Scan(&marker.ID, &marker.UserId, &marker.Description, &marker.Latitude, &marker.Longitude, &marker.Datetime, &marker.Base64EncodedString)
		if err != nil {
			continue
		}

		markers = append(markers, marker)
	}

	return
}

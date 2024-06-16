package database

import (
	"database/sql"
	"fmt"
)

type Settings struct {
	Username string `json:"username"`
	Password string `json:"password"`
	DbName   string `json:"dbname"`
	Host     string `json:"host"`
	Port     uint   `json:"port"`
}

var database *sql.DB

func (settings Settings) DbCreateConnection() (err error) {
	connectionString := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", settings.Username, settings.Password, settings.Host, settings.Port, settings.DbName)

	database, err = sql.Open("postgres", connectionString)

	if err != nil {
		return
	}

	err = database.Ping()

	return
}

func (settings *Settings) Defaults() {
	if settings.Port == 0 {
		settings.Port = 5432
	}
	if settings.Username == "" {
		settings.Username = "username"
	}
	if settings.Password == "" {
		settings.Password = "password"
	}
	if settings.DbName == "" {
		settings.DbName = "loafmap_db"
	}
	if settings.Host == "" {
		settings.Host = "localhost"
	}
}

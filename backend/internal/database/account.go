package database

type Account struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
}

func (account Account) Add() {
	database.Query("")
}

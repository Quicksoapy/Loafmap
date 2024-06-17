package database

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
}

//TODO maybe do the hashing in frontend for additional security?

func (account Account) Add() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	_, err = database.Exec("INSERT INTO accounts (username, password, created_on, last_login) VALUES ($1, $2, $3, $4); ", account.Username, hashedPassword, time.Now(), time.Unix(0, 0))

	return err
}

func (account Account) Authenticate() bool {
	rows, err := database.Query("SELECT password FROM accounts WHERE username = $1", account.Username)
	if err != nil {
		print(err)
	}
	defer rows.Close()

	var password string

	for rows.Next() {
		err = rows.Scan(&password)
		if err != nil {
			print(err)
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(account.Password))

	return err == nil
}

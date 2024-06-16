package main

import (
	"loafmap/backend/internal/config"
	"loafmap/backend/internal/flags"
	"log"
)

func main() {
	options := flags.Parse()
	settings, err := config.Get(options.Config)
	_ = settings

	if err != nil {
		log.Fatal(err)
	}

	if err = settings.Database.DbCreateConnection(); err != nil {
		log.Fatal(err)
	}
}

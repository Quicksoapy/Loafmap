package main

import (
	"loafmap/backend/internal/api"
	"loafmap/backend/internal/config"
	"loafmap/backend/internal/flags"
	"log"

	_ "github.com/lib/pq"
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

	api.HandleRequests(settings.Api)
}

package config

import (
	"encoding/json"
	"loafmap/backend/internal/api"
	"loafmap/backend/internal/database"
	"os"
)

type Settings struct {
	Database database.Settings `json:"database"`
	Api      api.Settings      `json:"api"`
}

func (settings *Settings) defaults() {
	settings.Database.Defaults()
	settings.Api.Defaults()
}

func Get(configPath string) (config Settings, err error) {
	content, err := os.ReadFile(configPath)

	if err != nil {
		return
	}

	err = json.Unmarshal(content, &config)

	config.defaults()

	return
}

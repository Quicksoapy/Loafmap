package database

type Settings struct {
	Username string `json:"username"`
	Password string `json:"password"`
	DbName   string `json:"dbname"`
	Host     string `json:"host"`
	Port     uint   `json:"port"`
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

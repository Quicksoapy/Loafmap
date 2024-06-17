package api

import (
	"loafmap/backend/internal/api/accounts"
	"loafmap/backend/internal/api/markers"
	"net/http"
	"strconv"

	"github.com/gorilla/handlers"
	_ "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func HandleRequests(settings Settings) (err error) {
	router := mux.NewRouter().StrictSlash(true)

	router.Handle("/marker", markers.GetAll).Methods("GET", "OPTIONS")
	router.Handle("/marker/add", markers.Add).Methods("POST", "OPTIONS")
	router.Handle("/marker/delete", markers.Delete).Methods("POST", "OPTIONS")

	router.Handle("/account/login", accounts.Login).Methods("GET", "OPTIONS")
	router.Handle("/account/register", accounts.Register).Methods("GET", "OPTIONS")

	origins := handlers.AllowedOrigins([]string{"*"})
	headers := handlers.AllowedHeaders([]string{"Content-Type"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	err = http.ListenAndServe(":"+strconv.Itoa(int(settings.Port)), handlers.CORS(origins, headers, methods)(router))
	return
}

type Settings struct {
	Port uint `json:"port"`
}

func (settings *Settings) Defaults() {
	if settings.Port == 0 {
		settings.Port = 80
	}
}

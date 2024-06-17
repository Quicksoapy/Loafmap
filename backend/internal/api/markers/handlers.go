package markers

import (
	"encoding/json"
	"fmt"
	"io"
	"loafmap/backend/internal/database"
	"log"
	"net/http"
)

var GetAll = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// return "OKOK"
	json.NewEncoder(w).Encode("OKOK")

	markers, err := database.GetAllMarkers()
	if err != nil {
		log.Print(err)
		return
	}

	v, err := json.Marshal(markers)
	if err != nil {
		log.Print(err)
		return
	}

	_, err = w.Write(v)
	if err != nil {
		log.Print(err)
	}
})

var Add = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// return "OKOK"

	json.NewEncoder(w).Encode("OKOK")

	v := &database.Marker{}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := fmt.Fprint(w, err)
		if err != nil {
			return
		}
		return
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := fmt.Fprint(w, err)
		if err != nil {
			return
		}
		return
	}

	database.Marker.Add(*v)
})

var Delete = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// return "OKOK"

	json.NewEncoder(w).Encode("OKOK")

	v := &database.Marker{}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := fmt.Fprint(w, err)
		if err != nil {
			return
		}
		return
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := fmt.Fprint(w, err)
		if err != nil {
			return
		}
		return
	}

	database.Marker.Delete(*v)
})

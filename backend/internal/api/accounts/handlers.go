package accounts

import (
	"encoding/json"
	"fmt"
	"io"
	"loafmap/backend/internal/api/cookies"
	"loafmap/backend/internal/database"
	"net/http"
	"strconv"
)

var Register = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// return "OKOK"

	json.NewEncoder(w).Encode("OKOK")

	v := &database.Account{}

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

	database.Account.Add(*v)
})

var Login = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// return "OKOK"

	json.NewEncoder(w).Encode("OKOK")

	v := &database.Account{}

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

	if !database.Account.Authenticate(*v) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cookies.SetAccountCookie(w, strconv.FormatUint(uint64(v.ID), 10))
})

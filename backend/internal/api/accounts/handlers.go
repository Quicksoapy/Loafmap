package accounts

import (
	"encoding/json"
	"loafmap/backend/internal/api/cookies"
	"loafmap/backend/internal/database"
	"net/http"
)

var Register = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// return "OKOK"

	json.NewEncoder(w).Encode("OKOK")

	a := database.Account{Username: r.URL.Query().Get("username"), Password: r.URL.Query().Get("password")}

	database.Account.Add(a)
})

var Login = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	a := database.Account{Username: r.URL.Query().Get("username"), Password: r.URL.Query().Get("password")}

	if !database.Account.Authenticate(a) {
		json.NewEncoder(w).Encode("wrong login")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cookies.SetAccountCookie(w, a.Username)
})

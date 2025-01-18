package markers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"loafmap/backend/internal/database"
	"log"
	"net/http"
	"strconv"

	exifremove "github.com/scottleedavis/go-exif-remove"
)

var GetAll = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// return "OKOK"
	// json.NewEncoder(w).Encode("OKOK")

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

	userId, err := strconv.ParseUint(r.URL.Query().Get("userid"), 10, 32)
	if err != nil {
		print(err)
		return
	}

	img := r.URL.Query().Get("image")

	imgDecoded, e := base64.StdEncoding.DecodeString(img)
	if e != nil {
		fmt.Println(e)
	}

	imgStrippedExif, err := exifremove.Remove(imgDecoded)

	if err != nil {
		fmt.Println(e)
		return
	}

	imgEncoded := base64.StdEncoding.EncodeToString(imgStrippedExif)

	m := database.Marker{UserId: uint(userId), Description: r.URL.Query().Get("description"), Latitude: r.URL.Query().Get("latitude"),
		Longitude: r.URL.Query().Get("longitude"), Base64EncodedString: imgEncoded}

	database.Marker.Add(m)
})

var Delete = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// return "OKOK"

	json.NewEncoder(w).Encode("OKOK")

	id, err := strconv.ParseUint(r.URL.Query().Get("id"), 10, 32)
	if err != nil {
		print(err)
		return
	}

	m := database.Marker{ID: uint(id)}

	database.Marker.Delete(m)
})

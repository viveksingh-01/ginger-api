package handlers

import (
	"fmt"
	"net/http"
	"os"
)

func HandleRestaurants(w http.ResponseWriter, r *http.Request) {
	apiURL := os.Getenv("RESTAURANTS_API_URL")

	query := r.URL.Query()
	if query.Get("lat") != "" && query.Get("lng") != "" {
		apiURL += fmt.Sprintf("?lat=%s&lng=%s", query.Get("lat"), query.Get("lng"))
	}
}

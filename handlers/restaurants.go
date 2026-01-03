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

	// Create HTTP request
	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating request: %v", err), http.StatusInternalServerError)
		return
	}
	fmt.Println(req.URL)
}

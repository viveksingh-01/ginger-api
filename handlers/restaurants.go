package handlers

import (
	"fmt"
	"io"
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

	// Make request to Swiggy API
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching data: %v", err), http.StatusInternalServerError)
		return
	}

	// Check status-code received from Swiggy API and handle non 200 response
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		http.Error(w, fmt.Sprintf("Swiggy API returned status %d: %s", resp.StatusCode, string(bodyBytes)), resp.StatusCode)
		return
	}
}

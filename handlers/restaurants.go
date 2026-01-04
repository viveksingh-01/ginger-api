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

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading response: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Write response to client
	if _, err := w.Write(body); err != nil {
		http.Error(w, fmt.Sprintf("Error writing response: %v", err), http.StatusInternalServerError)
		return
	}
}

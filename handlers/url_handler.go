package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"url_shorterner_m/services"
)

type ShortenRequest struct {
	LongURL string `json:"long_url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

// POST /shorten
func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ShortenRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.LongURL == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	shortCode := services.CreateShortURL(req.LongURL)

	response := ShortenResponse{
		ShortURL: "http://localhost:8080/" + shortCode,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GET /{code}
func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	code := strings.TrimPrefix(r.URL.Path, "/")

	longURL, exists := services.GetLongURL(code)
	if !exists {
		http.NotFound(w, r)
		return
	}

	// THIS LINE is the key 🔥
	http.Redirect(w, r, longURL, http.StatusFound)
}

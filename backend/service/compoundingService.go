package main

import (
	"net/http"
	"os"
)

// CORS:
func setupResponse(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "text/html; charset=utf-8")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

// For http://localhost:8080
func healthCheckResponse(w http.ResponseWriter, req *http.Request) {
	_, err := w.Write([]byte("Nothing here, used for health check. Try /banks instead."))
	if err != nil {
		return
	}
}

func showBanks(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w)

	body, err := os.ReadFile("banks-response.json")
	if err != nil {
		// Handle error, perhaps by sending an HTTP error response
		http.Error(w, "Unable to read file", http.StatusInternalServerError)
		return
	}

	// Let the web server know it's JSON
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(body)
	if err != nil {
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(healthCheckResponse))
	mux.Handle("/banks", http.HandlerFunc(showBanks))
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return
	}
}

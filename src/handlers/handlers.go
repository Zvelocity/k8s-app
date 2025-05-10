package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

// HomeHandler returns a welcome message
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to Go Web App! This application is running in Kubernetes.\n"))
}

// HealthCheckHandler returns a 200 OK status for health checks
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := map[string]string{
		"status": "ok",
		"info":   "service is healthy",
	}
	json.NewEncoder(w).Encode(resp)
}

// TimeHandler returns the current server time
func TimeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := map[string]string{
		"time": time.Now().Format(time.RFC3339),
	}
	json.NewEncoder(w).Encode(resp)
}

// EchoHandler echoes back the request body
func EchoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := map[string]interface{}{
		"echoed": string(body),
	}
	json.NewEncoder(w).Encode(resp)
}

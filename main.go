package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
)

type LogEntry struct {
	Severity string `json:"severity"`
	Message  string `json:"message"`
	Time     string `json:"time"`
}

func logMessage(severity, message string) {
	entry := LogEntry{
		Severity: severity,
		Message:  message,
		Time:     time.Now().Format(time.RFC3339),
	}
	json.NewEncoder(os.Stdout).Encode(entry)
}

func logHandler(severity string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logMessage(severity, "This is a "+severity+" log")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Logged a " + severity + " message"))
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/error", logHandler("ERROR"))
	http.HandleFunc("/warn", logHandler("WARNING"))
	http.HandleFunc("/info", logHandler("INFO"))
	http.HandleFunc("/debug", logHandler("DEBUG"))

	logMessage("INFO", "Server starting on port "+port)
	http.ListenAndServe(":"+port, nil)
}

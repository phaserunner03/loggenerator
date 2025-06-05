package main

import (
	"encoding/json"
	"fmt"
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
	_ = json.NewEncoder(os.Stdout).Encode(entry)
}

func logHandler(severity string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logMessage(severity, "This is a "+severity+" log")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Logged a %s message\n", severity)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	logMessage("INFO", "Starting server on port "+port)

	http.HandleFunc("/info", logHandler("INFO"))
	http.HandleFunc("/warn", logHandler("WARNING"))
	http.HandleFunc("/error", logHandler("ERROR"))
	http.HandleFunc("/debug", logHandler("DEBUG"))

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		logMessage("ERROR", "Server failed: "+err.Error())
	}
}

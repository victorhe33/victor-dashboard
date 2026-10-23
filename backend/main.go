package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type widget struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Value string `json:"value"`
}

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/health", handleHealth)
	mux.HandleFunc("GET /api/widgets", handleWidgets)
	return mux
}

func handleHealth(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, map[string]string{
		"status": "ok",
		"time":   time.Now().UTC().Format(time.RFC3339),
	})
}

// Placeholder data until real sources are wired in.
func handleWidgets(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, []widget{
		{ID: "tickets", Title: "Open tickets", Value: "—"},
		{ID: "prs", Title: "PRs awaiting review", Value: "—"},
		{ID: "deploys", Title: "Recent deploys", Value: "—"},
	})
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("encode response: %v", err)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4581"
	}
	addr := ":" + port
	log.Printf("listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, newMux()))
}

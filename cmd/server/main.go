package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dukerupert/a-team-asphalt/internal/handlers"
	"github.com/dukerupert/a-team-asphalt/internal/templates"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	tmpl, err := templates.Load("templates")
	if err != nil {
		log.Fatalf("Failed to load templates: %v", err)
	}

	h := handlers.New(tmpl)

	mux := http.NewServeMux()

	// Static files
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Pages — all concepts are single-page designs for now
	mux.HandleFunc("GET /{$}", h.Home)
	mux.HandleFunc("GET /about", h.About)
	mux.HandleFunc("GET /services", h.Services)
	mux.HandleFunc("GET /contact", h.Contact)

	// Concept switcher
	mux.HandleFunc("POST /set-concept", h.SetConcept)

	fmt.Printf("Server starting on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

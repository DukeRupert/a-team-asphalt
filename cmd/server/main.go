package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dukerupert/a-team-asphalt/internal/handlers"
	"github.com/dukerupert/a-team-asphalt/internal/mailer"
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

	// Postmark mailer — gracefully nil if not configured
	m := mailer.New(mailer.Config{
		Token: os.Getenv("POSTMARK_TOKEN"),
		From:  os.Getenv("POSTMARK_FROM"),
		To:    os.Getenv("ESTIMATE_TO"),
	})
	if m == nil {
		log.Println("WARN: POSTMARK_TOKEN not set — estimate emails disabled")
	}

	h := handlers.New(tmpl, m)

	mux := http.NewServeMux()

	// Static files
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// robots.txt at root
	mux.HandleFunc("GET /robots.txt", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/robots.txt")
	})

	// Form handlers
	mux.HandleFunc("POST /estimate", h.Estimate)

	// Pages — industrial is the primary landing page
	mux.HandleFunc("GET /{$}", h.Home)
	mux.HandleFunc("GET /about", h.About)
	mux.HandleFunc("GET /services", h.Services)
	mux.HandleFunc("GET /contact", h.Contact)
	mux.HandleFunc("GET /placard", h.Placard)
	mux.HandleFunc("GET /placard/services", h.PlacardServices)
	mux.HandleFunc("GET /placard/about", h.PlacardAbout)
	mux.HandleFunc("GET /placard/contact", h.PlacardContact)

	fmt.Printf("Server starting on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

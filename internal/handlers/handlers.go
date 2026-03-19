package handlers

import (
	"log"
	"net/http"
	"strings"

	"github.com/dukerupert/a-team-asphalt/internal/mailer"
	"github.com/dukerupert/a-team-asphalt/internal/templates"
)

// Handlers holds dependencies for HTTP handlers.
type Handlers struct {
	tmpl   *templates.Templates
	mailer *mailer.Mailer // nil if not configured
}

// New creates a new Handlers instance.
func New(tmpl *templates.Templates, m *mailer.Mailer) *Handlers {
	return &Handlers{tmpl: tmpl, mailer: m}
}

func (h *Handlers) render(w http.ResponseWriter, r *http.Request, page string) {
	data := templates.PageData{
		Concept:     "placard",
		CurrentPage: r.URL.Path,
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.tmpl.Render(w, "placard", page, data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// Home renders the home page.
func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	h.render(w, r, "home")
}

// About renders the about page.
func (h *Handlers) About(w http.ResponseWriter, r *http.Request) {
	h.render(w, r, "home")
}

// Services renders the services page.
func (h *Handlers) Services(w http.ResponseWriter, r *http.Request) {
	h.render(w, r, "home")
}

// Contact renders the contact page.
func (h *Handlers) Contact(w http.ResponseWriter, r *http.Request) {
	h.render(w, r, "home")
}

// Estimate handles the estimate request form submission.
func (h *Handlers) Estimate(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	name := strings.TrimSpace(r.FormValue("name"))
	phone := strings.TrimSpace(r.FormValue("phone"))
	email := strings.TrimSpace(r.FormValue("email"))
	projectType := strings.TrimSpace(r.FormValue("project_type"))
	description := strings.TrimSpace(r.FormValue("description"))

	if name == "" || phone == "" || projectType == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	log.Printf("Estimate request: name=%q phone=%q email=%q type=%q",
		name, phone, email, projectType)

	if h.mailer != nil {
		err := h.mailer.SendEstimateNotification(mailer.EstimateRequest{
			Name:        name,
			Phone:       phone,
			Email:       email,
			ProjectType: projectType,
			Description: description,
		})
		if err != nil {
			log.Printf("ERROR sending estimate email: %v", err)
			// Still redirect — don't lose the lead over an email failure
		}
	}

	http.Redirect(w, r, "/?submitted=1#contact", http.StatusSeeOther)
}

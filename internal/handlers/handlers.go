package handlers

import (
	"log"
	"net/http"
	"strings"

	"github.com/dukerupert/a-team-asphalt/internal/mailer"
	"github.com/dukerupert/a-team-asphalt/internal/services"
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

func (h *Handlers) renderConcept(w http.ResponseWriter, r *http.Request, concept, page string) {
	data := templates.PageData{
		Concept:     concept,
		CurrentPage: r.URL.Path,
		Params:      map[string]string{},
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.tmpl.Render(w, concept, page, data); err != nil {
		log.Printf("ERROR rendering %s/%s: %v", concept, page, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *Handlers) render(w http.ResponseWriter, r *http.Request, page string) {
	h.renderConcept(w, r, "industrial", page)
}

// Home renders the home page.
func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	h.render(w, r, "home")
}

// About renders the about page.
func (h *Handlers) About(w http.ResponseWriter, r *http.Request) {
	h.render(w, r, "about")
}

// Services renders the services page.
func (h *Handlers) Services(w http.ResponseWriter, r *http.Request) {
	h.render(w, r, "services")
}

// Contact renders the contact page with optional service query parameter.
func (h *Handlers) Contact(w http.ResponseWriter, r *http.Request) {
	data := templates.PageData{
		Concept:     "industrial",
		CurrentPage: r.URL.Path,
		Params:      map[string]string{},
	}
	if svc := r.URL.Query().Get("service"); svc != "" {
		data.Params["service"] = svc
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.tmpl.Render(w, "industrial", "contact", data); err != nil {
		log.Printf("ERROR rendering industrial/contact: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// ServiceDetail renders an individual service page by slug.
func (h *Handlers) ServiceDetail(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	svc := services.BySlug(slug)
	if svc == nil {
		http.NotFound(w, r)
		return
	}
	data := templates.PageData{
		Concept:     "industrial",
		CurrentPage: "/services",
		Params:      map[string]string{},
		Service:     svc,
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.tmpl.Render(w, "industrial", "service-detail", data); err != nil {
		log.Printf("ERROR rendering service-detail for %s: %v", slug, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// Placard renders the placard concept home page.
func (h *Handlers) Placard(w http.ResponseWriter, r *http.Request) {
	h.renderConcept(w, r, "placard", "home")
}

// PlacardServices renders the placard services page.
func (h *Handlers) PlacardServices(w http.ResponseWriter, r *http.Request) {
	h.renderConcept(w, r, "placard", "services")
}

// PlacardAbout renders the placard about page.
func (h *Handlers) PlacardAbout(w http.ResponseWriter, r *http.Request) {
	h.renderConcept(w, r, "placard", "about")
}

// PlacardContact renders the placard contact page.
func (h *Handlers) PlacardContact(w http.ResponseWriter, r *http.Request) {
	h.renderConcept(w, r, "placard", "contact")
}

// validRedirects lists allowed redirect targets for form submissions.
var validRedirects = map[string]bool{
	"/":               true,
	"/contact":         true,
	"/placard":         true,
	"/placard/contact": true,
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

	if name == "" || (phone == "" && email == "") {
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

	// Redirect back to the originating page
	redirect := "/"
	if dest := r.FormValue("redirect"); validRedirects[dest] {
		redirect = dest
	}
	http.Redirect(w, r, redirect+"?submitted=1", http.StatusSeeOther)
}

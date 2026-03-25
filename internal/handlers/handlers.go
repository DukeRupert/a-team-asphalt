package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dukerupert/a-team-asphalt/internal/mailer"
	"github.com/dukerupert/a-team-asphalt/internal/services"
	"github.com/dukerupert/a-team-asphalt/internal/templates"
)

const baseURL = "https://ateamasphalt.com"

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
		Concept:       concept,
		CurrentPage:   r.URL.Path,
		CanonicalPath: r.URL.Path,
		BaseURL:       baseURL,
		Params:        map[string]string{},
		Year:          time.Now().Year(),
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
		Concept:       "industrial",
		CurrentPage:   r.URL.Path,
		CanonicalPath: "/contact",
		BaseURL:       baseURL,
		Params:        map[string]string{},
		Year:          time.Now().Year(),
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
		h.NotFound(w, r)
		return
	}
	data := templates.PageData{
		Concept:       "industrial",
		CurrentPage:   "/services",
		CanonicalPath: r.URL.Path,
		BaseURL:       baseURL,
		Params:        map[string]string{},
		Service:       svc,
		Year:          time.Now().Year(),
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.tmpl.Render(w, "industrial", "service-detail", data); err != nil {
		log.Printf("ERROR rendering service-detail for %s: %v", slug, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// NotFound renders a branded 404 page.
func (h *Handlers) NotFound(w http.ResponseWriter, r *http.Request) {
	data := templates.PageData{
		Concept:       "industrial",
		CurrentPage:   "",
		CanonicalPath: r.URL.Path,
		BaseURL:       baseURL,
		Params:        map[string]string{},
		Year:          time.Now().Year(),
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	if err := h.tmpl.Render(w, "industrial", "404", data); err != nil {
		log.Printf("ERROR rendering 404: %v", err)
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}

// Privacy renders the privacy policy page.
func (h *Handlers) Privacy(w http.ResponseWriter, r *http.Request) {
	h.render(w, r, "privacy")
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

// Sitemap generates a sitemap.xml with all public pages.
func (h *Handlers) Sitemap(w http.ResponseWriter, r *http.Request) {
	now := time.Now().Format("2006-01-02")
	w.Header().Set("Content-Type", "application/xml; charset=utf-8")

	fmt.Fprint(w, `<?xml version="1.0" encoding="UTF-8"?>`)
	fmt.Fprint(w, `<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">`)

	// Static pages with priorities
	static := []struct {
		path     string
		priority string
	}{
		{"/", "1.0"},
		{"/services", "0.9"},
		{"/about", "0.7"},
		{"/contact", "0.8"},
		{"/privacy", "0.3"},
	}
	for _, p := range static {
		fmt.Fprintf(w, `<url><loc>%s%s</loc><lastmod>%s</lastmod><priority>%s</priority></url>`,
			baseURL, p.path, now, p.priority)
	}

	// Service detail pages
	for _, svc := range services.All() {
		fmt.Fprintf(w, `<url><loc>%s/services/%s</loc><lastmod>%s</lastmod><priority>0.8</priority></url>`,
			baseURL, svc.Slug, now)
	}

	fmt.Fprint(w, `</urlset>`)
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
	// Limit request body to 64KB to prevent abuse.
	r.Body = http.MaxBytesReader(w, r.Body, 64<<10)
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Honeypot field — if filled, it's a bot.
	if r.FormValue("company") != "" {
		redirect := "/"
		if dest := r.FormValue("redirect"); validRedirects[dest] {
			redirect = dest
		}
		http.Redirect(w, r, redirect+"?submitted=1", http.StatusSeeOther)
		return
	}

	name := truncate(strings.TrimSpace(r.FormValue("name")), 100)
	phone := truncate(strings.TrimSpace(r.FormValue("phone")), 20)
	email := truncate(strings.TrimSpace(r.FormValue("email")), 254)
	projectType := truncate(strings.TrimSpace(r.FormValue("project_type")), 100)
	description := truncate(strings.TrimSpace(r.FormValue("description")), 2000)

	if name == "" || (phone == "" && email == "") {
		// Redirect back with error flag so the form can show an inline message.
		redirect := "/"
		if dest := r.FormValue("redirect"); validRedirects[dest] {
			redirect = dest
		}
		http.Redirect(w, r, redirect+"?error=missing", http.StatusSeeOther)
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

// truncate returns s cut to at most max bytes on a rune boundary.
func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	// Walk back to avoid splitting a multi-byte rune.
	for max > 0 && max < len(s) && s[max]>>6 == 2 {
		max--
	}
	return s[:max]
}

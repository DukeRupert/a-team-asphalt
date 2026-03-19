package handlers

import (
	"net/http"

	"github.com/dukerupert/a-team-asphalt/internal/templates"
)

// Handlers holds dependencies for HTTP handlers.
type Handlers struct {
	tmpl *templates.Templates
}

// New creates a new Handlers instance.
func New(tmpl *templates.Templates) *Handlers {
	return &Handlers{tmpl: tmpl}
}

const conceptCookie = "concept"
const defaultConcept = "dark-nostalgia"

func getConcept(r *http.Request) string {
	c, err := r.Cookie(conceptCookie)
	if err != nil || c.Value == "" {
		return defaultConcept
	}
	for _, valid := range templates.Concepts {
		if c.Value == valid.Slug {
			return c.Value
		}
	}
	return defaultConcept
}

func (h *Handlers) render(w http.ResponseWriter, r *http.Request, page string) {
	concept := getConcept(r)
	data := templates.PageData{
		Concept:     concept,
		CurrentPage: r.URL.Path,
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.tmpl.Render(w, concept, page, data); err != nil {
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

// SetConcept sets the concept cookie and redirects back.
func (h *Handlers) SetConcept(w http.ResponseWriter, r *http.Request) {
	concept := r.FormValue("concept")
	valid := false
	for _, c := range templates.Concepts {
		if concept == c.Slug {
			valid = true
			break
		}
	}
	if !valid {
		concept = defaultConcept
	}
	http.SetCookie(w, &http.Cookie{
		Name:  conceptCookie,
		Value: concept,
		Path:  "/",
	})
	ref := r.Header.Get("Referer")
	if ref == "" {
		ref = "/"
	}
	http.Redirect(w, r, ref, http.StatusSeeOther)
}

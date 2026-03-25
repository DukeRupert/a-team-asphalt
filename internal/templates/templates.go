package templates

import (
	"fmt"
	"html/template"
	"io"
	"path/filepath"

	"github.com/dukerupert/a-team-asphalt/internal/services"
)

// Templates holds parsed template sets, one per concept.
type Templates struct {
	concepts map[string]*template.Template
}

// PageData is the data passed to every page template.
type PageData struct {
	Concept       string            // "placard"
	CurrentPage   string
	CanonicalPath string              // URL path for canonical/OG (no query params)
	BaseURL       string              // e.g. "https://ateamasphalt.com"
	Params        map[string]string   // query parameters
	Service       *services.Service   // populated on service detail pages
	Year          int                 // current year for copyright
}

// Concepts lists valid concept names in display order.
var Concepts = []struct {
	Slug  string
	Label string
}{
	{"industrial", "Industrial"},
	{"placard", "Placard"},
}

// Load parses each concept directory into a template set.
func Load(dir string) (*Templates, error) {
	concepts := make(map[string]*template.Template)

	for _, c := range Concepts {
		pattern := filepath.Join(dir, c.Slug, "*.html")
		files, err := filepath.Glob(pattern)
		if err != nil {
			return nil, fmt.Errorf("glob %s: %w", pattern, err)
		}
		if len(files) == 0 {
			return nil, fmt.Errorf("no templates found for concept %q at %s", c.Slug, pattern)
		}

		t, err := template.ParseFiles(files...)
		if err != nil {
			return nil, fmt.Errorf("parse concept %q: %w", c.Slug, err)
		}
		concepts[c.Slug] = t
	}

	return &Templates{concepts: concepts}, nil
}

// Render executes a concept's page template into the writer.
func (t *Templates) Render(w io.Writer, concept, page string, data PageData) error {
	tmpl, ok := t.concepts[concept]
	if !ok {
		return fmt.Errorf("concept %q not found", concept)
	}
	return tmpl.ExecuteTemplate(w, page, data)
}

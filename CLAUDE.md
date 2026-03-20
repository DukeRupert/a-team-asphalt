# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Brochure site for A-Team Asphalt LLC, a family-owned asphalt contractor in Sumner, WA (30+ years, co-owned by Gary Jorgensen and Don Lee). Go server-side rendered with `html/template`, no frontend framework, no external Go dependencies.

## Commands

```bash
# Run locally (default port 8080, override with PORT env)
go run ./cmd/server/

# Build binary
go build -o server ./cmd/server

# Docker
docker build -t a-team-asphalt .
```

No test suite or linter is configured.

## Architecture

**Entry point**: `cmd/server/main.go` — sets up routes, loads templates, initializes Postmark mailer.

**Internal packages**:
- `internal/handlers/` — HTTP handlers. All page handlers call `Templates.Render()` with a concept name and page name. The `render()` helper hardcodes "industrial" as the active concept.
- `internal/templates/` — Loads all `*.html` files per concept directory into a `template.Template` set. `PageData` struct is passed to every template (Concept, CurrentPage, Params, Service).
- `internal/services/` — Hardcoded catalog of 16 asphalt/concrete services as Go structs. Each service has slug, name, tagline, category, full SEO copy, process steps, related links. Look up by slug via `services.BySlug()`.
- `internal/mailer/` — Postmark email integration for estimate form submissions. Gracefully nil if `POSTMARK_TOKEN` not set.

**Template system**: Templates are organized by "concept" (theme variant). Each concept is a directory under `templates/` containing `*.html` files. The active concept is **industrial**. Templates use Go's `{{ define "name" }}` / `{{ template "name" . }}` pattern with `_partials.html` providing shared nav, footer, head styles, and scripts.

**Routes**:
- `GET /` — Home
- `GET /about` — About
- `GET /services` — Services listing with jump nav
- `GET /services/{slug}` — Service detail (16 services)
- `GET /contact` — Contact form (accepts `?service=` query param to pre-fill project type)
- `POST /estimate` — Form handler, sends email via Postmark, redirects with `?submitted=1`
- `GET /static/` — Static file server

**The `/placard` routes and `templates/placard/` directory are paused — do not modify or extend.**

## Styling

Two complete CSS files handle dark/light mode, swapped via `localStorage` + `prefers-color-scheme`:
- **Dark**: `static/css/industrial.css` (~2700 lines)
- **Light**: `static/css/cool-steel.css` (~2900 lines)

Both files must be kept in sync structurally — a new component needs styles in both.

Hero and footer sections remain dark in both themes.

## Design System

Full reference in `.impeccable.md`. Key constraints:

- **Typography**: Bebas Neue (display), Oswald UPPERCASE (nav/labels/buttons), Source Sans 3 (body)
- **Red `#C8102E`**: Reserved for CTAs and accent only — never on large surface areas
- **Clipped-corner buttons**: `clip-path: polygon(...)` — no `border-radius` anywhere in the design
- **Swooping dividers**: Red-gradient bar with right-pointing arrow between content zones
- **Section eyebrows**: Oswald 11px UPPERCASE, wide tracking, red text, with animated line rule
- **Chrome/silver accents**: Cool metallic tones, NOT warm gold (gold belongs to paused placard)
- **SVG noise texture**: `feTurbulence` overlay on dark sections

## Brand Voice

Full guide in `docs/brand-guide.md`. Write like a craftsman, not a marketer.

- Short sentences. Plain words. Say the true thing.
- **Banned words**: solutions, leverage, utilize, seamless, passionate, journey, world-class, partner/partnership, amazing, incredible, exciting
- No question-mark headlines, no exclamation points
- Lead with what the customer gets, not what we do
- The A-Team reference is an occasional easter egg, never a gimmick

## Deployment

GitHub Actions workflow (`.github/workflows/deploy-dev.yml`) builds a Docker image, pushes to Docker Hub (`dukerupert/a-team-asphalt:dev-{sha}`), and deploys to a VPS via SSH.

Environment variables (see `.env.example`):
- `PORT` — Listen port (default 8080)
- `POSTMARK_TOKEN`, `POSTMARK_FROM`, `ESTIMATE_TO` — Email config

## Adding a New Page

1. Create `templates/industrial/newpage.html` using `{{ define "newpage" }}` block
2. Include shared partials via `{{ template "_partials" }}` pattern
3. Add handler method in `internal/handlers/handlers.go` calling `h.render(w, r, "newpage")`
4. Register route in `cmd/server/main.go`
5. Add styles to both `industrial.css` and `cool-steel.css`

## Adding/Editing a Service

All 16 services are defined in `internal/services/services.go`. Each `Service` struct contains the full content for its SEO landing page (tagline, pitch, includes, why-us points, process steps, service areas, good-to-know callouts, related service links). The service detail template reads from this struct — no separate content files.

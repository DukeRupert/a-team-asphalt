# A-Team Asphalt Co. — Web UI Reference Guide

*Derived from analysis of the existing business card and brand research.*
*Owner: Gary A. Jorgensen Jr. · Sumner, WA 98390 · Lic. #TEAMAAL855JE*

---

## Brand Snapshot

| | |
|---|---|
| **Company** | A-Team Asphalt Co. |
| **Owner** | Gary A. Jorgensen Jr. |
| **Location** | Sumner, WA (Pierce County — Tacoma metro) |
| **Phone** | (253) 365-1288 |
| **License** | TEAMAAL855JE |
| **Tagline** | "As good as the best, different than the rest" |
| **Credential** | Over 20 years of paving excellence |
| **Market** | Western Washington — Puget Sound, Pierce & King counties |

---

## What the Business Card Is Telling Us

Gary's card is not a generic contractor card. It's an artifact with a clear point of view. Reading it closely:

**The format is a placard, not a flyer.** The double-rule border, labeled zones, rank-style "owner" badge above his name — this is the visual language of a military service record or a professional credential plate. It signals *authority* and *legitimacy* without saying either word.

**The typography is Western slab serif.** The A-Team Asphalt Co. name uses a high-contrast slab serif with a distressed or shadow treatment — old west poster, not military stencil. This is warm, bold, and slightly irreverent. It says "we've been around" more than "we're a corporation."

**The color system is warmer than we assumed.** The "silver" in our earlier concepts was cool chrome. The card's metallic is closer to champagne gold — warm, aged, almost brass-like. This is a crucial distinction. Cool chrome reads modern and industrial. Warm gold reads established and earned.

**The tagline is unpolished and genuine.** "As good as the best, different then the rest" has a typo (then/than) that Gary clearly didn't care enough to fix — or didn't notice. This is not a liability. It's evidence that the brand voice is authentically his, not marketing-polished. The web copy should have that same quality of *a real person talking*, not a brand consultant's version of a real person.

**"Over 20 years" is the lead credential.** Not certifications, not awards — time in the trade. That's what Gary is proudest of and what customers respond to most in this category.

---

## Color System

The card corrects our earlier assumptions. This is the real palette:

### Primary Colors

```
Near-Black Base
  Hex: #141414
  Use: Page backgrounds, nav, hero sections, footer
  Card equivalent: The deep black ground the entire card sits on

Champagne Gold (the real "silver")
  Hex: #C8B882
  Light variant: #DDD0A0
  Dark variant: #9A8C60
  Use: Primary display type, borders, decorative rules, owner name treatments
  Card equivalent: The metallic foil used for the border system and name

A-Team Red
  Hex: #C0200A
  Bright variant: #D93010
  Muted variant: #8B1508
  Use: Company name, CTAs, accent rules, hover states
  Card equivalent: The slab serif company name treatment

```

### Secondary / Support Colors

```
Dark Band
  Hex: #1E1C18
  Use: Nav band behind links, labeled section headers ("owner"-style badges)
  Card equivalent: The dark horizontal band behind "owner" label

Off-White / Warm Cream
  Hex: #F2EFE6
  Use: Body text on dark backgrounds, fine print, contact details
  Card equivalent: The lighter contact info text

Warm Mid-Gray
  Hex: #6A6560
  Use: Muted labels, secondary text, divider rules
  Card equivalent: The dotted separators and secondary info zones
```

### Color Roles

| Element | Color | Notes |
|---|---|---|
| Page background | `#141414` | Near-black, not pure black |
| Nav background | `#1E1C18` | Slightly warmer/lighter than base |
| Hero headline | `#C0200A` | Company name always in red |
| Display subheads | `#C8B882` | Champagne gold — Gary's name treatment |
| Body text (on dark) | `#F2EFE6` | Warm off-white, not pure white |
| Muted text / labels | `#6A6560` | Section labels, fine print |
| Primary CTA | `#C0200A` bg, `#F2EFE6` text | Red button |
| Secondary CTA | transparent, `#C8B882` border + text | Ghost button in gold |
| Borders / rules | `#C8B882` at 40–70% opacity | Double-rule card border system |
| Accent separators | `#C0200A` thin line, 1–2px | Red accent rules |

---

## Typography

### The Card's Type Approach

The card uses two distinct register sizes and weights:
1. **Display / Identity** — large, slab serif, high contrast, slightly distressed feel. Company name and owner name.
2. **Information / Detail** — smaller, upright, condensed, functional. Contact info, labels, credentials.

### Web Font Recommendations

**Display (company name, hero headlines, section titles):**
```
Playfair Display — bold/black weight
Clarendon (via web) — if available
Zilla Slab — 700/800 weight
Bitter — 700/900 weight

Recommended pairing: Zilla Slab Black or Playfair Display Black
Reasoning: Matches the slab serif character of the card name treatment.
High x-height, strong stroke contrast, readable at large sizes.
```

**Secondary Display (owner name, subheadings):**
```
Oswald — 600/700 weight, uppercase, tracked wide
Barlow Condensed — 700/800 weight
```

**Body / Functional:**
```
Barlow — 400/500 weight
IBM Plex Sans — 400/500 weight
Source Sans 3 — 400/600 weight
```

**Recommended Google Fonts import:**
```css
@import url('https://fonts.googleapis.com/css2?family=Zilla+Slab:wght@700&family=Oswald:wght@500;600;700&family=Barlow:wght@400;500;600&family=Barlow+Condensed:wght@600;700&display=swap');
```

### Type Scale

```css
/* Identity / Hero */
--type-hero:        clamp(56px, 8vw, 100px);   /* Company name, hero H1 */
--type-display:     clamp(36px, 5vw, 60px);    /* Section headlines */
--type-title:       clamp(22px, 3vw, 32px);    /* Card titles, subheads */

/* Functional */
--type-label:       10px;                       /* "owner"-style badge labels */
--type-nav:         11px;                       /* Nav links */
--type-body:        15px;                       /* Body copy */
--type-small:       12px;                       /* Fine print, credentials */

/* Tracking */
--track-label:      0.18em;   /* Uppercase label tracking */
--track-nav:        0.12em;   /* Nav link tracking */
--track-title:      0.04em;   /* Title tracking */
--track-hero:       -0.02em;  /* Hero headline — tight */
```

### Type Treatments from the Card

**The "owner" badge:**
A small all-caps label in a dark band, tight letter-spacing, appearing above a person's name or section title. Translates directly to website section eyebrows and role labels.

```css
.eyebrow {
  font-family: 'Oswald', sans-serif;
  font-size: 10px;
  font-weight: 600;
  letter-spacing: 0.2em;
  text-transform: uppercase;
  color: #6A6560;
  background: #1E1C18;
  padding: 0.2rem 0.6rem;
  display: inline-block;
}
```

**The double-rule border:**
Two thin parallel rules framing a section — the card's most distinctive structural element. On the web, use as a section divider or around callout blocks.

```css
.double-rule {
  border-top: 1px solid rgba(200, 184, 130, 0.5);
  box-shadow: 0 2px 0 0 rgba(200, 184, 130, 0.2);
  /* Creates the visual of two rules with the inner one slightly dimmer */
}
```

**The dotted separator:**
Between contact info zones on the card — a row of dots rather than a solid line. On the web, use between list items or as a visual pause within a content block.

```css
.dot-rule {
  border-bottom: 2px dotted rgba(200, 184, 130, 0.3);
  padding-bottom: 0.75rem;
  margin-bottom: 0.75rem;
}
```

---

## Layout Principles

### The Card's Structure, Translated

The card is organized as a **placard** — a formal display plate with labeled zones. This translates to web as:

**Zone-based sections** rather than free-flowing content. Each section of the page has a clear container with labeled entry points (eyebrow labels), defined content, and ruled exits. No content bleeds ambiguously between sections.

**Vertical hierarchy within a horizontal space.** The card runs the company name vertically while contact info runs horizontally — a tension between the identity (tall, bold, occupying its own axis) and the information (organized, readable, practical). On the web this maps to: hero section that asserts identity with large type, then structured content sections below that deliver practical information.

**Labels before content.** Every zone on the card is labeled before the content appears — "owner" before the name, implicit zone headers before the address. On the web, every section should have a small eyebrow label before the headline.

### Grid System

```css
/* Page max-width */
--page-max: 1280px;
--page-pad: clamp(1.5rem, 5vw, 4rem);

/* Card border system — use on contained sections */
.card-border {
  border: 1px solid rgba(200, 184, 130, 0.4);
  box-shadow: 0 0 0 3px #141414,
              0 0 0 4px rgba(200, 184, 130, 0.2);
  /* Creates double-border effect matching the card's ruled frame */
}

/* Section rhythm */
--section-pad-lg: clamp(4rem, 8vw, 7rem);
--section-pad-md: clamp(2.5rem, 5vw, 4rem);
```

### The Placard Component

The card's most direct translation to a web component — a contained block with double-rule border, eyebrow label, and bold content. Use for testimonials, stat callouts, featured service blocks.

```
┌─────────────────────────────────┐  ← outer rule (champagne, 1px)
│                                 │
│  LABEL                          │  ← eyebrow: dark band, small caps
├─────────────────────────────────┤  ← inner rule (champagne, dimmer)
│                                 │
│  Content                        │
│                                 │
└─────────────────────────────────┘  ← closing rule
```

---

## Component Library

### Navigation

```
Background: #1E1C18
Height: 56px
Logo: "A-TEAM ASPHALT CO." in Zilla Slab, red (#C0200A), ~18px
Logo sub: "Lic. #TEAMAAL855JE" in Barlow Condensed, champagne, 9px — builds in the license number as a trust signal
Links: Oswald 11px, 0.12em tracking, uppercase, muted gold (#9A8C60)
Links hover: champagne (#C8B882)
CTA button: Red bg, off-white text, no border-radius (square corners)
Phone: Barlow 13px, champagne, prominent in nav-right
Border-bottom: 1px, double-rule treatment
```

### Hero Section

```
Background: #141414 with diagonal stripe texture (ink, very low opacity)
Layout: Two-column — content left, photo/stat panel right
Headline: Zilla Slab Black, red, large — company name as hero
Sub-headline: Oswald 600, champagne, tracking 0.04em
Body: Barlow 15px, off-white
CTA: Red primary button + champagne ghost button
Trust flag: "over 20 years · Pierce County · WA Lic. #TEAMAAL855JE"
         Oswald, 9px, tracked, muted — beneath the headline
```

### Trust / Credential Band

Directly from the card's credential zone. A full-width dark band with the key trust signals in labeled placard style.

```
Background: #1E1C18
Border-top + border-bottom: double-rule in champagne
Items: flex row, each item has:
  - small eyebrow label (Oswald, muted, 9px)
  - value (Oswald 600, champagne, 14px)
  Separated by thin vertical rules

Items:
  "EXPERIENCE" / 20+ Years
  "LICENSED" / TEAMAAL855JE
  "COVERAGE" / Pierce & King Co.
  "ESTIMATES" / Always Free
```

### Tagline Treatment

"As good as the best, different than the rest" — do not correct the typo in formal contexts, but use the corrected version on the website. Consider treating it as the card does: smaller, angled or rotated label text set apart from the main identity.

```
Font: Zilla Slab, italic, ~18px
Color: Champagne (#C8B882)
Treatment: Set below the company name with a short red rule above it
          Like a subtitle on a title card
```

### Service Cards

```
Border: Double-rule in champagne (card-border mixin)
Background: #1E1C18
Eyebrow: Zone label (Oswald, muted, 10px, uppercase, 0.18em)
Title: Oswald 700, champagne, 20px
Body: Barlow 13px, #9A9590
Hover: Red top-border sweep (2px, animated right to full width)
```

### Buttons

```css
/* Primary — Red */
.btn-primary {
  background: #C0200A;
  color: #F2EFE6;
  font-family: 'Oswald', sans-serif;
  font-size: 12px;
  font-weight: 600;
  letter-spacing: 0.12em;
  text-transform: uppercase;
  padding: 0.8rem 2rem;
  border: none;
  border-radius: 0;           /* Square — matches card aesthetic */
  cursor: pointer;
  transition: background 0.15s;
}
.btn-primary:hover { background: #D93010; }

/* Secondary — Champagne ghost */
.btn-secondary {
  background: transparent;
  color: #C8B882;
  font-family: 'Oswald', sans-serif;
  font-size: 12px;
  font-weight: 600;
  letter-spacing: 0.12em;
  text-transform: uppercase;
  padding: 0.8rem 2rem;
  border: 1.5px solid #C8B882;
  border-radius: 0;
  cursor: pointer;
  transition: background 0.15s, color 0.15s;
}
.btn-secondary:hover {
  background: #C8B882;
  color: #141414;
}
```

### Form Inputs

```css
/* Matches the card's info zone aesthetic */
input, textarea, select {
  background: #1E1C18;
  border: 1px solid rgba(200, 184, 130, 0.3);
  border-radius: 0;
  color: #F2EFE6;
  font-family: 'Barlow', sans-serif;
  font-size: 14px;
  padding: 0.7rem 1rem;
  transition: border-color 0.15s;
}
input:focus, textarea:focus, select:focus {
  border-color: #C8B882;
  outline: none;
}
input::placeholder { color: #6A6560; }
```

---

## Texture & Atmosphere

### Background Texture

The card has a matte black surface with subtle depth. On the web, avoid pure black. Use `#141414` as base and layer a subtle diagonal stripe pattern at very low opacity — echoing the card's physical texture without being literal about it.

```css
.dark-section {
  background-color: #141414;
  background-image: repeating-linear-gradient(
    -55deg,
    transparent,
    transparent 40px,
    rgba(200, 184, 130, 0.018) 40px,
    rgba(200, 184, 130, 0.018) 41px
  );
}
```

### The Double-Rule Border

The card's most distinctive visual element. Use it as a system — not just on one component, but consistently throughout wherever there's a contained block.

```css
/* Full double-rule border */
.placard {
  border: 1px solid rgba(200, 184, 130, 0.5);
  outline: 3px solid #141414;
  outline-offset: -5px;
  /* Achieves the card's nested-border look using outline */
}

/* Single top-bottom rule (section divider) */
.section-rule {
  border-top: 1px solid rgba(200, 184, 130, 0.45);
  box-shadow: inset 0 1px 0 rgba(200, 184, 130, 0.15);
}
```

### The Rotated Wordmark

The card runs the company name vertically (rotated -90°). This is a memorable spatial choice — on the web it can be echoed in a sidebar element, a section label running vertically along the left edge, or a decorative background watermark.

```css
/* Vertical section label — echoes card's rotated company name */
.vertical-label {
  writing-mode: vertical-rl;
  text-orientation: mixed;
  transform: rotate(180deg);
  font-family: 'Zilla Slab', serif;
  font-weight: 700;
  font-size: 11px;
  letter-spacing: 0.15em;
  text-transform: uppercase;
  color: rgba(200, 184, 130, 0.25);
}
```

---

## Key Corrections from Earlier Concepts

Now that we have the actual card, several things from our earlier explorations need updating:

| Earlier assumption | Reality | Update |
|---|---|---|
| Eastern Washington / Spokane | Sumner WA, Pierce County — Tacoma metro | Geography, climate references, area codes all change |
| Cool chrome silver | Warm champagne gold | Entire secondary palette warms significantly |
| Condensed gothic / military stencil | Slab serif / western poster | Typography direction shifts to Zilla Slab or similar |
| "A-Team Asphalt" | "A-Team Asphalt Co." | The "Co." matters — slightly more formal |
| Invented tagline | "As good as the best, different than the rest" | Real tagline, use it |
| 30 years | 20+ years | Update the credential |
| (509) area code | (253) area code | Tacoma not Spokane |
| Invented license | TEAMAAL855JE | Real license number |
| Invented logo mark | The card suggests no existing logomark | The type treatment IS the identity — respect that |

---

## What to Preserve from the Card

Some things the card gets exactly right that the website should honor:

1. **The "Co." in the name.** Don't drop it. It adds a note of old-school formality that fits the slab-serif identity.

2. **The license number as a credential.** The card puts it prominently — `Lic. #TEAMAAL855JE`. On the web it should appear in the nav sub-line, the hero trust strip, and the footer. It's not fine print; it's a badge.

3. **The placard/zone structure.** The card's division of information into labeled zones is a clear identity expression. The web should echo this — not mimic it literally, but use the same logic: labeled containers, defined zones, ruled separators.

4. **The warmth of the gold.** The champagne metallic is the card's soul. It reads as *earned* — not flashy chrome, not corporate silver, but something that's been around and proved itself. That warmth should come through in every element of the web palette.

5. **Gary's name is prominent.** The card names him as "owner" with significant visual weight. The website should have a clear "about" or owner presence — customers in the residential and light commercial space in Pierce County want to know there's a real person behind the company.

---

## Sumner / Pierce County Context

This matters for copy and imagery direction:

- **Climate:** Western Washington — wet, gray, significant rainfall. Asphalt freeze-thaw is less about extreme cold and more about constant moisture and temperature fluctuation. Reference Pacific Northwest conditions, not mountain winters.
- **Market:** Tacoma metro suburbs — Sumner, Auburn, Puyallup, Bonney Lake, Enumclaw corridor. Mix of residential, light industrial, and commercial. Boeing supply chain presence. Agricultural transition zones.
- **Competitive landscape:** More dense than Eastern WA — more contractors, more price competition, more savvy commercial buyers.
- **Imagery direction:** Green — evergreens, overcast skies, wet asphalt, suburban residential streets. Not desert, not mountains, not open plains.

---

*A-Team Asphalt Co. · Sumner, WA · Lic. #TEAMAAL855JE · Firefly Software UI reference · 2026*
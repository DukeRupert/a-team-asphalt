package services

// Service holds all content for a single service detail page.
type Service struct {
	Slug        string
	Name        string
	Tagline     string
	Category    string
	CategoryNum string
	MetaDesc    string
	Pitch       []string
	Includes    []string
	WhyUs       []WhyPoint
	ContactType string // pre-fills contact form project_type
	PhotoDesc   string // describes ideal stock photo for HTML comment
	PhotoURL    string // path to hero image (e.g. "/static/img/asphalt-paving-800.webp"), empty = placeholder

	// Optional sections — populated for SEO landing pages, empty for standard services.
	HeroEyebrow    string         // overrides default category eyebrow (e.g. "Driveway Paving · Puget Sound, WA")
	Process         []ProcessStep  // "How We Do It" numbered steps
	GoodToKnow      []string       // educational callouts
	ServiceArea     string         // city list paragraph
	RelatedServices []RelatedLink  // cross-links to other service pages
	CustomCTAHead   string         // overrides "Ready to get started?"
	CustomCTABody   string         // overrides default CTA body
}

// WhyPoint is a single trust-building point.
type WhyPoint struct {
	Title string
	Desc  string
}

// ProcessStep is a numbered step in the "How We Do It" section.
type ProcessStep struct {
	Num   string // "01", "02", etc.
	Title string
	Desc  string
}

// RelatedLink points to another service page.
type RelatedLink struct {
	Name string
	Slug string
	Desc string
}

// BySlug returns the service matching the given slug, or nil.
func BySlug(slug string) *Service {
	s, ok := catalog[slug]
	if !ok {
		return nil
	}
	return &s
}

// All returns every service in display order.
func All() []Service {
	return order
}

var catalog = make(map[string]Service)

func init() {
	for _, s := range order {
		catalog[s.Slug] = s
	}
}

var order = []Service{
	// ── 01 · ASPHALT & CONCRETE ──
	{
		Slug:        "asphalt-paving",
		Name:        "Asphalt Paving & Installation",
		Tagline:     "New pavement that lasts — because we do the prep other crews skip.",
		Category:    "Asphalt & Concrete",
		CategoryNum: "01",
		MetaDesc:    "Professional asphalt paving for driveways, parking lots, and roads in the Puget Sound area. Proper base prep, laid to spec, compacted right. Free estimates: (253) 365-1283.",
		Pitch: []string{
			"New driveways, parking lots, and roads. We prep the base correctly, lay to spec, and compact it properly — because that's what makes asphalt last twenty years instead of five.",
			"Every paving job starts with proper base work. We don't show up and pour over whatever's there. We grade it, compact it, and make sure water goes where it should before the first load of hot-mix arrives.",
			"One crew handles the entire job from demolition through final compaction. No subcontractors, no hand-offs between companies, no surprises on the invoice.",
		},
		Includes: []string{
			"Residential driveways and private roads",
			"Commercial parking lots and access roads",
			"Base preparation and grading",
			"Hot-mix asphalt installation",
			"Compaction and finish rolling",
			"Clean-up and site restoration",
		},
		WhyUs: []WhyPoint{
			{Title: "Proper base prep", Desc: "We grade and compact before we pave. That's why our work lasts."},
			{Title: "Owner on every job", Desc: "Gary or Don is on-site. The person who quoted you runs the crew."},
			{Title: "One crew, start to finish", Desc: "No subs, no hand-offs. The same team from demo through compaction."},
		},
		ContactType: "Asphalt paving",
		PhotoDesc:   "Fresh hot-mix asphalt being laid on a residential driveway or commercial lot. Roller compactor visible, crew in hi-vis gear, steam rising from black surface. PNW setting.",
		PhotoURL:    "/static/img/asphalt-paving-800.webp",
	},
	{
		Slug:        "asphalt-repair",
		Name:        "Asphalt Repair & Patching",
		Tagline:     "Honest assessment — we'll tell you what needs replaced and what doesn't.",
		Category:    "Asphalt & Concrete",
		CategoryNum: "01",
		MetaDesc:    "Asphalt repair and patching in the Puget Sound area. Potholes, cracks, and failing sections fixed right. No upsell to replacement when repair will do. Free estimates.",
		Pitch: []string{
			"Potholes, cracks, and failing sections. We'll tell you what actually needs to come out and what can be patched. No upsell to replacement when repair will do.",
			"Not every damaged surface needs a full tearout. We assess the base condition and give you a straight answer about what repair will hold and what's too far gone. You get the fix that fits the problem.",
		},
		Includes: []string{
			"Pothole repair and infrared patching",
			"Cut-and-patch of failed sections",
			"Alligator cracking removal and replacement",
			"Edge repair along curbing and transitions",
			"Base assessment and honest recommendation",
		},
		WhyUs: []WhyPoint{
			{Title: "No unnecessary upsells", Desc: "If a patch will hold, we'll patch it. We tell you the truth about what your pavement needs."},
			{Title: "Same-crew quality", Desc: "Repairs get the same materials and compaction standards as our new installs."},
			{Title: "30 years of reading pavement", Desc: "We know what's a surface problem and what's a base failure. That saves you money."},
		},
		ContactType: "Asphalt paving",
		PhotoDesc:   "Close-up of asphalt patching in progress — fresh black hot-mix next to existing weathered pavement. Hand tools or small roller visible.",
	},
	{
		Slug:        "concrete",
		Name:        "Concrete Installation & Repair",
		Tagline:     "Concrete where it belongs. Done to code. Done to last.",
		Category:    "Asphalt & Concrete",
		CategoryNum: "01",
		MetaDesc:    "Curbs, sidewalks, slabs, ADA ramps, and flatwork throughout the Puget Sound. A-Team handles concrete work alongside asphalt projects or standalone. Free estimates — Tacoma, Puyallup, Auburn, and surrounding areas.",
		HeroEyebrow: "Concrete Installation & Repair · Puget Sound, WA",
		Pitch: []string{
			"Not everything is asphalt. Curbs, sidewalks, garage aprons, ADA ramps, retaining footings, and concrete pads are part of most commercial and many residential paving projects — and they need a contractor who can handle both materials, not two separate crews trying to coordinate around each other.",
			"We do the concrete work. That means you get one point of contact, one schedule, and transitions between asphalt and concrete surfaces that are built together and fit together. For commercial property managers dealing with ADA compliance, cracked sidewalks, or failing curb sections, we can handle the concrete scope directly.",
		},
		Includes: []string{
			"Concrete curb and gutter (new installation and replacement)",
			"Sidewalk installation and repair",
			"ADA ramp installation (to current federal and WA state standards)",
			"Concrete flatwork — slabs, pads, aprons, approaches",
			"Garage approach and transition slabs",
			"Concrete removal and haul-off",
			"Retaining curbs and wheel stops",
			"Concrete patching for utility cuts and localized failures",
		},
		Process: []ProcessStep{
			{Num: "01", Title: "Scope and code review", Desc: "For sidewalks, ADA ramps, and curb work, we verify the applicable municipal standards before forming anything. Cities across the Puget Sound each have their own requirements for slope tolerances, panel dimensions, joint placement, and surface finish. We build to those requirements."},
			{Num: "02", Title: "Demo and subgrade prep", Desc: "Existing failed concrete comes out cleanly. We excavate to the correct depth, compact the subgrade, and place base material appropriate to the application and load. Concrete poured on an unprepared subgrade will crack."},
			{Num: "03", Title: "Forming", Desc: "Forms are set to the correct grade and slope. For ADA ramps, slope tolerances are tight — a ramp that exceeds the allowable running or cross slope is non-compliant even if it looks fine. We set forms to spec, not to eye."},
			{Num: "04", Title: "Pour and finish", Desc: "We pour, strike, and finish the concrete to the required surface texture — broom finish for pedestrian surfaces, smooth for slabs where specified. Control joints are cut at the right spacing to direct cracking where it won't matter."},
			{Num: "05", Title: "Cure and protection", Desc: "Concrete needs time to gain strength. We apply curing compound and protect the surface from traffic for the appropriate period. Rushing traffic onto fresh concrete is one of the most common causes of premature surface failure."},
		},
		WhyUs: []WhyPoint{
			{Title: "Municipal-grade work", Desc: "We build to code. When your project needs permits and inspection, it passes."},
			{Title: "Proper cure time", Desc: "We don't rush the finish. Concrete needs time to reach full strength."},
			{Title: "One contractor for everything", Desc: "Asphalt and concrete on the same project? One crew, one invoice."},
		},
		GoodToKnow: []string{
			"ADA compliance for parking lots isn't just about accessible stall markings — it extends to the accessible route from the stall to your building entrance. If that route crosses a sidewalk with an out-of-tolerance cross slope or a trip hazard, you have a compliance exposure.",
			"In Washington, property owners are responsible for maintaining sidewalks adjacent to their property. Cities can compel repairs or do them and lien the property. Getting it done voluntarily is significantly less expensive.",
			"When an asphalt project involves transitions to concrete — garage aprons, curb cuts, sidewalk connections — the transition detail matters. When we're doing both materials, we control those details.",
		},
		ServiceArea:     "We provide concrete installation and repair throughout the Puget Sound — Tacoma, Puyallup, Auburn, Sumner, Federal Way, SeaTac, Covington, Maple Valley, Bellevue, and the greater Seattle area.",
		RelatedServices: []RelatedLink{
			{Name: "Parking Lot Paving", Slug: "parking-lot-paving", Desc: "Full lot projects including concrete curb and flatwork"},
			{Name: "Striping & Bollards", Slug: "striping", Desc: "ADA compliance markings alongside concrete work"},
			{Name: "Site Prep & Grading", Slug: "site-prep", Desc: "Subgrade prep before concrete or asphalt"},
		},
		CustomCTAHead: "Need concrete work alongside a paving project — or standalone?",
		CustomCTABody: "Call us. We handle the full scope or just the concrete. No obligation.",
		ContactType:   "Concrete / pavers / hardscape",
		PhotoDesc:     "Fresh concrete being poured or finished — smooth wet surface with forms visible. Worker with float or trowel.",
		PhotoURL:      "/static/img/concrete-pour-800.webp",
	},
	{
		Slug:        "milling-overlay",
		Name:        "Asphalt Milling & Overlay",
		Tagline:     "New surface. No tearout. A fraction of the cost.",
		Category:    "Asphalt & Concrete",
		CategoryNum: "01",
		MetaDesc:    "Asphalt milling and overlay is the cost-effective alternative to full replacement when your base is structurally sound. A-Team serves the Puget Sound — free estimates for driveways, parking lots, and commercial surfaces.",
		HeroEyebrow: "Asphalt Milling & Overlay · Puget Sound, WA",
		Pitch: []string{
			"Most asphalt doesn't fail from the bottom up — it fails from the surface down. UV oxidation, water infiltration, traffic wear, and Washington's freeze-thaw cycles all work on the top two inches first. If the base beneath it is still structurally sound, tearing the whole thing out is unnecessary.",
			"Milling removes the failed surface layer to a precise, consistent depth. The exposed base is cleaned and inspected, any localized failures are repaired, and fresh hot mix asphalt goes down over a surface that's actually ready to bond to it. The result is effectively new asphalt — for significantly less than full replacement.",
			"It's not the right answer for every situation. If the base has failed, an overlay just delays the inevitable and wastes money. We assess every job honestly before recommending it.",
		},
		Includes: []string{
			"Milling (grinding and removal of existing surface layer)",
			"Base inspection and spot repair of sub-base failures",
			"Tack coat application for proper bonding",
			"Hot mix asphalt overlay — typically 2 inches",
			"Edge and transition work",
			"Haul-off and recycling of milled material",
			"Striping and finish work",
		},
		Process: []ProcessStep{
			{Num: "01", Title: "Assessment — is this a milling candidate?", Desc: "We evaluate the base: type and extent of cracking, whether failures are surface or structural, and whether the grade and drainage are correct. If the base is compromised, we'll tell you — because an overlay over a bad base will mirror the same failures within a season or two."},
			{Num: "02", Title: "Milling", Desc: "A milling machine grinds the existing surface down to a consistent depth — typically 2 inches — without disturbing the base. This removes cracking, oxidation, and irregularities, leaving a textured surface that bonds well to new asphalt. The milled material is loaded and removed for recycling."},
			{Num: "03", Title: "Base repair", Desc: "Once the old surface is off, we can see exactly what the base looks like. Any soft spots, base failures, or drainage problems that weren't visible get addressed now — not after the new asphalt is down."},
			{Num: "04", Title: "Tack coat", Desc: "Asphalt emulsion is applied to the milled surface before paving. It's the bonding agent between old and new. This step is sometimes skipped by contractors looking to save time. Skipping it means the overlay can delaminate under traffic. We don't skip it."},
			{Num: "05", Title: "Overlay and compaction", Desc: "Fresh hot mix asphalt is laid at the right temperature and immediately compacted. Edges are cut clean and transitions to adjacent surfaces are finished properly."},
		},
		WhyUs: []WhyPoint{
			{Title: "Honest assessment first", Desc: "We'll tell you if overlay will work or if the base needs attention. No guessing."},
			{Title: "Cost-effective", Desc: "When the base is solid, overlay saves 40-60% versus full replacement."},
			{Title: "Clean transitions", Desc: "We grade the edges properly so water drains and the surface meets curbs cleanly."},
		},
		GoodToKnow: []string{
			"The tell-tale sign that milling and overlay is the right call: surface cracking without the alligator pattern. Alligator cracking — the web-like, interconnected pattern — indicates base failure. Overlay won't fix it.",
			"An overlay typically adds about 2 inches to the surface height. At transitions — garage doors, drainage grates, curb cuts, and ADA ramps — adjustment may be needed. We flag these during assessment.",
			"Milled asphalt is 100% recyclable and commonly used as aggregate in new asphalt mix. Choosing milling over full replacement is the more environmentally responsible approach.",
		},
		ServiceArea:     "We provide milling and overlay services for driveways, parking lots, and commercial surfaces throughout the Puget Sound — Tacoma, Puyallup, Auburn, Federal Way, SeaTac, Sumner, Covington, Maple Valley, Bellevue, and greater Seattle.",
		RelatedServices: []RelatedLink{
			{Name: "Parking Lot Paving", Slug: "parking-lot-paving", Desc: "For lots that need full base replacement"},
			{Name: "Driveway Paving", Slug: "driveway-paving", Desc: "New installation or full tearout and repave"},
			{Name: "Crack Filling", Slug: "crack-filling", Desc: "Address surface cracks before they become structural"},
		},
		CustomCTAHead: "Not sure if your pavement qualifies for overlay?",
		CustomCTABody: "Call us. We'll come out, look at the base condition, and give you a straight answer — milling and overlay, full replacement, or just maintenance. No obligation.",
		ContactType:   "Asphalt paving",
		PhotoDesc:     "Milling machine removing top layer of asphalt, exposing rough textured base. Fresh overlay being laid in background or adjacent area.",
		PhotoURL:      "/static/img/asphalt-paving-800.webp",
	},
	{
		Slug:        "driveway-paving",
		Name:        "Driveway Paving",
		Tagline:     "A driveway that holds — not one that looks good for a year and fails in three.",
		Category:    "Asphalt & Concrete",
		CategoryNum: "01",
		MetaDesc:    "A-Team Asphalt installs and repairs residential driveways throughout the Puget Sound. Proper base work, clean edges, free estimates. Serving Puyallup, Tacoma, Auburn, Federal Way, and surrounding areas.",
		HeroEyebrow: "Driveway Paving · Puget Sound, WA",
		Pitch: []string{
			"Your driveway takes more abuse than most people realize. Vehicle weight, freeze-thaw cycling, rainwater drainage, surface oxidation — it handles all of it, year after year. Done right, an asphalt driveway lasts twenty to twenty-five years with basic maintenance. Done wrong, you're calling someone back in five.",
			"The difference almost always comes down to what happens before the asphalt is laid. Subgrade preparation, base depth, compaction, and drainage grade are where driveways succeed or fail. We don't cut those corners. We've seen enough failed driveways to know exactly what causes them — and it's always the prep work.",
		},
		Includes: []string{
			"New driveway installation on bare ground or over gravel",
			"Full driveway replacement (tearout and repave)",
			"Driveway extension or widening",
			"Turnarounds and parking aprons",
			"Proper edge work and border finish",
			"Grading and drainage to move water away from your home",
			"Transitions to existing concrete, garage slabs, or street aprons",
		},
		Process: []ProcessStep{
			{Num: "01", Title: "Site evaluation and grade planning", Desc: "Before anything is excavated, we look at how water currently moves on your property. Your driveway needs to drain away from your foundation and not pool at the low end. We plan the grade before we dig."},
			{Num: "02", Title: "Excavation and subgrade prep", Desc: "We remove existing material to the proper depth for your soil conditions and traffic load. A standard residential driveway needs a minimum of 4 inches of compacted base — more on softer soils. Skimping on depth is the most common reason driveways fail early."},
			{Num: "03", Title: "Base installation and compaction", Desc: "We bring in the right crushed rock, spread it in lifts, and compact each layer before the next goes down. Compaction matters. Uncompacted base settles unevenly and takes the asphalt surface with it."},
			{Num: "04", Title: "Asphalt lay and finish", Desc: "Hot mix asphalt goes down at the right temperature and is compacted immediately. We cut clean edges and finish transitions to your garage, street, or existing concrete. No ragged borders, no mismatched heights."},
		},
		WhyUs: []WhyPoint{
			{Title: "Proper base prep", Desc: "We grade and compact before we pave. That's why our work lasts."},
			{Title: "Owner on every job", Desc: "Gary or Don is on-site. The person who quoted you runs the crew."},
			{Title: "One crew, start to finish", Desc: "No subs, no hand-offs. The same team from demo through compaction."},
		},
		GoodToKnow: []string{
			"New asphalt needs to cure before it's sealed — typically 6 to 12 months. The oils in fresh asphalt need to off-gas first. Any contractor who offers to sealcoat your brand-new driveway the same day is doing you a disservice. We'll tell you when it's ready.",
			"Puget Sound winters are wet, not brutal — but freeze-thaw cycles still happen, especially in elevated areas of Puyallup, Maple Valley, and Covington. Water that gets under unsupported edges will eventually heave them. We compact the subgrade to minimize this.",
			"If your old driveway is failing along the edges first, that's almost always a base problem, not a surface problem. An overlay will look fine for a season, then crack along the same lines. We'll tell you honestly whether your driveway is a mill-and-overlay candidate or whether the base needs to come out.",
		},
		ServiceArea:     "We install and replace residential driveways throughout the Puget Sound — Puyallup, Sumner, Auburn, Federal Way, Tacoma, Covington, Maple Valley, SeaTac, and the greater Seattle area including Bellevue and Issaquah.",
		RelatedServices: []RelatedLink{
			{Name: "Asphalt Milling & Overlay", Slug: "milling-overlay", Desc: "When the base is solid but the surface is shot"},
			{Name: "Sealcoating", Slug: "sealcoating", Desc: "Protect your new driveway after the first season"},
			{Name: "Crack Filling", Slug: "crack-filling", Desc: "Get ahead of small problems before they become big ones"},
		},
		CustomCTAHead: "Get a straight answer on your driveway.",
		CustomCTABody: "Call or fill out the form. We'll come out, look at what you've got, and tell you exactly what it needs — and whether repair or replacement makes more sense.",
		ContactType:   "Asphalt paving",
		PhotoDesc:     "Fresh asphalt driveway being laid at a residential home. Clean edges, roller compactor, crew working. PNW residential neighborhood with trees.",
		PhotoURL:      "/static/img/fresh-asphalt-800.webp",
	},
	{
		Slug:        "parking-lot-paving",
		Name:        "Parking Lot Paving",
		Tagline:     "Your parking lot is the first thing customers see. It shouldn't be an obstacle.",
		Category:    "Asphalt & Concrete",
		CategoryNum: "01",
		MetaDesc:    "Commercial parking lot installation, repair, and maintenance throughout the Puget Sound. A-Team handles paving, sealcoating, striping, and ADA compliance. Free estimates for property managers and business owners.",
		HeroEyebrow: "Commercial Parking Lot Paving · Puget Sound, WA",
		Pitch: []string{
			"A parking lot is infrastructure. For most commercial property owners and managers, it's also a liability — potholes, uneven surfaces, and faded or missing ADA markings all create real exposure. Deferred maintenance turns a manageable resurfacing job into a full tearout that costs three times as much.",
			"We work with property managers, business owners, HOAs, and commercial developers throughout the Puget Sound. Whether you need a new lot built from scratch, an existing lot replaced, or a maintenance program that keeps a good lot from becoming a bad one — we handle the full scope in-house. No subcontractors, no surprises.",
		},
		Includes: []string{
			"New parking lot construction (full base build and pave)",
			"Full lot replacement (demolition, base rebuild, repave)",
			"Asphalt overlay (mill and repave where base is sound)",
			"Pothole and patch repair",
			"Sealcoating (standalone or alongside paving)",
			"Striping — new layout or re-stripe, ADA compliant",
			"Bollard installation and repair",
			"Curb and gutter work",
			"Drainage grading and surface contouring",
			"Snow removal and de-icing contracts",
		},
		Process: []ProcessStep{
			{Num: "01", Title: "Site assessment", Desc: "We walk the lot, evaluate base condition, map drainage patterns, and look at how traffic moves through the space. A lot that drains poorly or has a base in failure needs different work than one that just needs a surface refresh. We tell you which situation you're in before any scope is set."},
			{Num: "02", Title: "Scope and design", Desc: "For new construction or full replacement, we plan layout for drainage, traffic flow, stall count, and ADA compliance before breaking ground. For maintenance, we define exactly what needs done and what can wait — commercial property managers don't benefit from being sold more than their lot needs."},
			{Num: "03", Title: "Base work", Desc: "Commercial lots see heavier and more sustained loads than driveways. We spec base depth to the load — a lot that takes delivery trucks needs more base than one that handles passenger cars. We don't use the same spec for both."},
			{Num: "04", Title: "Paving", Desc: "Hot mix asphalt laid at the right temperature, compacted to spec, with clean edges and smooth transitions. Large lots are phased to minimize disruption to your operations."},
			{Num: "05", Title: "Finish work", Desc: "Striping, bollards, signage, and curb work go in after the surface is ready. We don't rush the finish — a freshly paved lot with sloppy lines or misaligned ADA stalls is still a problem."},
		},
		WhyUs: []WhyPoint{
			{Title: "Full scope, one crew", Desc: "Paving, sealcoating, striping, bollards, signage — we handle it all. One invoice."},
			{Title: "ADA compliance built in", Desc: "Stall dimensions, slopes, signage, and accessible routes done to current standards."},
			{Title: "Commercial experience", Desc: "We work with property managers, HOAs, and developers regularly. We know the requirements."},
		},
		GoodToKnow: []string{
			"ADA compliance isn't optional — and it's more specific than most property owners realize. Accessible stall dimensions, slope tolerances, signage placement, and the accessible route from stalls to your building entrance all have requirements. We stripe to current ADA standards and can flag layout issues before they become a compliance problem.",
			"If your lot is showing widespread surface cracking but the base is still sound, a mill and overlay costs significantly less than full replacement and delivers a surface that will last another fifteen to twenty years. We'll tell you honestly whether your base qualifies.",
			"For property managers handling multiple sites, a scheduled maintenance program — annual crack filling, sealcoating on a 3-5 year cycle, re-stripe as needed — is the most cost-effective way to extend pavement life and avoid emergency repair costs.",
		},
		ServiceArea:     "We serve commercial property owners and managers throughout the Puget Sound — Tacoma, Seattle, Auburn, Federal Way, Bellevue, Puyallup, SeaTac, Renton, Kent, Covington, and the surrounding region.",
		RelatedServices: []RelatedLink{
			{Name: "Asphalt Milling & Overlay", Slug: "milling-overlay", Desc: "When the base is good and the surface needs replacing"},
			{Name: "Sealcoating", Slug: "sealcoating", Desc: "Protect your lot and extend its service life"},
			{Name: "Parking Lot Striping", Slug: "striping", Desc: "ADA-compliant layouts, re-stripe, and new marking"},
		},
		CustomCTAHead: "Property managers: let's talk about your lot.",
		CustomCTABody: "We work with commercial property owners and managers regularly. Call us or fill out the form — we'll schedule a site visit, assess what your lot actually needs, and give you a straight quote. No obligation.",
		ContactType:   "Asphalt paving",
		PhotoDesc:     "Aerial or wide shot of a freshly paved commercial parking lot. Clean black asphalt, fresh striping, ADA stalls visible. Commercial building in background.",
		PhotoURL:      "/static/img/asphalt-paving-800.webp",
	},

	// ── 02 · MAINTENANCE & PROTECTION ──
	{
		Slug:        "sealcoating",
		Name:        "Sealcoating",
		Tagline:     "Protect what you've already paid for.",
		Category:    "Maintenance & Protection",
		CategoryNum: "02",
		MetaDesc:    "Professional asphalt sealcoating for driveways and parking lots throughout the Puget Sound. Two-coat application built for Washington winters. Free estimates — Tacoma, Puyallup, Auburn, and surrounding areas.",
		HeroEyebrow: "Sealcoating · Puget Sound, WA",
		Pitch: []string{
			"Fresh asphalt is porous. The binder that holds it together starts oxidizing from the day it's laid — UV exposure, rainwater infiltration, petroleum drips from vehicles, and freeze-thaw cycling all work on it continuously. A surface that looks solid is actually degrading at the molecular level every season you leave it unsealed.",
			"Sealcoating creates a protective barrier that slows all of it down. Applied on schedule, it more than doubles the effective lifespan of your pavement and costs a fraction of what resurfacing or replacement does. The math is straightforward — but only if the sealcoating is done properly.",
			"We apply two coats because one isn't enough for this climate. And we don't seal over problems — if there are cracks that need filling or sections that need repair, those come first. Sealing over damage just traps it and accelerates failure beneath a surface that looks fine.",
		},
		Includes: []string{
			"Residential driveways (new and existing)",
			"Commercial parking lots",
			"Multi-family property lots and drives",
			"Industrial and warehouse surfaces",
			"Surface prep and cleaning before application",
			"Crack filling prior to sealing (performed or quoted separately)",
			"Two-coat application to manufacturer spec",
			"Cure time guidance and traffic restrictions",
		},
		Process: []ProcessStep{
			{Num: "01", Title: "Surface assessment", Desc: "We look at the pavement condition before quoting anything. If there are structural problems — base failures, significant cracking, alligator patterns — sealcoating won't fix them and we'll tell you so. We only seal surfaces that are ready to be sealed."},
			{Num: "02", Title: "Cleaning and prep", Desc: "We pressure wash, blow out debris from cracks, and treat oil spots. Sealcoat doesn't bond properly to a dirty surface, and oil contamination will cause it to peel. Prep is where the application's lifespan is determined. Many contractors skip or rush this step. We don't."},
			{Num: "03", Title: "Crack filling (if needed)", Desc: "Any cracks wider than 1/8 inch get filled with hot-apply rubberized filler before the first coat goes down. Sealing over open cracks traps them — they'll telegraph through the sealcoat within a season. We address them first."},
			{Num: "04", Title: "First coat", Desc: "We apply the first coat uniformly and allow it to cure to the manufacturer's specification — not until it looks dry, but until it's chemically ready for a second application. Rushing this step is the most common reason sealcoating fails prematurely."},
			{Num: "05", Title: "Second coat", Desc: "The second coat goes on after the first is fully cured. Two coats provide the coverage depth and durability the Pacific Northwest climate requires. One coat is a corner cut."},
		},
		WhyUs: []WhyPoint{
			{Title: "Two coats, always", Desc: "Single-coat sealcoating is a shortcut. We apply two because that's what Washington weather demands."},
			{Title: "Proper prep", Desc: "We clean the surface and fill cracks first. Sealcoat over dirt and open cracks is money wasted."},
			{Title: "Right conditions", Desc: "We don't sealcoat in the rain or below 50°F. If the weather's wrong, we reschedule."},
		},
		GoodToKnow: []string{
			"New asphalt needs 6 to 12 months to cure before it should be sealed. The oils in fresh asphalt need to off-gas first. If a contractor offers to sealcoat a brand-new driveway immediately after paving, that's a problem — you'll be sealing in oils that should have evaporated.",
			"There is such a thing as too much sealcoating. If sealcoat builds up too thick from multiple applications without proper assessment, it can crack and flake off in sheets. Before reapplying on an older surface, we evaluate what's already there.",
			"The best time of year to sealcoat in the Puget Sound is late spring through early fall — when surface temperatures are consistently above 50°F and no rain is forecast for 24-48 hours after application.",
			"For property managers with multiple sites, a scheduled sealcoating rotation — typically on a 3-5 year cycle depending on traffic — is the most cost-effective way to maintain pavement across a portfolio.",
		},
		ServiceArea:     "We provide sealcoating for residential driveways and commercial parking lots throughout the Puget Sound — Sumner, Puyallup, Tacoma, Auburn, Federal Way, SeaTac, Covington, Maple Valley, Bellevue, and greater Seattle.",
		RelatedServices: []RelatedLink{
			{Name: "Crack Repair", Slug: "crack-filling", Desc: "Address cracks before they get sealed over"},
			{Name: "Parking Lot Paving", Slug: "parking-lot-paving", Desc: "When the surface is past saving"},
			{Name: "Driveway Paving", Slug: "driveway-paving", Desc: "New installation when the base has failed"},
		},
		CustomCTAHead: "Ready to protect your pavement?",
		CustomCTABody: "Call or send us a message. We'll assess the surface, tell you what it needs, and give you a straight quote. No obligation.",
		ContactType:   "Sealcoating / repair",
		PhotoDesc:     "Worker applying sealcoat with squeegee or spray rig. Wet glossy black surface contrasting with unsealed faded gray asphalt.",
		PhotoURL:      "/static/img/fresh-asphalt-800.webp",
	},
	{
		Slug:        "crack-filling",
		Name:        "Crack Repair & Filling",
		Tagline:     "A $300 crack fill today. Or a $3,000 repair next year.",
		Category:    "Maintenance & Protection",
		CategoryNum: "02",
		MetaDesc:    "Professional asphalt crack filling and repair for driveways and parking lots across the Puget Sound. Stop water infiltration before it becomes a base failure. Free estimates — serving Tacoma, Puyallup, Auburn, and surrounding areas.",
		HeroEyebrow: "Crack Repair & Filling · Puget Sound, WA",
		Pitch: []string{
			"Asphalt cracks because it moves. It expands in heat, contracts in cold, flexes under load, and shifts when the base beneath it settles or softens. Small surface cracks are normal — they're not structural failure. But every crack is an entry point for water.",
			"Water gets into a crack, works its way to the base, softens the subgrade, and starts the cycle that ends in potholes and base failure. In the Puget Sound, where rain is consistent and freeze-thaw events happen more than people expect, that cycle moves faster than it does in drier climates.",
			"The right response to a crack depends on what kind it is and what caused it. Surface cracks (transverse and longitudinal) running perpendicular or parallel to the direction of paving are good candidates for filling — the base is intact. Edge cracks along the perimeter often indicate drainage or support issues. Block cracking signals an aged, brittle surface nearing end of life. And alligator cracking — the web-like interconnected pattern — means the base has failed. Filling won't fix it.",
			"We identify the type before recommending a fix — because crack filling is the right answer for surface cracking, and something else entirely is the right answer for structural failure.",
		},
		Includes: []string{
			"Crack routing and cleaning",
			"Hot-pour rubberized crack sealant",
			"Surface crack treatment (transverse, longitudinal, edge)",
			"Block crack assessment and treatment",
			"Joint sealing",
			"Assessment of underlying base condition",
		},
		Process: []ProcessStep{
			{Num: "01", Title: "Inspection and crack classification", Desc: "We walk the pavement and classify the cracking before touching anything. The type of crack determines the right fix. We don't recommend filling cracks that indicate base failure — it would be wasting your money."},
			{Num: "02", Title: "Crack preparation", Desc: "We clean cracks with a wire broom, compressed air, or a hot air lance to remove debris, vegetation, and loose material. Filler needs to bond to clean asphalt. Cracks that aren't properly cleaned fail quickly."},
			{Num: "03", Title: "Hot-apply rubberized crack filler", Desc: "The material is heated to approximately 380-400°F, poured into the crack, and tooled flush with the surface. Hot-apply filler penetrates the crack, bonds to the sidewalls, and remains flexible through temperature cycling. Cold-pour products are weaker and less durable. We don't use them."},
			{Num: "04", Title: "Cure and traffic", Desc: "The filler cures quickly — typically 30 to 60 minutes before foot traffic, a few hours before vehicle traffic. We'll give you specific guidance based on conditions on the day of the job."},
		},
		WhyUs: []WhyPoint{
			{Title: "Preventive, not cosmetic", Desc: "We fill cracks to stop water infiltration, not just to make the surface look better."},
			{Title: "Hot-pour sealant", Desc: "We use rubberized hot-pour material that flexes with temperature changes — not cold-pour from a jug."},
			{Title: "Honest about limits", Desc: "If cracks indicate base failure, we'll tell you. Filling over a failed base doesn't help."},
		},
		GoodToKnow: []string{
			"Cracks narrower than 1/8 inch generally shouldn't be filled — there isn't enough width for the filler to bond properly. Cracks that wide or wider should be addressed promptly.",
			"The best time to fill cracks in the Puget Sound is spring and fall, when temperatures are moderate and crack widths are near their median — not at their widest (winter) or narrowest (summer peak heat).",
			"If you're planning to sealcoat, crack filling comes first. Always. Sealing over open cracks traps them — they'll push through the sealcoat within a season.",
			"Crack filling is maintenance, not a permanent structural fix. For surfaces with widespread cracking or base issues, the honest answer is that filling individual cracks won't hold the pavement together. We'll tell you when repair has crossed the line into replacement territory.",
		},
		ServiceArea:     "We provide crack filling and repair for driveways, parking lots, and commercial surfaces throughout the Puget Sound — Puyallup, Sumner, Tacoma, Auburn, Federal Way, Covington, Maple Valley, SeaTac, and the greater Seattle area.",
		RelatedServices: []RelatedLink{
			{Name: "Sealcoating", Slug: "sealcoating", Desc: "Protect the repaired surface and the rest of the pavement"},
			{Name: "Asphalt Milling & Overlay", Slug: "milling-overlay", Desc: "When surface cracking is widespread"},
			{Name: "Parking Lot Paving", Slug: "parking-lot-paving", Desc: "When alligator cracking indicates base replacement is needed"},
		},
		CustomCTAHead: "Not sure what type of cracking you're dealing with?",
		CustomCTABody: "Send us a photo or call us out for a free assessment. We'll tell you what you're looking at and what it'll take to fix it.",
		ContactType:   "Sealcoating / repair",
		PhotoDesc:     "Close-up of crack filling in progress — hot-pour sealant being applied to clean crack in asphalt surface.",
	},
	{
		Slug:        "pressure-washing",
		Name:        "Pressure Washing",
		Tagline:     "Clean surface. Better results. Longer-lasting work.",
		Category:    "Maintenance & Protection",
		CategoryNum: "02",
		MetaDesc:    "Professional pressure washing for asphalt and concrete surfaces throughout the Puget Sound — surface prep before sealing, standalone lot cleaning, and more. Serving Tacoma, Puyallup, Auburn, and surrounding areas.",
		HeroEyebrow: "Pressure Washing · Puget Sound, WA",
		Pitch: []string{
			"A clean surface is the foundation of every maintenance service we provide. Sealcoat applied over oil stains, moss, and road debris won't bond properly — it peels within a season. Striping over a contaminated surface fades fast and lifts at the edges. Crack filler applied to debris-packed cracks doesn't seal the crack; it just sits on top of it.",
			"Pressure washing is how we prepare surfaces for the work that follows. It's also a standalone service for commercial property managers who want their lots, drives, and concrete surfaces looking professional — or who are dealing with moss, algae, or oil contamination that's affecting surface condition.",
		},
		Includes: []string{
			"Parking lot surface cleaning (asphalt and concrete)",
			"Driveway and access drive cleaning",
			"Concrete sidewalk and flatwork cleaning",
			"Oil and fuel spot treatment and washing",
			"Moss and algae removal from asphalt surfaces",
			"Pre-sealcoating surface prep washing",
			"Pre-striping surface prep",
			"Loading dock and industrial surface washing",
		},
		Process: []ProcessStep{
			{Num: "01", Title: "Surface assessment", Desc: "We identify areas of heavy contamination — oil spots, moss, standing stain patterns — before washing begins. Oil spots require a degreaser treatment prior to pressure washing; otherwise the oil gets spread rather than lifted."},
			{Num: "02", Title: "Degreaser treatment (where needed)", Desc: "Commercial degreaser is applied to oil and fuel contamination and allowed to dwell before washing. This breaks down the petroleum residue so it lifts cleanly rather than getting pushed around the surface."},
			{Num: "03", Title: "Pressure washing", Desc: "Hot water pressure washing is more effective than cold water for petroleum contamination and moss removal. We use equipment appropriate to the surface — high pressure on concrete, lower pressure on older or weathered asphalt where surface aggregate is exposed."},
			{Num: "04", Title: "Rinse and inspection", Desc: "The surface is rinsed clean and inspected for residual contamination. Oil-saturated areas may require a second treatment. We confirm the surface is ready for whatever comes next — sealcoating, striping, or just a clean finished appearance."},
		},
		WhyUs: []WhyPoint{
			{Title: "Prep that matters", Desc: "We pressure wash before sealcoating because it makes the sealcoat last longer. Not everyone does."},
			{Title: "Commercial equipment", Desc: "Truck-mounted pressure washers, not consumer units. The difference in cleaning power is significant."},
		},
		GoodToKnow: []string{
			"Oil contamination in asphalt doesn't just look bad — it breaks down the binder. Petroleum from vehicle leaks softens the asphalt surface and accelerates raveling. Heavy oil saturation may require that section to be removed and patched rather than just cleaned.",
			"The Puget Sound's mild, wet climate is ideal for moss and algae growth on asphalt and concrete. Moss holds moisture against the surface, accelerating deterioration. Regular pressure washing — typically annually — prevents accumulation from becoming a structural issue.",
			"Pressure washing alone doesn't seal, waterproof, or protect your asphalt. It prepares the surface for the work that does. If you're washing to prep for sealcoating, plan to seal within a few days after washing.",
		},
		ServiceArea:     "We provide pressure washing for asphalt and concrete surfaces throughout the Puget Sound — Puyallup, Sumner, Tacoma, Auburn, Federal Way, SeaTac, Covington, Maple Valley, Bellevue, and the greater Seattle area.",
		RelatedServices: []RelatedLink{
			{Name: "Sealcoating", Slug: "sealcoating", Desc: "The natural next step after a thorough surface cleaning"},
			{Name: "Crack Repair", Slug: "crack-filling", Desc: "Address cracks after washing reveals their full extent"},
			{Name: "Parking Lot Striping", Slug: "striping", Desc: "Re-stripe after washing exposes the old line condition"},
		},
		CustomCTAHead: "Ready for a clean surface?",
		CustomCTABody: "Whether you're prepping for sealcoating or just need a lot cleaned up, call us. No obligation.",
		ContactType:   "Other",
		PhotoDesc:     "Pressure washer cleaning asphalt surface — visible clean/dirty line where the spray has passed. Commercial lot or driveway.",
	},

	// ── 03 · SITE WORK ──
	{
		Slug:        "site-prep",
		Name:        "Site Prep, Grading & Demolition",
		Tagline:     "The paving job starts long before the asphalt arrives.",
		Category:    "Site Work",
		CategoryNum: "03",
		MetaDesc:    "Complete site preparation for paving projects — demolition, excavation, grading, and base work throughout the Puget Sound. A-Team handles the full job from dirt to pavement. Free estimates.",
		HeroEyebrow: "Site Prep, Grading & Demolition · Puget Sound, WA",
		Pitch: []string{
			"Most pavement failures aren't asphalt failures. They're base failures. The asphalt is fine — it's what's underneath that gave out. The subgrade wasn't properly excavated. The base rock was the wrong material or the wrong thickness. The grade didn't move water away from the surface. Somebody skipped a step to save time or money, and the pavement above it paid for it.",
			"We handle site work because we've seen what happens when someone else does it wrong and we have to pave over it. We'd rather control the whole job — from the first excavator bucket to the final load of asphalt — so that what we build holds.",
		},
		Includes: []string{
			"Existing asphalt demolition (saw cutting and full tearout)",
			"Existing concrete removal",
			"Material haul-off and recycling",
			"Excavation to design depth",
			"Subgrade assessment and treatment (soil stabilization where needed)",
			"Imported base rock installation (specified by load and soil type)",
			"Base compaction (lift-by-lift, not single pass)",
			"Drainage grading — slope and cross-slope for runoff management",
			"Rough and finish grade for paving",
			"Sod removal and site clearing",
		},
		Process: []ProcessStep{
			{Num: "01", Title: "Site assessment and grade design", Desc: "Before any excavation begins, we evaluate existing drainage patterns, soil conditions, and the intended use of the finished surface. Passenger vehicles vs. heavy truck traffic requires different base specs. We plan the grade to move water away from the pavement and away from adjacent structures."},
			{Num: "02", Title: "Demo and excavation", Desc: "Existing asphalt and concrete come out to the required depth. We saw-cut clean edges where the work meets existing pavement that's staying in place. Demo material is loaded and hauled off — asphalt is 100% recyclable."},
			{Num: "03", Title: "Subgrade evaluation", Desc: "Once the surface material is removed, we can see and probe the subgrade. Soft spots, organic material, or unstable soil sections need to be addressed now — not after base rock is placed on top of them."},
			{Num: "04", Title: "Base rock installation and compaction", Desc: "We specify base depth based on the intended load. A standard residential driveway carries very different loads than a commercial lot with delivery trucks. We compact in lifts — each lift gets compacted before the next goes down. Cutting this step is the single most common cause of premature pavement failure."},
			{Num: "05", Title: "Final grade", Desc: "The finished base grade is checked for slope and drainage before paving begins. Water standing on a newly paved surface after a rain event is a sign the grade wasn't set correctly. We check it before it's covered up."},
		},
		WhyUs: []WhyPoint{
			{Title: "One contractor, whole job", Desc: "We do demo, base, and paving. No waiting on three different companies."},
			{Title: "Clean removal", Desc: "Sawcut edges, proper hauling, site left clean. Not a mess for you to deal with."},
			{Title: "Right equipment", Desc: "Skid-steers, excavators, dump trucks — we bring what the job needs."},
		},
		GoodToKnow: []string{
			"Puget Sound soils vary significantly by location. Areas near water features, filled land, or organic-rich soils — common in parts of the Auburn valley, the Kent lowlands, and near-water properties in Tacoma — can have subgrade conditions that affect base depth and drainage design.",
			"If you're paving over an existing gravel surface, the existing gravel may or may not be usable as base material depending on gradation and compaction history. We assess it before deciding. Sometimes you save on imported base; sometimes the existing gravel needs to come out.",
			"Drainage design matters as much as base depth for pavement longevity. A parking lot with inadequate grade will hold standing water, which accelerates asphalt deterioration and base saturation. Getting the drainage right during grade work is not optional.",
		},
		ServiceArea:     "We provide site prep, grading, and demolition for paving projects throughout the Puget Sound — Tacoma, Puyallup, Auburn, Sumner, Federal Way, Covington, Maple Valley, SeaTac, Bellevue, and the greater Seattle area.",
		RelatedServices: []RelatedLink{
			{Name: "Asphalt Paving", Slug: "asphalt-paving", Desc: "Once the site is prepped and ready"},
			{Name: "Parking Lot Paving", Slug: "parking-lot-paving", Desc: "Full commercial lot projects"},
			{Name: "Driveway Paving", Slug: "driveway-paving", Desc: "Residential installs from bare ground up"},
		},
		CustomCTAHead: "Starting from scratch or tearing out what's there?",
		CustomCTABody: "Call us. We handle the full scope from demolition through finish grade. No obligation.",
		ContactType:   "Grading / site work",
		PhotoDesc:     "Excavator removing old asphalt, broken pavement being loaded into a dump truck. Cleared site with clean edges visible.",
		PhotoURL:      "/static/img/site-work-800.webp",
	},
	{
		Slug:        "grading",
		Name:        "Grading & Excavation",
		Tagline:     "Proper drainage starts with proper grading.",
		Category:    "Site Work",
		CategoryNum: "03",
		MetaDesc:    "Professional grading and excavation for paving projects in the Puget Sound area. Proper slopes for drainage and stable pavement. Free estimates.",
		Pitch: []string{
			"Proper grading is why pavement drains and stays stable. We grade it right before we pave it.",
			"If water pools on your pavement, the grade is wrong. We establish the right slopes for drainage, compact the subgrade to prevent settling, and make sure the surface you're paving over is solid and uniform. This is where pavement longevity starts.",
		},
		Includes: []string{
			"Subgrade preparation and compaction",
			"Slope grading for proper drainage",
			"Excavation and soil removal",
			"Fine grading to specified tolerances",
			"Compaction testing",
		},
		WhyUs: []WhyPoint{
			{Title: "We grade before we pave", Desc: "Sounds obvious, but some crews skip it. We don't."},
			{Title: "Drainage-first thinking", Desc: "Every grade we set considers where water goes. Standing water destroys pavement."},
			{Title: "Compaction standards", Desc: "We compact subgrade to spec. Soft spots under pavement become potholes."},
		},
		ContactType: "Grading / site work",
		PhotoDesc:   "Skid-steer or grader working on exposed earth, establishing grade. Laser level or string lines visible. Prepared subgrade surface.",
		PhotoURL:    "/static/img/site-work-800.webp",
	},
	{
		Slug:        "base-work",
		Name:        "Rocking & Base Work",
		Tagline:     "Where most pavement failures start. We don't skip it.",
		Category:    "Site Work",
		CategoryNum: "03",
		MetaDesc:    "Crushed rock base preparation for asphalt paving in the Puget Sound area. Proper base work that prevents premature pavement failure.",
		Pitch: []string{
			"Base preparation is where most pavement failures start. We don't skip it.",
			"Crushed rock base is the foundation your asphalt sits on. Wrong depth, wrong material, or poor compaction and the pavement fails in years instead of decades. We install the right depth of properly graded crushed rock and compact it to spec before any asphalt goes down.",
		},
		Includes: []string{
			"Crushed rock delivery and placement",
			"Base depth to specification",
			"Mechanical compaction",
			"Subgrade stabilization when needed",
			"Geotextile fabric installation",
		},
		WhyUs: []WhyPoint{
			{Title: "We don't skip base work", Desc: "The most common shortcut in paving. We've never taken it."},
			{Title: "Right material, right depth", Desc: "We spec the base for your soil conditions and traffic load, not a one-size-fits-all."},
			{Title: "Compacted to spec", Desc: "We mechanically compact in lifts. Dumping rock and rolling once isn't compaction."},
		},
		ContactType: "Grading / site work",
		PhotoDesc:   "Crushed rock base being spread and compacted. Vibratory roller or plate compactor on gravel surface. Prepared base ready for paving.",
		PhotoURL:    "/static/img/site-work-800.webp",
	},

	// ── 04 · STRIPING, SIGNAGE & SAFETY ──
	{
		Slug:        "striping",
		Name:        "Striping, Bollards & Signage",
		Tagline:     "A well-marked lot is safer, more organized, and less of a liability.",
		Category:    "Striping, Signage & Safety",
		CategoryNum: "04",
		MetaDesc:    "Professional parking lot striping, bollard installation, and signage for commercial properties throughout the Puget Sound. ADA-compliant layouts, fire lane marking, and more. Free estimates.",
		HeroEyebrow: "Striping, Bollards & Signage · Puget Sound, WA",
		Pitch: []string{
			"Striping and signage are the last step in a paving project and the most visible result. They're also where compliance requirements are sharpest — improperly marked ADA stalls, missing fire lane markings, and inadequate signage create real liability exposure for property owners.",
			"We handle striping as part of larger paving and sealcoating projects, and as standalone work for lots that need a refresh. We also install bollards and parking signage — so the finish scope of a commercial lot project doesn't require a second contractor.",
		},
		Includes: []string{
			"Standard parking stall layout and striping",
			"ADA accessible stall striping (current federal and WA standards)",
			"Re-stripe of faded lots (surface prep included)",
			"Directional arrows, crosswalks, and fire lane marking",
			"Curb painting (red, yellow, blue)",
			"Steel pipe bollard installation (surface mount and embedded)",
			"Bollard replacement and repair",
			"ADA parking signage, stop signs, directional signs",
			"Custom property signage installation",
		},
		Process: []ProcessStep{
			{Num: "01", Title: "Layout planning", Desc: "For new lots or re-layouts, we plan stall count, traffic flow, and ADA accessible stall placement before a line goes down. ADA requirements govern not just the stall itself but the accessible route from the stall to the building entrance — we design for the full compliance picture."},
			{Num: "02", Title: "Surface preparation", Desc: "We clean the surface before striping. Paint applied over loose debris, dust, or sealcoat residue doesn't last. On re-stripe jobs, existing line location is marked before cleaning so new lines go in the right places."},
			{Num: "03", Title: "Layout and stripes", Desc: "Lines go down with a mechanical striper — straight, consistent width, sharp edges. We don't hand-paint. ADA stall dimensions, access aisle widths, and slope requirements are verified against current standards."},
			{Num: "04", Title: "Stencils, signage, and bollards", Desc: "Handicap symbols, arrows, curb markings, and any stenciled text go in after the stall lines are down. Signs and bollards are installed at the correct location and height per applicable codes."},
		},
		WhyUs: []WhyPoint{
			{Title: "ADA included, always", Desc: "Handicap stalls, access aisles, and signage are part of every lot we stripe. Not an add-on."},
			{Title: "Traffic-grade paint", Desc: "We use commercial traffic paint, not hardware store latex. It lasts through Washington weather."},
			{Title: "Full finish scope", Desc: "Striping, bollards, signage, curb paint — one crew handles the complete finish."},
		},
		GoodToKnow: []string{
			"ADA accessible parking stalls require a minimum 8-foot width with a 5-foot adjacent access aisle. Van-accessible stalls require either 11-foot width or an 8-foot stall with an 8-foot aisle. Slope within the stall and aisle must not exceed 2% in any direction.",
			"Paint on a freshly sealcoated surface needs time to adhere — typically 24-48 hours minimum. If your contractor is doing same-day sealing and striping, the paint won't last.",
			"Fire lane markings are governed by local fire codes, which vary by municipality. Tacoma, Auburn, Federal Way, and Puyallup each have specific requirements. We stripe to the local standard.",
			"Bollards protect against low-speed vehicle impacts — storefronts, pedestrian zones, fuel islands, loading docks. If your property has areas where accidental vehicle incursion would create serious liability, bollards are worth a conversation.",
		},
		ServiceArea:     "We provide parking lot striping, bollard installation, and signage throughout the Puget Sound — Tacoma, Puyallup, Auburn, Federal Way, Sumner, SeaTac, Covington, Maple Valley, Bellevue, and the greater Seattle area.",
		RelatedServices: []RelatedLink{
			{Name: "Parking Lot Paving", Slug: "parking-lot-paving", Desc: "Full lot construction or replacement"},
			{Name: "Sealcoating", Slug: "sealcoating", Desc: "Protect the surface before re-striping"},
			{Name: "Concrete Installation", Slug: "concrete", Desc: "ADA ramps and sidewalk work alongside striping"},
		},
		CustomCTAHead: "Need your lot striped, marked, or protected?",
		CustomCTABody: "Call us or fill out the form. We'll come out, assess the lot, and give you a straight quote. Free estimate.",
		ContactType:   "Striping & marking",
		PhotoDesc:     "Freshly striped parking lot — crisp white or yellow lines on dark asphalt, ADA-blue handicap stalls visible.",
		PhotoURL:      "/static/img/striped-lot-800.webp",
	},
	{
		Slug:        "bollards",
		Name:        "Bollard Installation & Repair",
		Tagline:     "Protect storefronts and commercial property.",
		Category:    "Striping, Signage & Safety",
		CategoryNum: "04",
		MetaDesc:    "Bollard installation and repair for commercial properties in the Puget Sound area. Protective posts for storefronts, parking structures, and property lines.",
		Pitch: []string{
			"Protective bollards for storefronts, parking structures, and commercial property. Install, repair, or replace.",
			"Bollards protect buildings, pedestrians, and infrastructure from vehicle impact. We install new bollard systems, repair damaged posts, and replace bollards that have done their job. Concrete-set or surface-mount, painted or bare — matched to your property's needs.",
		},
		Includes: []string{
			"New bollard installation",
			"Damaged bollard repair and replacement",
			"Concrete-set and surface-mount options",
			"Painting and visibility marking",
			"ADA-compliant placement",
		},
		WhyUs: []WhyPoint{
			{Title: "Properly anchored", Desc: "We set bollards in concrete to spec. A bollard that folds on impact isn't protecting anything."},
			{Title: "Paired with paving", Desc: "Installing bollards during a paving project is cheaper and cleaner than retrofitting later."},
		},
		ContactType: "Striping & marking",
		PhotoDesc:   "Yellow-painted steel bollards in front of a commercial storefront or parking structure. Clean installation, fresh paint.",
		PhotoURL:    "/static/img/striped-lot-800.webp",
	},
	{
		Slug:        "signage",
		Name:        "Signage",
		Tagline:     "Parking, directional, and safety signage — installed right.",
		Category:    "Striping, Signage & Safety",
		CategoryNum: "04",
		MetaDesc:    "Parking lot signage installation in the Puget Sound area. Directional, safety, and regulatory signs for commercial properties.",
		Pitch: []string{
			"Parking, directional, and safety signage installed alongside your paving project or standalone.",
			"Signs get hit, fade, and fall over. We install new signage and replace what's damaged. Handicap signs, stop signs, directional signs, speed limit signs — everything a commercial lot needs to be safe and compliant.",
		},
		Includes: []string{
			"Handicap parking signs (ADA-compliant)",
			"Stop signs and yield signs",
			"Directional and wayfinding signs",
			"Speed limit and traffic control signs",
			"Post installation and replacement",
		},
		WhyUs: []WhyPoint{
			{Title: "Compliance-aware", Desc: "We know which signs are required and where they need to be. You won't get cited."},
			{Title: "Bundled with paving", Desc: "Adding signage to a paving or striping project saves a separate mobilization cost."},
		},
		ContactType: "Striping & marking",
		PhotoDesc:   "Handicap parking sign properly installed on a post in a freshly paved and striped commercial lot.",
		PhotoURL:    "/static/img/striped-lot-800.webp",
	},

	// ── 05 · YEAR-ROUND ──
	{
		Slug:        "snow-removal",
		Name:        "Snow Removal & De-Icing",
		Tagline:     "When it snows in the Puget Sound, it shuts things down. We show up.",
		Category:    "Year-Round Services",
		CategoryNum: "05",
		MetaDesc:    "Commercial snow removal and de-icing for parking lots and drives throughout the Puget Sound. Seasonal contracts available. Don't get caught unprepared — serving Tacoma, Puyallup, Auburn, and surrounding areas.",
		HeroEyebrow: "Snow Removal & De-Icing · Puget Sound, WA",
		Pitch: []string{
			"Western Washington doesn't get the relentless winters of eastern Washington or the mountain passes. But the Puget Sound gets enough cold weather to create serious problems for commercial properties — and the region isn't built to handle it the way colder climates are.",
			"What the Puget Sound gets most is ice. Temperatures hover at or near freezing, precipitation comes down as rain, then sleet, then freezes overnight on pavement that drained during the day. A thin layer of clear ice on an unprotected parking lot is nearly invisible and genuinely dangerous.",
			"When that happens, the property owners who have service agreements in place get plowed. The ones who start calling around in the morning don't.",
		},
		Includes: []string{
			"Parking lot snow plowing and clearing",
			"Access drive and entrance clearing",
			"Snow haul-off for large accumulation events",
			"Granular de-icing application (rock salt and blended products)",
			"Liquid anti-icing application (pre-treatment before forecast events)",
			"Walkway and sidewalk clearing",
			"Ongoing seasonal service contracts (priority response)",
			"Single-event response (as availability permits)",
		},
		Process: []ProcessStep{
			{Num: "01", Title: "Pre-season contract setup", Desc: "The most reliable winter coverage is a seasonal contract established before the weather arrives. We assess your property, understand the critical areas — main lot access, fire lane clearance, accessible parking routes — and set a response protocol."},
			{Num: "02", Title: "Weather monitoring", Desc: "We watch forecasts, not just actuals. Liquid anti-icing applied before precipitation freezes is significantly more effective — and uses less material — than de-icing applied after ice has already formed."},
			{Num: "03", Title: "Snow event response", Desc: "When snow accumulates to trigger threshold, we dispatch. Contract clients have priority over call-in requests. For significant events, we work in phases — initial clearing to open traffic lanes, followed by cleanup and edge work."},
			{Num: "04", Title: "De-icing and sand application", Desc: "After plowing, we apply de-icing material to high-traffic areas and pedestrian zones. Material selection is matched to conditions — granular products for standard freeze events, blended products for persistent ice."},
			{Num: "05", Title: "Post-event check", Desc: "Freeze-thaw cycles after a snow event can create ice on surfaces that were clear the night before. We monitor conditions and return for re-application where contracts include ongoing coverage."},
		},
		WhyUs: []WhyPoint{
			{Title: "We show up", Desc: "When it snows at 3 AM, we're already moving. Your lot is clear by open."},
			{Title: "Contract or per-event", Desc: "Seasonal contracts for guaranteed service, or call us when you need us."},
			{Title: "Same crew you trust", Desc: "The company that paves your lot clears your snow. One relationship, year-round."},
		},
		GoodToKnow: []string{
			"Puget Sound snow events are often short-duration but high-impact. Three inches on a Tuesday morning in Tacoma creates more disruption than a foot of snow in a climate built to handle it. The liability exposure is real.",
			"Black ice is the most dangerous condition we deal with in this region. It forms when rain or melting snow refreezes on pavement overnight — particularly on north-facing surfaces and shaded lots. Pre-treatment before forecast freeze events is the most effective mitigation.",
			"Snow removal contracts fill up fast when an event is actually in the forecast. If you need reliable winter coverage, the time to set up a contract is late summer or early fall.",
			"De-icing materials can damage vegetation, concrete, and metal when over-applied. We use the appropriate products at the appropriate rates. The goal is a safe surface, not a white parking lot.",
		},
		ServiceArea:     "We provide commercial snow removal and de-icing throughout the Puget Sound — Tacoma, Puyallup, Auburn, Federal Way, Sumner, Covington, Maple Valley, SeaTac, and the surrounding region.",
		RelatedServices: []RelatedLink{
			{Name: "Parking Lot Paving", Slug: "parking-lot-paving", Desc: "Ensure your lot is properly graded for drainage before winter"},
			{Name: "Sealcoating", Slug: "sealcoating", Desc: "Protect your pavement from freeze-thaw damage"},
			{Name: "Striping & Bollards", Slug: "striping", Desc: "Clear accessible routes and lot markings before winter"},
		},
		CustomCTAHead: "Don't wait until it snows to think about this.",
		CustomCTABody: "Seasonal contracts fill early. Call us now to get your property assessed and your winter coverage set up before the season arrives.",
		ContactType:   "Snow plowing",
		PhotoDesc:     "Plow truck clearing a commercial parking lot in early morning light. Heavy wet PNW snow, headlights on, cleared lanes visible.",
		PhotoURL:      "/static/img/snow-removal-800.webp",
	},
}

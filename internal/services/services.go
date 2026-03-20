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
		Tagline:     "Slabs, curbing, sidewalks, and flatwork — residential and commercial.",
		Category:    "Asphalt & Concrete",
		CategoryNum: "01",
		MetaDesc:    "Concrete installation and repair in the Puget Sound area. Slabs, curbing, sidewalks, and flatwork to municipal standards. Free estimates: (253) 365-1283.",
		Pitch: []string{
			"Slabs, curbing, sidewalks, and flatwork. Residential and commercial. Done to municipal standards when required.",
			"We pour concrete the same way we lay asphalt — with proper prep, proper forms, and no shortcuts on the mix or the cure time. The result is concrete that doesn't crack in two winters.",
		},
		Includes: []string{
			"Concrete slabs and pads",
			"Curbing and gutter work",
			"Sidewalks and walkways",
			"Flatwork and aprons",
			"Concrete removal and replacement",
			"ADA-compliant installations",
		},
		WhyUs: []WhyPoint{
			{Title: "Municipal-grade work", Desc: "We build to code. When your project needs permits and inspection, it passes."},
			{Title: "Proper cure time", Desc: "We don't rush the finish. Concrete needs time to reach full strength."},
			{Title: "One contractor for everything", Desc: "Asphalt and concrete on the same project? One crew, one invoice."},
		},
		ContactType: "Concrete / pavers / hardscape",
		PhotoDesc:   "Fresh concrete being poured or finished — smooth wet surface with forms visible. Worker with float or trowel. Residential or commercial setting.",
		PhotoURL:    "/static/img/concrete-pour-800.webp",
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
		Tagline:     "Two coats standard — because one isn't enough for a Washington winter.",
		Category:    "Maintenance & Protection",
		CategoryNum: "02",
		MetaDesc:    "Professional sealcoating in the Puget Sound area. Two-coat application that protects asphalt from UV, water, and freeze-thaw damage. Free estimates: (253) 365-1283.",
		Pitch: []string{
			"Asphalt needs protection from UV, water infiltration, and freeze-thaw cycling. We apply two coats because one isn't sufficient for Washington conditions. Done right, sealcoating extends pavement life significantly.",
			"Sealcoat is your pavement's first line of defense. It blocks UV damage that makes asphalt brittle, prevents water from reaching the base, and gives the surface a clean, uniform appearance. We prep the surface properly and apply two full coats — because that's what it takes to last through our climate.",
		},
		Includes: []string{
			"Surface cleaning and preparation",
			"Crack filling prior to sealcoat",
			"Two-coat sealant application",
			"Oil spot priming",
			"Barricading and dry-time management",
			"Residential and commercial properties",
		},
		WhyUs: []WhyPoint{
			{Title: "Two coats, always", Desc: "Single-coat sealcoating is a shortcut. We apply two because that's what Washington weather demands."},
			{Title: "Proper prep", Desc: "We clean the surface and fill cracks first. Sealcoat over dirt and open cracks is money wasted."},
			{Title: "Right conditions", Desc: "We don't sealcoat in the rain or below 50°F. If the weather's wrong, we reschedule."},
		},
		ContactType: "Sealcoating / repair",
		PhotoDesc:   "Worker applying sealcoat with squeegee or spray rig. Wet glossy black surface contrasting with unsealed faded gray asphalt. Driveway or small lot.",
		PhotoURL:    "/static/img/fresh-asphalt-800.webp",
	},
	{
		Slug:        "crack-filling",
		Name:        "Crack Filling",
		Tagline:     "Fix cracks now or fix pavement later. Your call.",
		Category:    "Maintenance & Protection",
		CategoryNum: "02",
		MetaDesc:    "Professional crack filling for asphalt in the Puget Sound area. Prevent water infiltration and structural failure. Free estimates: (253) 365-1283.",
		Pitch: []string{
			"Get ahead of it. Cracks that get water in them fail fast. We clean and fill before they become structural problems.",
			"Water is asphalt's worst enemy. When it gets into cracks and reaches the base, freeze-thaw cycles break the pavement apart from underneath. Crack filling is the cheapest maintenance you can do — and the one that prevents the most expensive repairs.",
		},
		Includes: []string{
			"Crack routing and cleaning",
			"Hot-pour rubberized crack sealant",
			"Linear and alligator crack treatment",
			"Joint sealing",
			"Assessment of underlying base condition",
		},
		WhyUs: []WhyPoint{
			{Title: "Preventive, not cosmetic", Desc: "We fill cracks to stop water infiltration, not just to make the surface look better."},
			{Title: "Hot-pour sealant", Desc: "We use rubberized hot-pour material that flexes with temperature changes — not cold-pour from a jug."},
			{Title: "Honest about limits", Desc: "If cracks indicate base failure, we'll tell you. Filling over a failed base doesn't help."},
		},
		ContactType: "Sealcoating / repair",
		PhotoDesc:   "Close-up of crack filling in progress — hot-pour sealant being applied to clean crack in asphalt surface. Applicator wand or pour pot visible.",
	},
	{
		Slug:        "pressure-washing",
		Name:        "Pressure Washing",
		Tagline:     "Clean pavement holds sealcoat better and looks better.",
		Category:    "Maintenance & Protection",
		CategoryNum: "02",
		MetaDesc:    "Commercial and residential pressure washing for asphalt and concrete surfaces in the Puget Sound area. Surface prep and standalone cleaning.",
		Pitch: []string{
			"Surface prep before sealing, or standalone. Clean pavement holds sealcoat better and looks better.",
			"Oil stains, moss, dirt, and debris prevent sealcoat from bonding properly. Pressure washing removes all of it and gives you a clean surface that holds treatment. We also do standalone washing for property managers who want their lots and sidewalks looking sharp.",
		},
		Includes: []string{
			"High-pressure surface cleaning",
			"Oil and grease stain treatment",
			"Moss and algae removal",
			"Pre-sealcoat surface preparation",
			"Parking lots, sidewalks, and driveways",
		},
		WhyUs: []WhyPoint{
			{Title: "Prep that matters", Desc: "We pressure wash before sealcoating because it makes the sealcoat last longer. Not everyone does."},
			{Title: "Commercial equipment", Desc: "Truck-mounted pressure washers, not consumer units. The difference in cleaning power is significant."},
		},
		ContactType: "Other",
		PhotoDesc:   "Pressure washer cleaning asphalt surface — visible clean/dirty line where the spray has passed. Commercial lot or driveway.",
	},

	// ── 03 · SITE WORK ──
	{
		Slug:        "site-prep",
		Name:        "Site Prep & Demolition",
		Tagline:     "Job site ready before the first load of material arrives.",
		Category:    "Site Work",
		CategoryNum: "03",
		MetaDesc:    "Site preparation and demolition for paving projects in the Puget Sound area. Asphalt and concrete removal, clearing, and grading. Free estimates.",
		Pitch: []string{
			"Existing asphalt or concrete removal, demo, and clearing. We handle it so the job site is ready before the first load of material arrives.",
			"Every paving job starts with what's already there. We remove old pavement, break out failed concrete, clear debris, and leave you with a clean surface ready for base work. One crew handles demo through final paving — no coordinating between contractors.",
		},
		Includes: []string{
			"Asphalt removal and disposal",
			"Concrete demolition and hauling",
			"Site clearing and debris removal",
			"Sawcutting for clean edges",
			"Coordination with utility locates",
		},
		WhyUs: []WhyPoint{
			{Title: "One contractor, whole job", Desc: "We do demo, base, and paving. No waiting on three different companies."},
			{Title: "Clean removal", Desc: "Sawcut edges, proper hauling, site left clean. Not a mess for you to deal with."},
			{Title: "Right equipment", Desc: "Skid-steers, excavators, dump trucks — we bring what the job needs."},
		},
		ContactType: "Grading / site work",
		PhotoDesc:   "Excavator removing old asphalt, broken pavement being loaded into a dump truck. Cleared site with clean edges visible.",
		PhotoURL:    "/static/img/site-work-800.webp",
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
		Name:        "Parking Lot Striping",
		Tagline:     "Clean lines, ADA compliance, professional appearance.",
		Category:    "Striping, Signage & Safety",
		CategoryNum: "04",
		MetaDesc:    "Parking lot striping and line marking in the Puget Sound area. New layouts, re-stripes, and ADA-compliant markings. Free estimates.",
		Pitch: []string{
			"Clean, straight, durable lines. New layouts or re-stripe of existing lots. ADA stalls included.",
			"A freshly striped lot is the easiest way to make commercial property look maintained and professional. We do new layouts from scratch and re-stripes of existing markings. Every job includes proper ADA-compliant handicap stalls and access aisles.",
		},
		Includes: []string{
			"New parking lot layout and design",
			"Re-striping of existing markings",
			"ADA-compliant handicap stalls and signage",
			"Fire lane markings",
			"Directional arrows and crosswalks",
			"Durable traffic-grade paint",
		},
		WhyUs: []WhyPoint{
			{Title: "ADA included, always", Desc: "Handicap stalls, access aisles, and signage are part of every lot we stripe. Not an add-on."},
			{Title: "Traffic-grade paint", Desc: "We use commercial traffic paint, not hardware store latex. It lasts through Washington weather."},
			{Title: "Clean geometry", Desc: "Straight lines, proper spacing, consistent width. It's detail work and we take it seriously."},
		},
		ContactType: "Striping & marking",
		PhotoDesc:   "Freshly striped parking lot — crisp white or yellow lines on dark asphalt, ADA-blue handicap stalls visible. Clean geometric composition.",
		PhotoURL:    "/static/img/striped-lot-800.webp",
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
		Tagline:     "When it snows, we show up.",
		Category:    "Year-Round Services",
		CategoryNum: "05",
		MetaDesc:    "Commercial snow removal and de-icing services in the Puget Sound area. Parking lot clearing, salt application, and ongoing service contracts.",
		Pitch: []string{
			"When it snows, we show up. Parking lot clearing, de-icing applications, and ongoing service contracts available for commercial properties.",
			"PNW snow is wet, heavy, and unpredictable. We run plows and salt trucks for commercial properties on a contract or per-event basis. When the forecast calls for it, we're staged and ready. Your lot is clear before your tenants arrive.",
		},
		Includes: []string{
			"Parking lot snow plowing",
			"Sidewalk and walkway clearing",
			"De-icing salt and chemical application",
			"Seasonal service contracts",
			"Per-event service available",
			"Early-morning priority clearing",
		},
		WhyUs: []WhyPoint{
			{Title: "We show up", Desc: "When it snows at 3 AM, we're already moving. Your lot is clear by open."},
			{Title: "Contract or per-event", Desc: "Seasonal contracts for guaranteed service, or call us when you need us."},
			{Title: "Same crew you trust", Desc: "The company that paves your lot clears your snow. One relationship, year-round."},
		},
		ContactType: "Snow plowing",
		PhotoDesc:   "Plow truck clearing a commercial parking lot in early morning light. Heavy wet PNW snow, headlights on, cleared lanes visible.",
		PhotoURL:    "/static/img/snow-removal-800.webp",
	},
}

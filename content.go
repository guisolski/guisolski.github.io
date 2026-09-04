package main

// TimelineEntry is one milestone on the career rail.
//
// Time is the machine-readable form of Date, emitted as <time datetime="...">
// so a crawler or a language model reads the same milestone a person does.
type TimelineEntry struct {
	Year  string
	Date  string
	Time  string
	Body  string
	Quote string // recognition text; a non-empty quote marks the entry as an honor
}

var timelineEntries = []TimelineEntry{
	{
		Year: "2015",
		Date: "February 2015",
		Time: "2015-02",
		Body: "Started a technical high-school program at Instituto Federal do Paraná (IFPR), specializing in digital game development.",
	},
	{
		Year: "2017",
		Date: "December 2017",
		Time: "2017-12",
		Body: "Graduated from the technical program. For the capstone I built a physiotherapy game using Microsoft Kinect that tracks a patient's movements and compares them against exercises recorded by the physiotherapist.",
	},
	{
		Year: "2018",
		Date: "February 2018",
		Time: "2018-02",
		Body: "Began a BSc in Computer Science at Pontifícia Universidade Católica do Paraná (PUCPR).",
	},
	{
		Year: "2019",
		Date: "January 2019",
		Time: "2019-01",
		Body: "Joined Banco do Brasil as a software engineering intern, building internal applications for a national retail bank using MVC with Java, JavaScript and SQL.",
	},
	{
		Year: "2020",
		Date: "March 2020",
		Time: "2020-03",
		Body: "Finished at Banco do Brasil and joined SaasTec Labs, developing and maintaining the company's PHP e-commerce platform and its backend processes.",
	},
	{
		Year: "2020",
		Date: "October 2020",
		Time: "2020-10",
		Body: "Became a PHP developer at SaasTec Labs and joined Go44 part-time as a Ruby developer, maintaining core systems and building a new module for the flagship product in support of Industry 4.0 initiatives.",
	},
	{
		Year: "2021",
		Date: "May 2021",
		Time: "2021-05",
		Body: "Joined Composable Bit as a self-employed Java developer, orchestrating microservices on Azure Kubernetes Service with Azure Functions for event-driven tasks.",
	},
	{
		Year: "2021",
		Date: "July 2021",
		Time: "2021-07",
		Body: "Joined ExxonMobil as a Java developer, building Kafka-based event integrations and synchronous Java APIs — configuring Confluent Cloud with an Avro schema registry so schemas evolve in a coordinated way across producers and consumers — and provisioning Azure infrastructure as code with Terraform.",
	},
	{
		Year: "2021",
		Date: "December 2021",
		Time: "2021-12",
		Body: "Earned my BSc in Computer Science.",
	},
	{
		Year: "2022",
		Date: "April 2022",
		Time: "2022-04",
		Body: "Started an MBA in Data Science & Analytics at USP / ESALQ.",
	},
	{
		Year: "2022",
		Date: "May 2022",
		Time: "2022-05",
		Body: "Promoted to regular Java developer at ExxonMobil and received a Recognition Award for contributions to the CSSP project.",
		Quote: "Your dedication to adapt, come together, and pivot helped bring the project to " +
			"its current stage. I appreciate your flexibility and collaboration in the early " +
			"project stages as a member of the Integration team. Thank you for your leadership!",
	},
	{
		Year: "2022",
		Date: "September 2022",
		Time: "2022-09",
		Body: "Received a Leadership Recognition Award at ExxonMobil for driving the MSP transition programme end to end — planning, execution, and cross-team coordination with an external vendor.",
		Quote: "Your servant leadership in leading the overall MSP transition activities " +
			"including planning, organizing meetings, executing and addressing key gaps are very " +
			"well organized and structured. Your excellent collaboration you made within the " +
			"program team and HCL as one team one goal spirit are truly recognized and " +
			"appreciated. Your commitment and mindset to enable our MSP to support the " +
			"continuity of products are contributing to the successful outcome of the flawless " +
			"transition and seamless experiences to our customers.",
	},
	{
		Year: "2022",
		Date: "October 2022",
		Time: "2022-10",
		Body: "Joined Mercado Libre as a Go engineer. I architect and develop microservices in the pricing and shipping domain for platforms serving 1M+ customers, own observability end to end with OpenTelemetry, and re-architected price-table promotion to be genuinely idempotent — so a retried or partially applied promotion converges on the same result, removing a class of data-integrity failures that only surfaced at thousands of requests per minute.",
	},
	{
		Year: "2024",
		Date: "2024",
		Time: "2024",
		Body: "Completed the MBA in Data Science & Analytics.",
	},
	{
		Year: "2025",
		Date: "May 2025",
		Time: "2025-05",
		Body: "Co-founded FINK AI part-time alongside full-time employment. I built the Open Finance integration layer end to end and the internal developer platform around it, taking the product from an empty repository to 20,000+ registered users and 4.2M+ financial records on a ~3,000-file Go monorepo.",
	},
}

// The intro reads as a lead line plus a body paragraph; aboutText is the two
// joined, and is what llms.txt and the JSON-LD description are built from.
const (
	aboutLead = "Backend engineer with 7+ years on high-throughput production systems, " +
		"and co-founder of FINK AI."

	aboutBody = "At Mercado Libre I write Go daily on e-commerce services that hold their " +
		"latency budget at thousands of requests per minute for 1M+ customers, and I own " +
		"observability end to end — OpenTelemetry traces, metrics and logs feeding Grafana, " +
		"Datadog, New Relic and Kibana. At FINK AI I took the product from an empty repository " +
		"to 20,000+ registered users and 4.2M+ financial records, on a ~3,000-file Go monorepo " +
		"I architected myself. Before that, Java at ExxonMobil — Kafka on Confluent Cloud, " +
		"Azure with Terraform, and an MSP transition that earned a Leadership Recognition " +
		"Award. Remote since 2020, based in Curitiba. Open to full-time backend roles, remote " +
		"or relocation."
)

const aboutText = aboutLead + " " + aboutBody

// Identity, as stated once and reused by the page, llms.txt and the JSON-LD.
const (
	personName            = "Guilherme Solski Alves"
	personJobTitle        = "Backend Software Engineer"
	personEmail           = "guilhermesolskialves@gmail.com"
	personPhone           = "+55 41 99628-6624"
	personPhoneE164       = "+55-41-99628-6624"
	personLocality        = "Curitiba"
	personRegion          = "PR"
	personCountry         = "BR"
	personLocation        = "Curitiba, Paraná, Brazil"
	personTimezone        = "UTC−3"
	personTagline         = "Backend engineer writing Go at Mercado Libre. Part-time co-founder of FINK AI."
	siteURL               = "https://guisolski.github.io"
	profileImage          = "/assets/images/profile.jpg"
	availability          = "Open to full-time backend roles"
	availabilityQualifier = "Remote or relocation"
)

// Focus is one of the areas the "What I work on" section describes. Hue is the
// entry's stop on the dawn spectrum, in reading order.
type Focus struct {
	Title string
	Body  string
	Hue   int
}

var focusAreas = []Focus{
	{
		Title: "Go at scale",
		Hue:   245,
		Body: "Primary language since 2022. Goroutines and channels for concurrent I/O, layered " +
			"and semantic caching, table-driven tests. Latency budgets held at thousands of " +
			"requests per minute, and write paths redesigned to converge idempotently under retry.",
	},
	{
		Title: "Third-party API integration",
		Hue:   288,
		Body: "OAuth flows, webhook ingestion and replay, provider clients behind stable internal " +
			"ports — with the reliability layer real integrations need: rate limiting that honours " +
			"upstream Retry-After, jittered capped backoff, bounded retries, circuit breaking.",
	},
	{
		Title: "Developer tooling",
		Hue:   331,
		Body: "Custom Go static-analysis linters and a golangci-lint plugin, shipped as thin CLIs " +
			"wired into pre-commit and CI; independently versioned internal Go modules each " +
			"service pins on its own schedule.",
	},
	{
		Title: "Observability",
		Hue:   35,
		Body: "OpenTelemetry traces, metrics and logs feeding Grafana, Datadog, New Relic, Kibana " +
			"and Pyroscope, with dashboards and alerts built end to end rather than inherited.",
	},
}

type Link struct {
	Label string
	Href  string
	Icon  string
}

var socialLinks = []Link{
	{Label: "GitHub", Href: "https://github.com/guisolski", Icon: "github"},
	{Label: "LinkedIn", Href: "https://www.linkedin.com/in/guilherme-solski-alves/", Icon: "linkedin"},
	{Label: "HackerRank", Href: "https://www.hackerrank.com/guilhermesolski", Icon: "code"},
	{Label: "URI Online Judge", Href: "https://www.urionlinejudge.com.br/judge/en/profile/296338", Icon: "trophy"},
}

var contactLinks = []Link{
	{Label: personEmail, Href: "mailto:" + personEmail, Icon: "mail"},
	{Label: personPhone, Href: "tel:+5541996286624", Icon: "phone"},
	{Label: personLocation, Href: "", Icon: "pin"},
}

var courseLinks = []Link{
	{Label: "MBA, Data Science & Analytics — USP / ESALQ (2022–2024)", Href: "", Icon: "badge"},
	{Label: "BSc, Computer Science — PUCPR (2018–2021)", Href: "/assets/pdf/faculdade.pdf", Icon: "badge"},
	{Label: "Digital Game Developer — IFPR technical (2015–2017)", Href: "/assets/images/ifpr_certified.jpeg", Icon: "badge"},
}

var resumeLinks = []Link{
	{Label: "CV (English, PDF)", Href: "/assets/pdf/cv.pdf", Icon: "file"},
	{Label: "Currículo (Português, PDF)", Href: "/assets/pdf/Curriculum/portugues.pdf", Icon: "file"},
}

// resumeLink is kept for callers/tests that expect the English CV path.
var resumeLink = resumeLinks[0]

// StackGroup is one labelled row of the Stack card.
type StackGroup struct {
	Label string
	Items []string
}

var stack = []StackGroup{
	{"Languages", []string{"Go", "TypeScript", "Java", "Python", "SQL"}},
	{"Data", []string{"PostgreSQL", "pgvector", "Redis", "Kafka"}},
	{"Platform", []string{"Kubernetes", "Docker", "Terraform", "Azure", "GCP"}},
}

type Tagged struct {
	Label string
	Tag   string
}

var spokenLanguages = []Tagged{
	{Label: "Portuguese", Tag: "Native"},
	{Label: "English", Tag: "Full professional"},
	{Label: "Spanish", Tag: "Professional working"},
}

// Machine-readable facts that have no visible counterpart on the page but
// belong in the structured data.
var (
	employers = []string{"Mercado Libre", "FINK AI"}

	schools = []struct{ Type, Name string }{
		{"CollegeOrUniversity", "USP / ESALQ"},
		{"CollegeOrUniversity", "Pontifícia Universidade Católica do Paraná"},
		{"EducationalOrganization", "Instituto Federal do Paraná"},
	}

	awards = []string{
		"Leadership Recognition Award, ExxonMobil (2022)",
		"Recognition Award, CSSP Project, ExxonMobil (2022)",
		"Recognition Award, CSIT, ExxonMobil (2022)",
	}

	knowsAbout = []string{
		"Go", "TypeScript", "Java", "Python", "SQL",
		"PostgreSQL", "pgvector", "Redis", "Apache Kafka",
		"Kubernetes", "Docker", "Terraform", "Microsoft Azure",
		"OpenTelemetry", "Observability", "Distributed systems",
		"OAuth", "Webhooks", "Idempotency",
	}

	knowsLanguage = []struct{ Name, Code string }{
		{"Portuguese", "pt-BR"},
		{"English", "en"},
		{"Spanish", "es"},
	}
)

package main

type TimelineEntry struct {
	Year  string
	Date  string
	Body  string
	Quote string // recognition text; a non-empty quote marks the entry as an honor
}

var timelineEntries = []TimelineEntry{
	{
		Year: "2015",
		Date: "February 2015",
		Body: "Started a technical high-school program at Instituto Federal do Paraná (IFPR), specializing in digital game development.",
	},
	{
		Year: "2017",
		Date: "December 2017",
		Body: "Graduated from the technical program. For the capstone project, I built a physiotherapy game using Microsoft Kinect that tracks a patient's movements and compares them against exercises recorded by the physiotherapist.",
	},
	{
		Year: "2018",
		Date: "February 2018",
		Body: "Began a Bachelor's degree in Computer Science at Pontifícia Universidade Católica do Paraná (PUCPR).",
	},
	{
		Year: "2019",
		Date: "January 2019",
		Body: "Joined Banco do Brasil as an intern, developing and maintaining web systems with Java, JavaScript, HTML, CSS, and MySQL, using Bootstrap and Semantic UI on the front end.",
	},
	{
		Year: "2020",
		Date: "March 2020",
		Body: "Completed the internship at Banco do Brasil and joined SaaSTec Labs as an intern, developing and maintaining e-commerce systems.",
	},
	{
		Year: "2020",
		Date: "October 2020",
		Body: "Became a PHP developer at SaaSTec Labs and joined Go44 part-time as a Ruby developer, building systems for Industry 4.0.",
	},
	{
		Year: "2021",
		Date: "May 2021",
		Body: "Joined Composable bit as a self-employed Java developer.",
	},
	{
		Year: "2021",
		Date: "July 2021",
		Body: "Joined ExxonMobil as a Java developer.",
	},
	{
		Year: "2021",
		Date: "December 2021",
		Body: "Earned my Bachelor's degree in Computer Science.",
	},
	{
		Year: "2022",
		Date: "April 2022",
		Body: "Started an MBA in Data Science & Analytics.",
	},
	{
		Year: "2022",
		Date: "May 2022",
		Body: "Promoted to regular Java developer at ExxonMobil and received a Recognition Award for contributions to the CSSP project.",
		Quote: "Your dedication to adapt, come together, and pivot helped bring the project to " +
			"its current stage. I appreciate your flexibility and collaboration in the early " +
			"project stages as a member of the Integration team. Thank you for your leadership!",
	},
	{
		Year: "2022",
		Date: "September 2022",
		Body: "Received a Recognition Award at ExxonMobil for contributions to the CSIT project.",
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
		Body: "Joined Mercado Libre as a software engineer writing Go for e-commerce platforms " +
			"serving 1M+ customers, owning scalable services and the observability chain end to end.",
	},
	{
		Year: "2023",
		Date: "2023",
		Body: "Completed the MBA in Data Science & Analytics.",
	},
	{
		Year: "2025",
		Date: "May 2025",
		Body: "Co-founded FINK AI part-time alongside full-time employment — a personal-finance " +
			"product with 20,000+ registered users and 4.2M+ financial records.",
	},
}

const aboutText = "Backend engineer with 7+ years in production systems. At Mercado Libre I write Go " +
	"for e-commerce platforms serving 1M+ customers and own observability end to end — metrics, " +
	"OpenTelemetry, and dashboards and alerts across New Relic, Grafana, Datadog, and Kibana. " +
	"Previously at ExxonMobil: Azure with Terraform, Confluent Cloud Kafka + Avro, and an MSP " +
	"transition (Leadership Recognition Award, 2022). Part-time co-founder of FINK AI since 2025 — " +
	"personal finance with 20,000+ users and 4.2M+ financial records. Open to full-time backend " +
	"roles (remote or relocation)."

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
	{Label: "guilhermesolskialves@gmail.com", Href: "mailto:guilhermesolskialves@gmail.com", Icon: "mail"},
	{Label: "+55 41 99628-6624", Href: "tel:+5541996286624", Icon: "phone"},
}

var courseLinks = []Link{
	{Label: "Digital Game Developer — technical certificate", Href: "/assets/images/ifpr_certified.jpeg", Icon: "badge"},
	{Label: "Bachelor of Computer Science — diploma", Href: "/assets/pdf/faculdade.pdf", Icon: "badge"},
	{Label: "MBA, Data Science & Analytics — USP / ESALQ (2022–2024)", Href: "", Icon: "badge"},
}

var resumeLinks = []Link{
	{Label: "CV (English, PDF)", Href: "/assets/pdf/cv.pdf", Icon: "file"},
	{Label: "Currículo (Português, PDF)", Href: "/assets/pdf/Curriculum/portugues.pdf", Icon: "file"},
}

// resumeLink is kept for callers/tests that expect the English CV path.
var resumeLink = resumeLinks[0]

type Tagged struct {
	Label string
	Tag   string
}

var programmingLanguages = []string{"Go", "Java", "Python", "SQL", "JavaScript"}

var spokenLanguages = []Tagged{
	{Label: "Portuguese", Tag: "Native"},
	{Label: "English", Tag: "Full professional"},
	{Label: "Spanish", Tag: "Professional working"},
}

package main

type TimelineEntry struct {
	Year string
	Date string
	Body string
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
		Date: "February 2019",
		Body: "Joined Banco do Brasil as an intern, developing and maintaining web systems with Java, JavaScript, HTML, CSS, and MySQL, using Bootstrap and Semantic UI on the front end.",
	},
	{
		Year: "2020",
		Date: "March 2020",
		Body: "Completed the internship at Banco do Brasil and joined SaaStec, developing and maintaining e-commerce systems.",
	},
	{
		Year: "2021",
		Date: "July 2021",
		Body: "Joined ExxonMobil as a back-end developer.",
	},
	{
		Year: "2021",
		Date: "October 2021",
		Body: "Started working at Go44, building systems for Industry 4.0.",
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
		Date: "2022",
		Body: "Moved from ExxonMobil to Mercado Libre.",
	},
	{
		Year: "2023",
		Date: "2023",
		Body: "Completed the MBA in Data Science & Analytics.",
	},
}

const aboutText = "I am a software developer with experience in banking and financial services, " +
	"energy, and e-commerce, currently working at Mercado Libre. I have worked with a range of " +
	"technologies — mostly Java, Spring, and JavaScript — and I am comfortable with practices " +
	"such as unit testing and Agile methodologies. My interests include software development, " +
	"data engineering, the Internet of Things, and cloud computing."

type Link struct {
	Label string
	Href  string
	Icon  string
}

var socialLinks = []Link{
	{Label: "GitHub", Href: "https://github.com/guisolski", Icon: "github"},
	{Label: "LinkedIn", Href: "https://www.linkedin.com/in/guilherme-solski-alves-566262160/", Icon: "linkedin"},
	{Label: "HackerRank", Href: "https://www.hackerrank.com/guilhermesolski", Icon: "code"},
	{Label: "URI Online Judge", Href: "https://www.urionlinejudge.com.br/judge/en/profile/296338", Icon: "trophy"},
}

var contactLinks = []Link{
	{Label: "guilhermesolskialves@gmail.com", Href: "mailto:guilhermesolskialves@gmail.com", Icon: "mail"},
	{Label: "guigusolski@gmail.com", Href: "mailto:guigusolski@gmail.com", Icon: "mail"},
	{Label: "+55 41 99628-6624", Href: "tel:+5541996286624", Icon: "phone"},
}

var courseLinks = []Link{
	{Label: "Digital Game Developer — technical certificate", Href: "/assets/images/ifpr_certified.jpeg", Icon: "badge"},
	{Label: "Bachelor of Computer Science — diploma", Href: "/assets/pdf/faculdade.pdf", Icon: "badge"},
}

var resumeLink = Link{Label: "Résumé (Portuguese, PDF)", Href: "/assets/pdf/Curriculum/portugues.pdf", Icon: "file"}

type Tagged struct {
	Label string
	Tag   string
}

var programmingLanguages = []string{"Java", "Python", "JavaScript", "PHP", "Go"}

var spokenLanguages = []Tagged{
	{Label: "Portuguese", Tag: "Native"},
	{Label: "English", Tag: "Advanced"},
	{Label: "Spanish", Tag: "Intermediate"},
}

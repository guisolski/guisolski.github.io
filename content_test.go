package main

import (
	"strings"
	"testing"
	"time"
)

func entryTime(t *testing.T, date string) time.Time {
	t.Helper()
	if parsed, err := time.Parse("January 2006", date); err == nil {
		return parsed
	}
	parsed, err := time.Parse("2006", date)
	if err != nil {
		t.Fatalf("unparseable date %q", date)
	}
	return parsed.AddDate(0, 11, 0)
}

func TestTimelineEntriesComplete(t *testing.T) {
	for _, entry := range timelineEntries {
		t.Run(entry.Date, func(t *testing.T) {
			if entry.Year == "" || entry.Date == "" || entry.Body == "" {
				t.Fatalf("entry has empty fields: %+v", entry)
			}
			if !strings.Contains(entry.Date, entry.Year) {
				t.Errorf("rail label %q does not match date %q", entry.Year, entry.Date)
			}
		})
	}
}

func TestTimelineEntriesChronological(t *testing.T) {
	for i := 1; i < len(timelineEntries); i++ {
		prev, cur := timelineEntries[i-1], timelineEntries[i]
		if entryTime(t, cur.Date).Before(entryTime(t, prev.Date)) {
			t.Errorf("entry %q comes before %q but is dated earlier", cur.Date, prev.Date)
		}
	}
}

func TestTimelineCareerFacts(t *testing.T) {
	all := make([]string, 0, len(timelineEntries))
	for _, entry := range timelineEntries {
		all = append(all, entry.Body)
	}
	joined := strings.Join(all, " ")

	for _, fact := range []string{"Mercado Libre", "ExxonMobil", "MBA in Data Science & Analytics", "Recognition Award", "FINK AI"} {
		if !strings.Contains(joined, fact) {
			t.Errorf("timeline is missing %q", fact)
		}
	}
	if !strings.Contains(aboutText, "Mercado Libre") {
		t.Errorf("about text should mention the current employer, Mercado Libre")
	}
}

func TestAboutTextPositioning(t *testing.T) {
	for _, fact := range []string{
		"Backend engineer",
		"7+ years",
		"Mercado Libre",
		"Go",
		"1M+ customers",
		"OpenTelemetry",
		"ExxonMobil",
		"FINK AI",
		"20,000+",
		"4.2M+",
	} {
		if !strings.Contains(aboutText, fact) {
			t.Errorf("about text is missing %q", fact)
		}
	}
	for _, old := range []string{"Internet of Things", "data engineering", "Agile methodologies"} {
		if strings.Contains(aboutText, old) {
			t.Errorf("about text still contains outdated interest phrasing %q", old)
		}
	}
}

func TestProgrammingLanguagesOrder(t *testing.T) {
	want := []string{"Go", "Java", "Python", "SQL", "JavaScript"}
	if len(programmingLanguages) != len(want) {
		t.Fatalf("programmingLanguages len = %d, want %d (%v)", len(programmingLanguages), len(want), programmingLanguages)
	}
	for i, lang := range want {
		if programmingLanguages[i] != lang {
			t.Errorf("programmingLanguages[%d] = %q, want %q", i, programmingLanguages[i], lang)
		}
	}
}

func TestSpokenLanguageLevels(t *testing.T) {
	byLabel := map[string]string{}
	for _, l := range spokenLanguages {
		byLabel[l.Label] = l.Tag
	}
	if byLabel["English"] == "Native" {
		t.Errorf("English must not claim Native proficiency")
	}
	if got := byLabel["English"]; got != "Full professional" && got != "Advanced" {
		t.Errorf("English tag = %q, want Full professional or Advanced", got)
	}
	if got := byLabel["Spanish"]; got != "Professional working" && got != "Intermediate" {
		t.Errorf("Spanish tag = %q, want Professional working or Intermediate", got)
	}
	if byLabel["Portuguese"] != "Native" {
		t.Errorf("Portuguese tag = %q, want Native", byLabel["Portuguese"])
	}
}

func TestLinkedInURL(t *testing.T) {
	const want = "https://www.linkedin.com/in/guilherme-solski-alves/"
	for _, link := range socialLinks {
		if link.Label != "LinkedIn" {
			continue
		}
		if link.Href != want {
			t.Errorf("LinkedIn href = %q, want %q", link.Href, want)
		}
		if strings.Contains(link.Href, "566262160") {
			t.Errorf("LinkedIn href still has the old numeric suffix: %q", link.Href)
		}
		return
	}
	t.Fatal("LinkedIn link not found in socialLinks")
}

func TestMercadoLibreAndFINKTimelineEntries(t *testing.T) {
	var foundMELI, foundFINK bool
	for _, entry := range timelineEntries {
		switch entry.Date {
		case "October 2022":
			foundMELI = true
			for _, fact := range []string{"Mercado Libre", "Go", "observability"} {
				if !strings.Contains(entry.Body, fact) {
					t.Errorf("October 2022 entry missing %q: %q", fact, entry.Body)
				}
			}
		case "May 2025":
			foundFINK = true
			for _, fact := range []string{"FINK AI", "part-time", "20,000+", "4.2M+"} {
				if !strings.Contains(entry.Body, fact) {
					t.Errorf("May 2025 entry missing %q: %q", fact, entry.Body)
				}
			}
		}
	}
	if !foundMELI {
		t.Error("missing October 2022 Mercado Libre timeline entry")
	}
	if !foundFINK {
		t.Error("missing May 2025 FINK AI timeline entry")
	}
}

func TestLinksValid(t *testing.T) {
	groups := map[string][]Link{
		"social":  socialLinks,
		"contact": contactLinks,
		"courses": courseLinks,
		"resume":  {resumeLink},
	}

	for group, links := range groups {
		for _, link := range links {
			t.Run(group+"/"+link.Label, func(t *testing.T) {
				if link.Label == "" {
					t.Fatal("empty label")
				}
				switch {
				case strings.HasPrefix(link.Href, "https://"),
					strings.HasPrefix(link.Href, "mailto:"),
					strings.HasPrefix(link.Href, "tel:"),
					strings.HasPrefix(link.Href, "/"):
				default:
					t.Errorf("unexpected href %q", link.Href)
				}
			})
		}
	}
}

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

	for _, fact := range []string{"Mercado Libre", "ExxonMobil", "MBA in Data Science & Analytics", "Recognition Award", "FinkAI"} {
		if !strings.Contains(joined, fact) {
			t.Errorf("timeline is missing %q", fact)
		}
	}
	if !strings.Contains(aboutText, "Mercado Libre") {
		t.Errorf("about text should mention the current employer, Mercado Libre")
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

package main

import (
	"encoding/json"
	"encoding/xml"
	"strings"
	"testing"
)

// llms.txt, sitemap.xml and the JSON-LD block are generated from content.go
// rather than kept by hand, so these tests guard the generators, not files.

func TestLLMsTxtCoversTheProfile(t *testing.T) {
	body := llmsTxt()

	for _, want := range []string{
		"# " + personName,
		"## Summary",
		"## Focus",
		"## Career",
		"## Stack",
		"## Honors",
		"## Links",
		"## Contact",
		"## Education",
		"## Availability",
		"Mercado Libre",
		"FINK AI",
		"ExxonMobil",
		"https://guisolski.github.io/",
		"https://www.linkedin.com/in/guilherme-solski-alves/",
		"https://github.com/guisolski",
		"https://guisolski.github.io/assets/pdf/cv.pdf",
		"CV (English, PDF)",
		"guilhermesolskialves@gmail.com",
	} {
		if !strings.Contains(body, want) {
			t.Errorf("llms.txt missing %q", want)
		}
	}

	// Relative hrefs must be absolute here: a model that reads this file has
	// no base URL to resolve "/assets/pdf/cv.pdf" against.
	for _, line := range strings.Split(body, "\n") {
		if _, href, ok := strings.Cut(line, ": /"); ok && strings.HasPrefix(line, "- ") {
			t.Errorf("llms.txt has a root-relative link: %q (%q)", line, href)
		}
	}
}

func TestLLMsTxtListsEveryMilestone(t *testing.T) {
	body := llmsTxt()
	for _, e := range timelineEntries {
		if !strings.Contains(body, e.Date) {
			t.Errorf("llms.txt is missing the %s milestone", e.Date)
		}
	}
}

func TestSitemapXML(t *testing.T) {
	const lastmod = "2026-01-31"
	body := sitemapXML(lastmod)

	var parsed struct {
		URLs []struct {
			Loc     string `xml:"loc"`
			LastMod string `xml:"lastmod"`
		} `xml:"url"`
	}
	if err := xml.Unmarshal([]byte(body), &parsed); err != nil {
		t.Fatalf("sitemap is not well-formed XML: %v\n%s", err, body)
	}
	if len(parsed.URLs) != 1 {
		t.Fatalf("sitemap has %d urls, want 1", len(parsed.URLs))
	}
	if got, want := parsed.URLs[0].Loc, siteURL+"/"; got != want {
		t.Errorf("sitemap loc = %q, want %q", got, want)
	}
	if parsed.URLs[0].LastMod != lastmod {
		t.Errorf("sitemap lastmod = %q, want %q", parsed.URLs[0].LastMod, lastmod)
	}
}

func TestTodayIsISODate(t *testing.T) {
	if got := today(); len(got) != 10 || strings.Count(got, "-") != 2 {
		t.Errorf("today() = %q, want an ISO 8601 date", got)
	}
}

// The JSON-LD is injected raw into <head>, so the one thing that must never
// happen is a payload that closes the script element around it.
func TestPersonJSONLDCannotEscapeItsScriptTag(t *testing.T) {
	tag := personJSONLD()
	if !strings.HasPrefix(tag, `<script type="application/ld+json">`) {
		t.Fatalf("JSON-LD is not wrapped in a script tag: %.60q", tag)
	}
	body := strings.TrimSuffix(strings.TrimPrefix(tag, `<script type="application/ld+json">`), "</script>")
	if strings.ContainsAny(body, "<>&") {
		t.Errorf("JSON-LD payload contains an unescaped <, > or &: %s", body)
	}
}

func TestPersonJSONLDShape(t *testing.T) {
	raw, err := json.Marshal(personDocument())
	if err != nil {
		t.Fatalf("marshalling the profile: %v", err)
	}

	var doc struct {
		Context    string `json:"@context"`
		Type       string `json:"@type"`
		MainEntity struct {
			Type          string   `json:"@type"`
			Name          string   `json:"name"`
			Email         string   `json:"email"`
			Image         string   `json:"image"`
			URL           string   `json:"url"`
			SameAs        []string `json:"sameAs"`
			KnowsAbout    []string `json:"knowsAbout"`
			Award         []string `json:"award"`
			KnowsLanguage []struct {
				Name string `json:"name"`
			} `json:"knowsLanguage"`
			WorksFor []struct {
				Name string `json:"name"`
			} `json:"worksFor"`
		} `json:"mainEntity"`
	}
	if err := json.Unmarshal(raw, &doc); err != nil {
		t.Fatalf("profile is not valid JSON: %v", err)
	}

	if doc.Context != "https://schema.org" || doc.Type != "ProfilePage" {
		t.Errorf("document is %s/%s, want https://schema.org/ProfilePage", doc.Context, doc.Type)
	}
	p := doc.MainEntity
	if p.Type != "Person" || p.Name != personName {
		t.Errorf("mainEntity = %s %q, want Person %q", p.Type, p.Name, personName)
	}
	if p.Email != personEmail {
		t.Errorf("email = %q, want %q", p.Email, personEmail)
	}
	// Absolute URLs: a consumer of the JSON-LD may never have fetched the page.
	for _, field := range []struct{ name, value string }{
		{"url", p.URL},
		{"image", p.Image},
	} {
		if !strings.HasPrefix(field.value, "https://") {
			t.Errorf("%s = %q, want an absolute URL", field.name, field.value)
		}
	}
	for _, list := range []struct {
		name  string
		items []string
	}{
		{"sameAs", p.SameAs},
		{"knowsAbout", p.KnowsAbout},
		{"award", p.Award},
	} {
		if len(list.items) == 0 {
			t.Errorf("%s is empty", list.name)
		}
	}
	if len(p.KnowsLanguage) != len(spokenLanguages) {
		t.Errorf("knowsLanguage has %d entries, want %d", len(p.KnowsLanguage), len(spokenLanguages))
	}
	if len(p.WorksFor) == 0 || p.WorksFor[0].Name != "Mercado Libre" {
		t.Errorf("worksFor should lead with the current employer, got %+v", p.WorksFor)
	}
}

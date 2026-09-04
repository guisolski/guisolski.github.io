package main

import (
	"bytes"
	"encoding/json"
)

// The structured-data layer. Everything here is built from the same values the
// page renders, so the JSON-LD cannot drift away from the visible prose.

type ldOrg struct {
	Type string `json:"@type"`
	Name string `json:"name"`
}

type ldAddress struct {
	Type     string `json:"@type"`
	Locality string `json:"addressLocality"`
	Region   string `json:"addressRegion"`
	Country  string `json:"addressCountry"`
}

type ldLanguage struct {
	Type          string `json:"@type"`
	Name          string `json:"name"`
	AlternateName string `json:"alternateName"`
}

type ldOccupation struct {
	Type   string `json:"@type"`
	Name   string `json:"name"`
	Skills string `json:"skills"`
}

type ldDemand struct {
	Type string `json:"@type"`
	Name string `json:"name"`
}

type ldPerson struct {
	Type          string       `json:"@type"`
	ID            string       `json:"@id"`
	Name          string       `json:"name"`
	JobTitle      string       `json:"jobTitle"`
	Description   string       `json:"description"`
	URL           string       `json:"url"`
	Image         string       `json:"image"`
	Email         string       `json:"email"`
	Telephone     string       `json:"telephone"`
	Address       ldAddress    `json:"address"`
	WorksFor      []ldOrg      `json:"worksFor"`
	AlumniOf      []ldOrg      `json:"alumniOf"`
	HasOccupation ldOccupation `json:"hasOccupation"`
	KnowsAbout    []string     `json:"knowsAbout"`
	KnowsLanguage []ldLanguage `json:"knowsLanguage"`
	Award         []string     `json:"award"`
	SameAs        []string     `json:"sameAs"`
	Seeks         ldDemand     `json:"seeks"`
}

type ldProfilePage struct {
	Context    string   `json:"@context"`
	Type       string   `json:"@type"`
	MainEntity ldPerson `json:"mainEntity"`
}

func personSameAs() []string {
	out := make([]string, 0, len(socialLinks))
	for _, l := range socialLinks {
		if isExternal(l.Href) {
			out = append(out, l.Href)
		}
	}
	return out
}

func personDocument() ldProfilePage {
	works := make([]ldOrg, 0, len(employers))
	for _, name := range employers {
		works = append(works, ldOrg{Type: "Organization", Name: name})
	}

	alumni := make([]ldOrg, 0, len(schools))
	for _, s := range schools {
		alumni = append(alumni, ldOrg{Type: s.Type, Name: s.Name})
	}

	langs := make([]ldLanguage, 0, len(knowsLanguage))
	for _, l := range knowsLanguage {
		langs = append(langs, ldLanguage{Type: "Language", Name: l.Name, AlternateName: l.Code})
	}

	return ldProfilePage{
		Context: "https://schema.org",
		Type:    "ProfilePage",
		MainEntity: ldPerson{
			Type:        "Person",
			ID:          siteURL + "/#person",
			Name:        personName,
			JobTitle:    personJobTitle,
			Description: aboutLead,
			URL:         siteURL + "/",
			Image:       siteURL + profileImage,
			Email:       personEmail,
			Telephone:   personPhoneE164,
			Address: ldAddress{
				Type:     "PostalAddress",
				Locality: personLocality,
				Region:   personRegion,
				Country:  personCountry,
			},
			WorksFor: works,
			AlumniOf: alumni,
			HasOccupation: ldOccupation{
				Type:   "Occupation",
				Name:   personJobTitle,
				Skills: "Go, distributed systems, observability, third-party API integration, developer tooling",
			},
			KnowsAbout:    knowsAbout,
			KnowsLanguage: langs,
			Award:         awards,
			SameAs:        personSameAs(),
			Seeks: ldDemand{
				Type: "Demand",
				Name: availability + ", " + "remote or relocation",
			},
		},
	}
}

// personJSONLD renders the <script> tag injected into every page's <head>.
// json.Marshal escapes <, > and & as \u003c, \u003e and \u0026, so the
// payload can never close the script element it sits in.
func personJSONLD() string {
	var buf bytes.Buffer
	buf.WriteString(`<script type="application/ld+json">`)
	body, err := json.Marshal(personDocument())
	if err != nil {
		// The document is built from constants; a failure here is a programming
		// error, and an empty object keeps the page valid either way.
		buf.WriteString("{}")
	} else {
		buf.Write(body)
	}
	buf.WriteString(`</script>`)
	return buf.String()
}

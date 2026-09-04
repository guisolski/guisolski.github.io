package main

import (
	"fmt"
	"strings"
)

// llms.txt and sitemap.xml are generated from the same values the page renders,
// so there is no second, hand-kept copy of these facts to drift out of date.

func writeLinkSection(b *strings.Builder, heading string, groups ...[]Link) {
	fmt.Fprintf(b, "\n## %s\n\n", heading)
	for _, group := range groups {
		for _, l := range group {
			if l.Href == "" {
				fmt.Fprintf(b, "- %s\n", l.Label)
				continue
			}
			href := l.Href
			if strings.HasPrefix(href, "/") {
				href = siteURL + href
			}
			fmt.Fprintf(b, "- %s: %s\n", l.Label, href)
		}
	}
}

func llmsTxt() string {
	var b strings.Builder

	fmt.Fprintf(&b, "# %s\n\n> %s\n", personName, personTagline)
	fmt.Fprintf(&b, "\n## Summary\n\n%s Based in %s (%s).\n", aboutText, personLocation, personTimezone)

	b.WriteString("\n## Focus\n\n")
	for _, f := range focusAreas {
		fmt.Fprintf(&b, "- %s: %s\n", f.Title, f.Body)
	}

	b.WriteString("\n## Career\n\n")
	for _, e := range timelineEntries {
		fmt.Fprintf(&b, "- %s — %s\n", e.Date, e.Body)
	}

	b.WriteString("\n## Stack\n\n")
	for _, g := range stack {
		fmt.Fprintf(&b, "- %s: %s\n", g.Label, strings.Join(g.Items, ", "))
	}

	b.WriteString("\n## Spoken languages\n\n")
	for _, l := range spokenLanguages {
		fmt.Fprintf(&b, "- %s (%s)\n", l.Label, strings.ToLower(l.Tag))
	}

	b.WriteString("\n## Honors\n\n")
	for _, a := range awards {
		fmt.Fprintf(&b, "- %s\n", a)
	}

	fmt.Fprintf(&b, "\n## Links\n\n- Site: %s/\n", siteURL)
	for _, l := range socialLinks {
		fmt.Fprintf(&b, "- %s: %s\n", l.Label, l.Href)
	}
	for _, l := range resumeLinks {
		fmt.Fprintf(&b, "- %s: %s%s\n", l.Label, siteURL, l.Href)
	}

	writeLinkSection(&b, "Contact", contactLinks)
	writeLinkSection(&b, "Education", courseLinks)

	fmt.Fprintf(&b, "\n## Availability\n\n%s — %s.\n", availability, strings.ToLower(availabilityQualifier))

	return b.String()
}

func sitemapXML(lastmod string) string {
	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <url>
    <loc>%s/</loc>
    <lastmod>%s</lastmod>
    <changefreq>monthly</changefreq>
    <priority>1.0</priority>
  </url>
</urlset>
`, siteURL, lastmod)
}

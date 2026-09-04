package main

import (
	"fmt"
	"strings"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// Each card takes its stop on the dawn spectrum from its place in the reading
// order, so the page warms as you scroll — the same way the career rail does.
type cardSpec struct {
	Title string
	Hue   int
	Body  func() []app.UI
}

var cardSpecs = []cardSpec{
	{"Around the web", 245, socialLinksBody},
	{"Contact", 288, contactLinksBody},
	{"CV", 320, resumeBody},
	{"Education", 352, courseLinksBody},
	{"Stack", 14, stackBody},
	{"Languages", 35, spokenLanguagesBody},
}

func socialLinksBody() []app.UI {
	return []app.UI{linkList(socialLinks)}
}

func contactLinksBody() []app.UI {
	return []app.UI{linkList(contactLinks)}
}

func courseLinksBody() []app.UI {
	return []app.UI{linkList(courseLinks)}
}

func resumeBody() []app.UI {
	return []app.UI{
		linkList(resumeLinks),
		app.P().Class("card__note").Body(
			app.Text(availability),
			app.Br(),
			app.Text(availabilityQualifier),
		),
	}
}

func renderStackGroup(i int) app.UI {
	g := stack[i]
	return app.Div().Class("stack__group").Body(
		app.Dt().Class("stack__label").Text(g.Label),
		app.Dd().Class("stack__items").Text(strings.Join(g.Items, " · ")),
	)
}

func stackBody() []app.UI {
	return []app.UI{
		app.Dl().Class("stack").Body(
			app.Range(stack).Slice(renderStackGroup),
		),
	}
}

func spokenLanguagesBody() []app.UI {
	return []app.UI{
		app.Ul().Class("tags").Body(
			app.Range(spokenLanguages).Slice(renderSpokenLanguagePill),
		),
	}
}

func renderSpokenLanguagePill(i int) app.UI {
	l := spokenLanguages[i]
	return app.Li().Class("tags__item").Body(
		app.Span().Text(l.Label),
		app.Span().Class("tags__level").Text(l.Tag),
	)
}

func renderCard(i int) app.UI {
	s := cardSpecs[i]
	return card(s, s.Body()...)
}

func Cards() app.UI {
	return app.Section().Class("cards").Aria("label", "Details").Body(
		app.Range(cardSpecs).Slice(renderCard),
	)
}

func card(s cardSpec, body ...app.UI) app.UI {
	head := []app.UI{app.H2().Class("card__title").Text(s.Title)}
	return app.Article().
		Class("card").
		Style("--card-h", fmt.Sprint(s.Hue)).
		Body(append(head, body...)...)
}

func renderFocus(i int) app.UI {
	f := focusAreas[i]
	return app.Article().
		Class("focus__item").
		Style("--focus-h", fmt.Sprint(f.Hue)).
		Body(
			app.Span().Class("focus__rule").Aria("hidden", "true"),
			app.H3().Class("focus__title").Text(f.Title),
			app.P().Class("focus__body").Text(f.Body),
		)
}

func FocusSection() app.UI {
	return app.Section().Class("focus").Aria("labelledby", "focus-heading").Body(
		app.H2().Class("section-label").ID("focus-heading").Text("What I work on"),
		app.Div().Class("focus__grid").Body(
			app.Range(focusAreas).Slice(renderFocus),
		),
	)
}

type linkRenderer struct {
	links []Link
}

func (r linkRenderer) render(i int) app.UI {
	l := r.links[i]
	if l.Href == "" {
		return app.Li().Body(
			app.Span().Class("links__item links__item--plain").Body(
				icon(l.Icon),
				app.Span().Text(l.Label),
			),
		)
	}
	a := app.A().Class("links__item").Href(l.Href).Body(
		icon(l.Icon),
		app.Span().Text(l.Label),
	)
	// go-app intercepts same-origin <a> clicks as SPA navigation unless
	// target=_blank or download is set; open assets in a new tab so PDFs load.
	if isExternal(l.Href) || isAssetLink(l.Href) {
		a = a.Target("_blank").Rel("noopener")
	}
	return app.Li().Body(a)
}

func isExternal(href string) bool {
	return len(href) > 4 && href[:4] == "http"
}

func isAssetLink(href string) bool {
	return strings.HasPrefix(href, "/assets/")
}

func linkList(links []Link) app.UI {
	return app.Ul().Class("links").Body(
		app.Range(links).Slice(linkRenderer{links}.render),
	)
}

func svgIcon(width, height int, class string, strokeWidth float64, inner string) app.UI {
	classAttr := ""
	if class != "" {
		classAttr = ` class="` + class + `"`
	}
	return app.Raw(fmt.Sprintf(
		`<svg%s viewBox="0 0 24 24" width="%d" height="%d" fill="none" stroke="currentColor" stroke-width="%g" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true">%s</svg>`,
		classAttr, width, height, strokeWidth, inner,
	))
}

// The ✦ that marks honors and opens the night scene is drawn rather than typed:
// a glyph depends on whatever font happens to resolve, a path does not.
func starIcon(size int) app.UI {
	return app.Raw(fmt.Sprintf(
		`<svg class="star" viewBox="0 0 24 24" width="%d" height="%d" fill="currentColor" aria-hidden="true"><path d="M12 1.5 13.7 9.6 21.8 11.3 13.7 13 12 21.1 10.3 13 2.2 11.3 10.3 9.6Z"/></svg>`,
		size, size,
	))
}

var iconPaths = map[string]string{
	"github":   `<path d="M9 19c-4.3 1.4-4.3-2.5-6-3m12 5v-3.5c0-1 .1-1.4-.5-2 2.8-.3 5.5-1.4 5.5-6a4.6 4.6 0 0 0-1.3-3.2 4.2 4.2 0 0 0-.1-3.2s-1.1-.3-3.5 1.3a12.3 12.3 0 0 0-6.2 0C6.5 2.8 5.4 3.1 5.4 3.1a4.2 4.2 0 0 0-.1 3.2A4.6 4.6 0 0 0 4 9.5c0 4.6 2.7 5.7 5.5 6-.6.6-.6 1.2-.5 2V21"/>`,
	"linkedin": `<path d="M16 8a6 6 0 0 1 6 6v7h-4v-7a2 2 0 0 0-4 0v7h-4v-13h4v2a5 5 0 0 1 2-2z"/><rect x="2" y="9" width="4" height="12"/><circle cx="4" cy="4" r="2"/>`,
	"code":     `<polyline points="16 18 22 12 16 6"/><polyline points="8 6 2 12 8 18"/>`,
	"trophy":   `<path d="M8 21h8m-4-4v4m-6-17h12v6a6 6 0 0 1-12 0zM6 4H3v2a4 4 0 0 0 3 3.9M18 4h3v2a4 4 0 0 1-3 3.9"/>`,
	"mail":     `<rect x="2" y="4" width="20" height="16" rx="2"/><polyline points="2,7 12,13 22,7"/>`,
	"phone":    `<path d="M22 16.9v3a2 2 0 0 1-2.2 2 19.8 19.8 0 0 1-8.6-3 19.5 19.5 0 0 1-6-6 19.8 19.8 0 0 1-3-8.7A2 2 0 0 1 4.1 2h3a2 2 0 0 1 2 1.7c.1.9.3 1.8.6 2.7a2 2 0 0 1-.4 2.1L8 9.8a16 16 0 0 0 6 6l1.3-1.3a2 2 0 0 1 2.1-.4c.9.3 1.8.5 2.7.6a2 2 0 0 1 1.9 2.2z"/>`,
	"pin":      `<path d="M12 21s7-5.5 7-11a7 7 0 0 0-14 0c0 5.5 7 11 7 11z"/><circle cx="12" cy="10" r="2.5"/>`,
	"badge":    `<circle cx="12" cy="8" r="6"/><path d="M15.5 13 17 22l-5-3-5 3 1.5-9"/>`,
	"file":     `<path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/><polyline points="14 2 14 8 20 8"/>`,
}

func icon(name string) app.UI {
	p, ok := iconPaths[name]
	if !ok {
		return app.Text("")
	}
	return svgIcon(16, 16, "icon", 1.8, p)
}

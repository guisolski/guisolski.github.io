package main

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type cardSpec struct {
	Title string
	Body  func() []app.UI
}

var cardSpecs = []cardSpec{
	{"Around the web", socialLinksBody},
	{"Contact", contactLinksBody},
	{"Education", courseLinksBody},
	{"Programming languages", programmingLanguagesBody},
	{"Languages", spokenLanguagesBody},
	{"Résumé", resumeBody},
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
		linkList([]Link{resumeLink}),
		app.P().Class("card__note").Text("Ask me for a Portuguese version by email."),
	}
}

func programmingLanguagesBody() []app.UI {
	return []app.UI{
		app.Ul().Class("pills").Body(
			app.Range(programmingLanguages).Slice(renderLanguagePill),
		),
	}
}

func renderLanguagePill(i int) app.UI {
	return app.Li().Class("pill").Text(programmingLanguages[i])
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
	return card(s.Title, s.Body()...)
}

func Cards() app.UI {
	return app.Section().Class("cards").Aria("label", "Details").Body(
		app.Range(cardSpecs).Slice(renderCard),
	)
}

func card(title string, body ...app.UI) app.UI {
	return app.Article().Class("card").Body(
		append([]app.UI{app.H2().Class("card__title").Text(title)}, body...)...,
	)
}

type linkRenderer struct {
	links []Link
}

func (r linkRenderer) render(i int) app.UI {
	l := r.links[i]
	a := app.A().Class("links__item").Href(l.Href).Body(
		icon(l.Icon),
		app.Span().Text(l.Label),
	)
	if isExternal(l.Href) {
		a = a.Target("_blank").Rel("noopener")
	}
	return app.Li().Body(a)
}

func isExternal(href string) bool {
	return len(href) > 4 && href[:4] == "http"
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

var iconPaths = map[string]string{
	"github":   `<path d="M9 19c-4.3 1.4-4.3-2.5-6-3m12 5v-3.5c0-1 .1-1.4-.5-2 2.8-.3 5.5-1.4 5.5-6a4.6 4.6 0 0 0-1.3-3.2 4.2 4.2 0 0 0-.1-3.2s-1.1-.3-3.5 1.3a12.3 12.3 0 0 0-6.2 0C6.5 2.8 5.4 3.1 5.4 3.1a4.2 4.2 0 0 0-.1 3.2A4.6 4.6 0 0 0 4 9.5c0 4.6 2.7 5.7 5.5 6-.6.6-.6 1.2-.5 2V21"/>`,
	"linkedin": `<path d="M16 8a6 6 0 0 1 6 6v7h-4v-7a2 2 0 0 0-4 0v7h-4v-13h4v2a5 5 0 0 1 2-2z"/><rect x="2" y="9" width="4" height="12"/><circle cx="4" cy="4" r="2"/>`,
	"code":     `<polyline points="16 18 22 12 16 6"/><polyline points="8 6 2 12 8 18"/>`,
	"trophy":   `<path d="M8 21h8m-4-4v4m-6-17h12v6a6 6 0 0 1-12 0zM6 4H3v2a4 4 0 0 0 3 3.9M18 4h3v2a4 4 0 0 1-3 3.9"/>`,
	"mail":     `<rect x="2" y="4" width="20" height="16" rx="2"/><polyline points="2,7 12,13 22,7"/>`,
	"phone":    `<path d="M22 16.9v3a2 2 0 0 1-2.2 2 19.8 19.8 0 0 1-8.6-3 19.5 19.5 0 0 1-6-6 19.8 19.8 0 0 1-3-8.7A2 2 0 0 1 4.1 2h3a2 2 0 0 1 2 1.7c.1.9.3 1.8.6 2.7a2 2 0 0 1-.4 2.1L8 9.8a16 16 0 0 0 6 6l1.3-1.3a2 2 0 0 1 2.1-.4c.9.3 1.8.5 2.7.6a2 2 0 0 1 1.9 2.2z"/>`,
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

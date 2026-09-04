package main

import (
	"time"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

const introEyebrow = "Backend engineer · Go · Remote, UTC−3"

type Home struct {
	app.Compo

	leaving bool
}

func pageClass(leaving bool) string {
	if leaving {
		return "page page--leaving"
	}
	return "page"
}

// OnPreRender runs while the static site is generated, which is the only place
// the canonical link can be set for a prerendered page.
func (h *Home) OnPreRender(ctx app.Context) {
	ctx.Page().SetCanonicalLink("/")
}

// The avatar wears the dawn spectrum as a slowly orbiting ring. The ring is its
// own element rather than a conic border, so the rotation is a plain transform
// on the compositor instead of an interpolated custom property.
func renderAvatar() app.UI {
	return app.Span().Class("avatar").Body(
		app.Span().Class("avatar__ring").Aria("hidden", "true"),
		app.Img().
			Class("avatar__img").
			Src(profileImage).
			Alt("Portrait of "+personName).
			Width(76).
			Height(76),
	)
}

// The h1 holds the name as text and nothing else: the night scene is a sibling
// link, so the page's only heading stays a heading rather than a control.
func (h *Home) renderIntro() app.UI {
	return app.Header().Class("intro").Body(
		renderAvatar(),
		app.P().Class("eyebrow").Text(introEyebrow),
		app.H1().Class("intro__name").Body(
			app.Span().Text(personName),
			app.A().
				Class("intro__star").
				Href("/relax").
				Title("Night scene").
				Aria("label", "Open the hidden night scene").
				OnClick(h.onStarClick).
				Body(starIcon(16)),
		),
		app.P().Class("intro__lead").Text(aboutLead),
		app.P().Class("intro__about").Text(aboutBody),
	)
}

func renderFooter() app.UI {
	return app.Footer().Class("footer").Body(
		app.Span().Class("footer__rule").Aria("hidden", "true"),
		app.Div().Class("footer__row").Body(
			app.P().Body(
				app.Text("© "+personName+" · Built with "),
				app.A().Href("https://go.dev").Text("Go"),
				app.Text(" and WebAssembly — "),
				app.A().Href("https://github.com/guisolski/guisolski.github.io").Text("source"),
			),
			app.P().Class("footer__place").Text(personLocation+" · "+personTimezone),
		),
	)
}

func (h *Home) Render() app.UI {
	return app.Main().Class(pageClass(h.leaving)).Body(
		h.renderIntro(),
		&Timeline{},
		FocusSection(),
		Cards(),
		renderFooter(),
	)
}

func (h *Home) navigateToRelax(ctx app.Context) {
	ctx.Navigate("/relax")
}

// The href is real so crawlers follow it; the click is intercepted only to let
// the page fade out before the night scene takes over.
func (h *Home) onStarClick(ctx app.Context, e app.Event) {
	e.PreventDefault()
	if h.leaving {
		return
	}
	h.leaving = true
	ctx.After(500*time.Millisecond, h.navigateToRelax)
}

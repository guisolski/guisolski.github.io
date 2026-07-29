package main

import (
	"time"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// Home is the single CV page: header, about, career timeline, and cards.
type Home struct {
	app.Compo

	leaving bool // true while the slide-out transition to /relax plays
}

func (h *Home) Render() app.UI {
	pageClass := "page"
	if h.leaving {
		pageClass += " page--leaving"
	}

	return app.Main().Class(pageClass).Body(
		app.Header().Class("intro").Body(
			app.Img().
				Class("intro__avatar").
				Src("/assets/images/profile.png").
				Alt("Portrait of Guilherme Solski Alves").
				Width(72).
				Height(72),
			app.P().Class("eyebrow").Text("Software developer · Mercado Libre"),
			app.H1().Class("intro__name").Body(
				app.Button().
					Class("intro__name-button").
					Title("Take a break").
					Aria("label", "Guilherme Solski Alves — open the hidden night scene").
					OnClick(h.onNameClick).
					Text("Guilherme Solski Alves"),
			),
			app.P().Class("intro__about").Text(aboutText),
		),
		&Timeline{},
		Cards(),
		app.Footer().Class("footer").Body(
			app.P().Body(
				app.Text("© Guilherme Solski Alves · Built with "),
				app.A().Href("https://go.dev").Text("Go"),
				app.Text(" and WebAssembly — "),
				app.A().Href("https://github.com/guisolski/guisolski.github.io").Text("source"),
			),
		),
	)
}

func (h *Home) onNameClick(ctx app.Context, e app.Event) {
	if h.leaving {
		return
	}
	h.leaving = true
	ctx.After(500*time.Millisecond, func(ctx app.Context) {
		ctx.Navigate("/relax")
	})
}

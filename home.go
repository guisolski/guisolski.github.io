package main

import (
	"time"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

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

func (h *Home) renderIntro() app.UI {
	return app.Header().Class("intro").Body(
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
	)
}

func renderFooter() app.UI {
	return app.Footer().Class("footer").Body(
		app.P().Body(
			app.Text("© Guilherme Solski Alves · Built with "),
			app.A().Href("https://go.dev").Text("Go"),
			app.Text(" and WebAssembly — "),
			app.A().Href("https://github.com/guisolski/guisolski.github.io").Text("source"),
		),
	)
}

func (h *Home) Render() app.UI {
	return app.Main().Class(pageClass(h.leaving)).Body(
		h.renderIntro(),
		&Timeline{},
		Cards(),
		renderFooter(),
	)
}

func (h *Home) navigateToRelax(ctx app.Context) {
	ctx.Navigate("/relax")
}

func (h *Home) onNameClick(ctx app.Context, e app.Event) {
	if h.leaving {
		return
	}
	h.leaving = true
	ctx.After(500*time.Millisecond, h.navigateToRelax)
}

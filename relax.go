package main

import (
	"math/rand"
	"time"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// Relax is the hidden night-scene page: an animated starfield, a fractal
// mountain silhouette, and a fractal tree, all computed in Go/WebAssembly.
// The scene layers appear in stages, like the original page did.
type Relax struct {
	app.Compo

	field        *starfield
	rng          *rand.Rand
	raf          app.Func
	resize       app.Func
	running      bool
	terrainShown bool
	treeShown    bool
}

func (r *Relax) Render() app.UI {
	return app.Div().Class("relax").Body(
		app.Canvas().ID("bgCanvas").Class("relax__canvas"),
		app.Canvas().ID("terCanvas").Class("relax__canvas"),
		app.Canvas().ID("treeCanvas").Class("relax__canvas"),
		app.A().Class("relax__back").Href("/").Text("← back"),
	)
}

func (r *Relax) OnMount(ctx app.Context) {
	r.running = true
	r.rng = rand.New(rand.NewSource(time.Now().UnixNano()))
	r.setupScene()

	r.raf = app.FuncOf(func(this app.Value, args []app.Value) any {
		if !r.running {
			return nil
		}
		r.field.step()
		app.Window().Call("requestAnimationFrame", r.raf)
		return nil
	})
	app.Window().Call("requestAnimationFrame", r.raf)

	r.resize = app.FuncOf(func(this app.Value, args []app.Value) any {
		if r.running {
			r.setupScene()
		}
		return nil
	})
	app.Window().Call("addEventListener", "resize", r.resize)

	ctx.After(2*time.Second, func(app.Context) {
		if r.running {
			r.terrainShown = true
			r.paintTerrain()
		}
	})
	ctx.After(4500*time.Millisecond, func(app.Context) {
		if r.running {
			r.treeShown = true
			r.paintTree()
		}
	})
}

func (r *Relax) OnDismount() {
	r.running = false
	if r.raf != nil {
		r.raf.Release()
		r.raf = nil
	}
	if r.resize != nil {
		app.Window().Call("removeEventListener", "resize", r.resize)
		r.resize.Release()
		r.resize = nil
	}
}

// setupScene sizes all three canvases to the viewport and (re)initializes
// whatever layers are currently visible.
func (r *Relax) setupScene() {
	w, h := app.Window().Size()
	for _, id := range []string{"bgCanvas", "terCanvas", "treeCanvas"} {
		canvas := app.Window().GetElementByID(id)
		if !canvas.Truthy() {
			return
		}
		canvas.Set("width", w)
		canvas.Set("height", h)
	}
	r.field = newStarfield(context2D("bgCanvas"), float64(w), float64(h), r.rng)
	if r.terrainShown {
		r.paintTerrain()
	}
	if r.treeShown {
		r.paintTree()
	}
}

func (r *Relax) paintTerrain() {
	w, h := app.Window().Size()
	drawTerrain(context2D("terCanvas"), float64(w), float64(h), r.rng)
}

func (r *Relax) paintTree() {
	w, h := app.Window().Size()
	fw, fh := float64(w), float64(h)
	drawTree(context2D("treeCanvas"), fw-fw/5, fh, fh, r.rng)
}

func context2D(canvasID string) app.Value {
	return app.Window().GetElementByID(canvasID).Call("getContext", "2d")
}

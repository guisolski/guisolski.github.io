package main

import (
	"math/rand"
	"time"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

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

func (r *Relax) onAnimationFrame(this app.Value, args []app.Value) any {
	if !r.running {
		return nil
	}
	r.field.step()
	app.Window().Call("requestAnimationFrame", r.raf)
	return nil
}

func (r *Relax) onResize(this app.Value, args []app.Value) any {
	if r.running {
		r.setupScene()
	}
	return nil
}

func (r *Relax) revealTerrain(app.Context) {
	if r.running {
		r.terrainShown = true
		r.paintTerrain()
	}
}

func (r *Relax) revealTree(app.Context) {
	if r.running {
		r.treeShown = true
		r.paintTree()
	}
}

func (r *Relax) OnMount(ctx app.Context) {
	r.running = true
	r.rng = rand.New(rand.NewSource(time.Now().UnixNano()))
	r.setupScene()

	r.raf = app.FuncOf(r.onAnimationFrame)
	app.Window().Call("requestAnimationFrame", r.raf)

	r.resize = app.FuncOf(r.onResize)
	app.Window().Call("addEventListener", "resize", r.resize)

	ctx.After(2*time.Second, r.revealTerrain)
	ctx.After(4500*time.Millisecond, r.revealTree)
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

func windowSizeF() (float64, float64) {
	w, h := app.Window().Size()
	return float64(w), float64(h)
}

func (r *Relax) paintTerrain() {
	w, h := windowSizeF()
	drawTerrain(context2D("terCanvas"), w, h, r.rng)
}

func (r *Relax) paintTree() {
	w, h := windowSizeF()
	drawTree(context2D("treeCanvas"), w-w/5, h, h, r.rng)
}

func context2D(canvasID string) app.Value {
	return app.Window().GetElementByID(canvasID).Call("getContext", "2d")
}

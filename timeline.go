package main

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

const swipeThreshold = 40 // px of horizontal travel that counts as a swipe

// The timeline is colored as a dawn: the first year sits at night indigo and
// the hue travels through dusk rose to star amber at the current year — the
// same night-to-day story as the scene hidden behind the name.
const (
	dawnStartHue = 245 // night indigo
	dawnHueSpan  = 150 // degrees to star amber (35°)
)

// yearHue maps a timeline index onto the dawn spectrum, in degrees.
func yearHue(i, last int) int {
	if last <= 0 {
		return dawnStartHue
	}
	return (dawnStartHue + dawnHueSpan*i/last) % 360
}

// fillClip reveals the dawn gradient up to the selected year by clipping the
// remainder of the line from its right edge.
func fillClip(selected, last int) string {
	remaining := 100.0
	if last > 0 {
		remaining = 100 * (1 - float64(selected)/float64(last))
	}
	return fmt.Sprintf("inset(0 %.1f%% 0 0)", remaining)
}

// swipeStep converts a horizontal touch delta into a timeline step:
// -1 (swipe right, go back), +1 (swipe left, go forward), or 0.
func swipeStep(delta float64) int {
	switch {
	case delta > swipeThreshold:
		return -1
	case delta < -swipeThreshold:
		return 1
	default:
		return 0
	}
}

// dateClass returns the CSS classes for the year dot at index i.
func dateClass(i, selected int) string {
	switch {
	case i == selected:
		return "timeline__date timeline__date--active"
	case i < selected:
		return "timeline__date timeline__date--passed"
	default:
		return "timeline__date"
	}
}

// eventClass returns the CSS classes for the event panel at index i; the
// selected panel slides in from the side it is approached from.
func eventClass(i, selected, previous int) string {
	if i != selected {
		return "timeline__event"
	}
	if selected >= previous {
		return "timeline__event timeline__event--selected timeline__event--enter-right"
	}
	return "timeline__event timeline__event--selected timeline__event--enter-left"
}

// chevron renders the prev/next arrow glyph. dir is -1 for left, 1 for right.
func chevron(dir int) app.UI {
	points := "15 18 9 12 15 6"
	if dir > 0 {
		points = "9 18 15 12 9 6"
	}
	return app.Raw(`<svg viewBox="0 0 24 24" width="18" height="18" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><polyline points="` + points + `"/></svg>`)
}

// Timeline is the reactive horizontal career timeline: a rail of year stops
// with an animated filling line, and one event panel shown at a time.
type Timeline struct {
	app.Compo

	selected int
	previous int     // last selected index, used to derive slide direction
	touchX   float64 // clientX where the current touch started
}

func (t *Timeline) Render() app.UI {
	last := len(timelineEntries) - 1

	return app.Section().
		Class("timeline").
		Aria("label", "Career timeline").
		TabIndex(0).
		Style("--year-h", fmt.Sprint(yearHue(t.selected, last))).
		OnKeyDown(t.onKeyDown).
		Body(
			app.Div().Class("timeline__rail").Body(
				app.Button().
					Class("timeline__nav timeline__nav--prev").
					Aria("label", "Previous event").
					Disabled(t.selected == 0).
					OnClick(t.shiftBy(-1)).
					Body(chevron(-1)),
				app.Div().Class("timeline__track").Body(
					app.Div().Class("timeline__strip").Body(
						app.Span().Class("timeline__line").Aria("hidden", "true"),
						app.Span().
							Class("timeline__fill").
							Aria("hidden", "true").
							Style("clip-path", fillClip(t.selected, last)),
						app.Ol().Class("timeline__dates").Body(
							app.Range(timelineEntries).Slice(func(i int) app.UI {
								return app.Li().Body(
									app.Button().
										Class(dateClass(i, t.selected)).
										Style("--year-h", fmt.Sprint(yearHue(i, last))).
										Aria("label", timelineEntries[i].Date).
										OnClick(t.selectEvent(i)).
										Text(timelineEntries[i].Year),
								)
							}),
						),
					),
				),
				app.Button().
					Class("timeline__nav timeline__nav--next").
					Aria("label", "Next event").
					Disabled(t.selected == last).
					OnClick(t.shiftBy(1)).
					Body(chevron(1)),
			),
			app.Div().
				Class("timeline__events").
				On("touchstart", t.onTouchStart).
				On("touchend", t.onTouchEnd).
				Body(
					app.Range(timelineEntries).Slice(func(i int) app.UI {
						return app.Article().
							Class(eventClass(i, t.selected, t.previous)).
							Style("--year-h", fmt.Sprint(yearHue(i, last))).
							Aria("hidden", fmt.Sprint(i != t.selected)).
							Body(
								app.H3().Class("timeline__event-date").Text(timelineEntries[i].Date),
								app.P().Class("timeline__event-body").Text(timelineEntries[i].Body),
							)
					}),
				),
		)
}

func (t *Timeline) selectEvent(i int) app.EventHandler {
	return func(ctx app.Context, e app.Event) {
		t.move(ctx, i)
	}
}

func (t *Timeline) shiftBy(delta int) app.EventHandler {
	return func(ctx app.Context, e app.Event) {
		t.move(ctx, t.selected+delta)
	}
}

// move applies the state change and, after the update, keeps the newly
// selected year visible when the rail overflows on narrow screens.
func (t *Timeline) move(ctx app.Context, i int) {
	t.moveTo(i)
	ctx.Defer(centerActiveDate)
}

func (t *Timeline) moveTo(i int) {
	if i < 0 || i >= len(timelineEntries) || i == t.selected {
		return
	}
	t.previous = t.selected
	t.selected = i
}

func centerActiveDate(ctx app.Context) {
	active := app.Window().Get("document").Call("querySelector", ".timeline__date--active")
	if active.Truthy() {
		active.Call("scrollIntoView", map[string]any{"block": "nearest", "inline": "center"})
	}
}

func (t *Timeline) onKeyDown(ctx app.Context, e app.Event) {
	switch e.Get("key").String() {
	case "ArrowLeft":
		e.PreventDefault()
		t.move(ctx, t.selected-1)
	case "ArrowRight":
		e.PreventDefault()
		t.move(ctx, t.selected+1)
	}
}

func (t *Timeline) onTouchStart(ctx app.Context, e app.Event) {
	touches := e.Get("changedTouches")
	if touches.Truthy() && touches.Length() > 0 {
		t.touchX = touches.Index(0).Get("clientX").Float()
	}
}

func (t *Timeline) onTouchEnd(ctx app.Context, e app.Event) {
	touches := e.Get("changedTouches")
	if !touches.Truthy() || touches.Length() == 0 {
		return
	}
	delta := touches.Index(0).Get("clientX").Float() - t.touchX
	if step := swipeStep(delta); step != 0 {
		t.move(ctx, t.selected+step)
	}
}

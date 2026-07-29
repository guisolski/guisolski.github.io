package main

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

const swipeThreshold = 40 // px of horizontal travel that counts as a swipe

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
	progress := 0.0
	if last > 0 {
		progress = float64(t.selected) / float64(last)
	}

	return app.Section().
		Class("timeline").
		Aria("label", "Career timeline").
		TabIndex(0).
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
							Style("transform", fmt.Sprintf("scaleX(%.4f)", progress)),
						app.Ol().Class("timeline__dates").Body(
							app.Range(timelineEntries).Slice(func(i int) app.UI {
								return app.Li().Body(
									app.Button().
										Class(dateClass(i, t.selected)).
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
		t.moveTo(i)
	}
}

func (t *Timeline) shiftBy(delta int) app.EventHandler {
	return func(ctx app.Context, e app.Event) {
		t.moveTo(t.selected + delta)
	}
}

func (t *Timeline) moveTo(i int) {
	if i < 0 || i >= len(timelineEntries) || i == t.selected {
		return
	}
	t.previous = t.selected
	t.selected = i
}

func (t *Timeline) onKeyDown(ctx app.Context, e app.Event) {
	switch e.Get("key").String() {
	case "ArrowLeft":
		e.PreventDefault()
		t.moveTo(t.selected - 1)
	case "ArrowRight":
		e.PreventDefault()
		t.moveTo(t.selected + 1)
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
		t.moveTo(t.selected + step)
	}
}

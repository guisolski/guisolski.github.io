package main

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

const swipeThreshold = 40

const (
	dawnStartHue = 245
	dawnHueSpan  = 150
)

func yearHue(i, last int) int {
	if last <= 0 {
		return dawnStartHue
	}
	return (dawnStartHue + dawnHueSpan*i/last) % 360
}

func fillClip(selected, last int) string {
	remaining := 100.0
	if last > 0 {
		remaining = 100 * (1 - float64(selected)/float64(last))
	}
	return fmt.Sprintf("inset(0 %.1f%% 0 0)", remaining)
}

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

func eventClass(i, selected, previous int) string {
	if i != selected {
		return "timeline__event"
	}
	if selected >= previous {
		return "timeline__event timeline__event--selected timeline__event--enter-right"
	}
	return "timeline__event timeline__event--selected timeline__event--enter-left"
}

var chevronPoints = map[int]string{
	-1: "15 18 9 12 15 6",
	1:  "9 18 15 12 9 6",
}

func chevron(dir int) app.UI {
	return svgIcon(18, 18, "", 2, `<polyline points="`+chevronPoints[dir]+`"/>`)
}

type Timeline struct {
	app.Compo

	selected int
	previous int
	touchX   float64
}

func lastIndex() int {
	return len(timelineEntries) - 1
}

func (t *Timeline) OnInit() {
	t.selected = lastIndex()
	t.previous = t.selected
}

// The rail scrolls once it holds more years than fit the page; without this
// the initially selected (most recent) dot mounts out of view.
func (t *Timeline) OnMount(ctx app.Context) {
	ctx.Defer(centerActiveDate)
}

// dateLabel keeps the label a single flex item — the date button is a flex
// column, so separate children would stack the star above the year and push
// the dot off the rail.
func dateLabel(e TimelineEntry) app.UI {
	if e.Quote == "" {
		return app.Text(e.Year)
	}
	return app.Span().Body(
		app.Span().Class("timeline__honor-star").Aria("hidden", "true").Text("✦"),
		app.Text(e.Year),
	)
}

func (t *Timeline) renderDateDot(i int) app.UI {
	last := lastIndex()
	return app.Li().Body(
		app.Button().
			Class(dateClass(i, t.selected)).
			Style("--year-h", fmt.Sprint(yearHue(i, last))).
			Aria("label", timelineEntries[i].Date).
			OnClick(t.selectEvent(i)).
			Body(dateLabel(timelineEntries[i])),
	)
}

func (t *Timeline) renderEventPanel(i int) app.UI {
	last := lastIndex()
	e := timelineEntries[i]
	body := []app.UI{
		app.H3().Class("timeline__event-date").Text(e.Date),
		app.P().Class("timeline__event-body").Text(e.Body),
	}
	if e.Quote != "" {
		body = append(body, app.Blockquote().Class("timeline__event-quote").Text(e.Quote))
	}
	return app.Article().
		Class(eventClass(i, t.selected, t.previous)).
		Style("--year-h", fmt.Sprint(yearHue(i, last))).
		Aria("hidden", fmt.Sprint(i != t.selected)).
		Body(body...)
}

func (t *Timeline) renderRail(last int) app.UI {
	return app.Div().Class("timeline__rail").Body(
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
					app.Range(timelineEntries).Slice(t.renderDateDot),
				),
			),
		),
		app.Button().
			Class("timeline__nav timeline__nav--next").
			Aria("label", "Next event").
			Disabled(t.selected == last).
			OnClick(t.shiftBy(1)).
			Body(chevron(1)),
	)
}

func (t *Timeline) renderEvents() app.UI {
	return app.Div().
		Class("timeline__events").
		On("touchstart", t.onTouchStart).
		On("touchend", t.onTouchEnd).
		Body(
			app.Range(timelineEntries).Slice(t.renderEventPanel),
		)
}

func (t *Timeline) Render() app.UI {
	last := lastIndex()
	return app.Section().
		Class("timeline").
		Aria("label", "Career timeline").
		TabIndex(0).
		Style("--year-h", fmt.Sprint(yearHue(t.selected, last))).
		OnKeyDown(t.onKeyDown).
		Body(t.renderRail(last), t.renderEvents())
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

var keyStep = map[string]int{
	"ArrowLeft":  -1,
	"ArrowRight": 1,
}

func (t *Timeline) onKeyDown(ctx app.Context, e app.Event) {
	step, ok := keyStep[e.Get("key").String()]
	if !ok {
		return
	}
	e.PreventDefault()
	t.move(ctx, t.selected+step)
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

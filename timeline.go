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

// The active dot beats; giving the animation its own class keeps it off the
// other fourteen, which would otherwise all repaint on every frame.
func dotClass(i, selected int) string {
	if i == selected {
		return "timeline__dot timeline__dot--beat"
	}
	return "timeline__dot"
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

// showsYear is false for the second and later milestones of a year, so the rail
// reads 2015 2017 2018 2019 2020 2021 2022 2024 2025 rather than repeating a
// year under every dot. The dots themselves all stay.
func showsYear(i int) bool {
	return i == 0 || timelineEntries[i].Year != timelineEntries[i-1].Year
}

func yearLabel(i int) string {
	if showsYear(i) {
		return timelineEntries[i].Year
	}
	return ""
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
	ctx.Defer(jumpToActiveDate)
}

// honorStar reserves its line whether or not the milestone earned a star, so
// every year label sits on the same baseline across the rail.
func honorStar(e TimelineEntry) app.UI {
	slot := app.Span().Class("timeline__date-star").Aria("hidden", "true")
	if e.Quote == "" {
		return slot
	}
	return slot.Body(starIcon(11))
}

func (t *Timeline) renderDateDot(i int) app.UI {
	last := lastIndex()
	e := timelineEntries[i]
	return app.Li().Body(
		app.Button().
			Type("button").
			Class(dateClass(i, t.selected)).
			Style("--year-h", fmt.Sprint(yearHue(i, last))).
			Aria("label", e.Date).
			OnClick(t.selectEvent(i)).
			Body(
				honorStar(e),
				app.Span().Class("timeline__date-year").Text(yearLabel(i)),
				app.Span().Class(dotClass(i, t.selected)).Aria("hidden", "true"),
			),
	)
}

func (t *Timeline) renderEventPanel(i int) app.UI {
	last := lastIndex()
	e := timelineEntries[i]
	body := []app.UI{
		app.P().Class("timeline__event-date").Body(
			app.Time().Attr("datetime", e.Time).Text(e.Date),
		),
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

func (t *Timeline) renderHeader() app.UI {
	last := lastIndex()
	return app.Div().Class("timeline__head").Body(
		app.H2().Class("section-label").Text("Career"),
		app.Div().Class("timeline__nav").Body(
			app.Button().
				Type("button").
				Class("timeline__step").
				Aria("label", "Previous milestone").
				Disabled(t.selected == 0).
				OnClick(t.shiftBy(-1)).
				Body(chevron(-1)),
			app.Button().
				Type("button").
				Class("timeline__step").
				Aria("label", "Next milestone").
				Disabled(t.selected == last).
				OnClick(t.shiftBy(1)).
				Body(chevron(1)),
		),
	)
}

// --dots lets the CSS inset the rail line by half a column without knowing how
// many milestones there are, so adding one does not leave the line overshooting
// the first and last dots.
func (t *Timeline) renderRail(last int) app.UI {
	return app.Div().Class("timeline__rail").Body(
		app.Div().
			Class("timeline__strip").
			Style("--dots", fmt.Sprint(len(timelineEntries))).
			Body(
				app.Span().Class("timeline__line").Aria("hidden", "true"),
				app.Span().
					Class("timeline__fill").
					Aria("hidden", "true").
					Style("clip-path", fillClip(t.selected, last)),
				app.Ol().Class("timeline__dates").Body(
					app.Range(timelineEntries).Slice(t.renderDateDot),
				),
			),
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

// renderStaticEntry is the prerendered form of a milestone: a dated article in
// document order, with nothing hidden.
func renderStaticEntry(i int) app.UI {
	last := lastIndex()
	e := timelineEntries[i]
	body := []app.UI{
		app.P().Class("timeline__event-date").Body(
			app.Time().Attr("datetime", e.Time).Text(e.Date),
		),
		app.P().Class("timeline__event-body").Text(e.Body),
	}
	if e.Quote != "" {
		body = append(body, app.Blockquote().Class("timeline__event-quote").Text(e.Quote))
	}
	return app.Li().Body(
		app.Article().
			Class("timeline__entry").
			Style("--year-h", fmt.Sprint(yearHue(i, last))).
			Body(body...),
	)
}

// The interactive rail shows one milestone at a time, which is right for a
// reader and wrong for anything that only sees the HTML: fourteen of fifteen
// panels would ship aria-hidden. Prerendering therefore emits the whole career
// as an ordered list with machine-readable dates, and WASM replaces it with the
// rail once it boots.
func (t *Timeline) renderStatic() app.UI {
	return app.Section().
		Class("timeline timeline--static").
		Aria("label", "Career timeline").
		Body(
			app.H2().Class("section-label").Text("Career"),
			app.Ol().Class("timeline__list").Body(
				app.Range(timelineEntries).Slice(renderStaticEntry),
			),
		)
}

func (t *Timeline) Render() app.UI {
	if app.IsServer {
		return t.renderStatic()
	}

	last := lastIndex()
	return app.Section().
		Class("timeline").
		Aria("label", "Career timeline").
		TabIndex(0).
		Style("--year-h", fmt.Sprint(yearHue(t.selected, last))).
		OnKeyDown(t.onKeyDown).
		Body(
			t.renderHeader(),
			t.renderRail(last),
			t.renderEvents(),
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

func (t *Timeline) move(ctx app.Context, i int) {
	t.moveTo(i)
	ctx.Defer(glideToActiveDate)
}

func (t *Timeline) moveTo(i int) {
	if i < 0 || i >= len(timelineEntries) || i == t.selected {
		return
	}
	t.previous = t.selected
	t.selected = i
}

// scrollIntoView walks every scrollable ancestor and, on a narrow screen,
// settles somewhere short of the selected year. Scrolling the rail itself by
// the dot's own offset is both exact and cheaper.
func centerActiveDate(behavior string) {
	doc := app.Window().Get("document")
	active := doc.Call("querySelector", ".timeline__date--active")
	rail := doc.Call("querySelector", ".timeline__rail")
	if !active.Truthy() || !rail.Truthy() {
		return
	}
	// offsetLeft is measured from .timeline__strip, which is the rail's scroll
	// origin, so it is already in scroll coordinates.
	left := active.Get("offsetLeft").Float() +
		active.Get("offsetWidth").Float()/2 -
		rail.Get("clientWidth").Float()/2
	rail.Call("scrollTo", map[string]any{"left": left, "behavior": behavior})
}

func prefersReducedMotion() bool {
	q := app.Window().Call("matchMedia", "(prefers-reduced-motion: reduce)")
	return q.Truthy() && q.Get("matches").Bool()
}

// On mount the rail should already be showing today. Gliding there from 2015
// would be a ten-year scroll nobody asked for — and under a virtual clock or
// a stalled compositor the animation may simply never land.
func jumpToActiveDate(ctx app.Context) {
	centerActiveDate("instant")
}

// Moving between years is a deliberate act, so the rail follows the eye —
// unless the reader has asked the machine to stop moving things.
func glideToActiveDate(ctx app.Context) {
	if prefersReducedMotion() {
		centerActiveDate("instant")
		return
	}
	centerActiveDate("smooth")
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

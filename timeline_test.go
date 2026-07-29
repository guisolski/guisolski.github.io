package main

import "testing"

func TestMoveTo(t *testing.T) {
	last := len(timelineEntries) - 1
	tests := []struct {
		name         string
		selected     int
		target       int
		wantSelected int
		wantPrevious int
	}{
		{name: "forward", selected: 0, target: 1, wantSelected: 1, wantPrevious: 0},
		{name: "backward", selected: 3, target: 2, wantSelected: 2, wantPrevious: 3},
		{name: "jump to date", selected: 0, target: last, wantSelected: last, wantPrevious: 0},
		{name: "below range is ignored", selected: 0, target: -1, wantSelected: 0, wantPrevious: 0},
		{name: "above range is ignored", selected: last, target: last + 1, wantSelected: last, wantPrevious: 0},
		{name: "same index is ignored", selected: 2, target: 2, wantSelected: 2, wantPrevious: 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tl := &Timeline{selected: tc.selected}
			tl.moveTo(tc.target)
			if tl.selected != tc.wantSelected || tl.previous != tc.wantPrevious {
				t.Errorf("moveTo(%d) from %d = (selected %d, previous %d), want (%d, %d)",
					tc.target, tc.selected, tl.selected, tl.previous, tc.wantSelected, tc.wantPrevious)
			}
		})
	}
}

func TestSwipeStep(t *testing.T) {
	tests := []struct {
		name  string
		delta float64
		want  int
	}{
		{name: "swipe right goes back", delta: swipeThreshold + 1, want: -1},
		{name: "swipe left goes forward", delta: -swipeThreshold - 1, want: 1},
		{name: "small right drag is ignored", delta: swipeThreshold, want: 0},
		{name: "small left drag is ignored", delta: -swipeThreshold, want: 0},
		{name: "tap is ignored", delta: 0, want: 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := swipeStep(tc.delta); got != tc.want {
				t.Errorf("swipeStep(%v) = %d, want %d", tc.delta, got, tc.want)
			}
		})
	}
}

func TestDateClass(t *testing.T) {
	tests := []struct {
		name     string
		i        int
		selected int
		want     string
	}{
		{name: "selected dot", i: 2, selected: 2, want: "timeline__date timeline__date--active"},
		{name: "passed dot", i: 1, selected: 2, want: "timeline__date timeline__date--passed"},
		{name: "upcoming dot", i: 3, selected: 2, want: "timeline__date"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := dateClass(tc.i, tc.selected); got != tc.want {
				t.Errorf("dateClass(%d, %d) = %q, want %q", tc.i, tc.selected, got, tc.want)
			}
		})
	}
}

func TestEventClass(t *testing.T) {
	tests := []struct {
		name     string
		i        int
		selected int
		previous int
		want     string
	}{
		{name: "hidden panel", i: 0, selected: 1, previous: 0, want: "timeline__event"},
		{
			name: "moving forward enters from the right",
			i:    2, selected: 2, previous: 1,
			want: "timeline__event timeline__event--selected timeline__event--enter-right",
		},
		{
			name: "moving backward enters from the left",
			i:    1, selected: 1, previous: 2,
			want: "timeline__event timeline__event--selected timeline__event--enter-left",
		},
		{
			name: "initial render enters from the right",
			i:    0, selected: 0, previous: 0,
			want: "timeline__event timeline__event--selected timeline__event--enter-right",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := eventClass(tc.i, tc.selected, tc.previous); got != tc.want {
				t.Errorf("eventClass(%d, %d, %d) = %q, want %q", tc.i, tc.selected, tc.previous, got, tc.want)
			}
		})
	}
}

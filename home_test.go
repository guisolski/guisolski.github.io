package main

import "testing"

func TestPageClass(t *testing.T) {
	tests := []struct {
		name    string
		leaving bool
		want    string
	}{
		{"resting", false, "page"},
		{"leaving", true, "page page--leaving"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := pageClass(tc.leaving); got != tc.want {
				t.Errorf("pageClass(%v) = %q, want %q", tc.leaving, got, tc.want)
			}
		})
	}
}

func TestIntroEyebrow(t *testing.T) {
	const want = "Backend software engineer · Go · Mercado Libre"
	if introEyebrow != want {
		t.Errorf("introEyebrow = %q, want %q", introEyebrow, want)
	}
}

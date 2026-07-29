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

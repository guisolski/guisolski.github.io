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

// The eyebrow answers the two questions a recruiter opens with — what he
// does and whether he can work from where he is — before the name is read.
func TestIntroEyebrow(t *testing.T) {
	const want = "Backend engineer · Go · Remote, UTC−3"
	if introEyebrow != want {
		t.Errorf("introEyebrow = %q, want %q", introEyebrow, want)
	}
}

package main

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestTerrainPoints(t *testing.T) {
	tests := []struct {
		name          string
		width, height float64
	}{
		{name: "phone portrait", width: 375, height: 667},
		{name: "laptop", width: 1440, height: 900},
		{name: "exact power of two", width: 1024, height: 768},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rng := rand.New(rand.NewSource(1))
			points := terrainPoints(tc.width, tc.height, rng)

			power := 1 << int(math.Ceil(math.Log2(tc.width)))
			if len(points) != power+1 {
				t.Fatalf("len(points) = %d, want %d", len(points), power+1)
			}
			if points[0] != tc.height/2 {
				t.Errorf("left edge = %v, want %v", points[0], tc.height/2)
			}
			if points[power] != tc.height {
				t.Errorf("right edge = %v, want %v", points[power], tc.height)
			}
			for i, p := range points {
				if p < 0 || p > tc.height+125 {
					t.Fatalf("points[%d] = %v out of range [0, %v]", i, p, tc.height+125)
				}
			}
		})
	}
}

func TestTerrainPointsDeterministic(t *testing.T) {
	a := terrainPoints(800, 600, rand.New(rand.NewSource(7)))
	b := terrainPoints(800, 600, rand.New(rand.NewSource(7)))
	for i := range a {
		if a[i] != b[i] {
			t.Fatalf("same seed produced different profiles at index %d: %v != %v", i, a[i], b[i])
		}
	}
}

func TestBranchEnd(t *testing.T) {
	tests := []struct {
		name         string
		x, y, angle  float64
		depth        int
		armLength    float64
		wantX, wantY float64
	}{
		{name: "straight up", x: 100, y: 200, angle: -90, depth: 5, armLength: 10, wantX: 100, wantY: 150},
		{name: "to the right", x: 100, y: 200, angle: 0, depth: 5, armLength: 10, wantX: 150, wantY: 200},
		{name: "zero depth stays put", x: 100, y: 200, angle: -45, depth: 0, armLength: 10, wantX: 100, wantY: 200},
	}

	const epsilon = 1e-9
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotX, gotY := branchEnd(tc.x, tc.y, tc.angle, tc.depth, tc.armLength)
			if math.Abs(gotX-tc.wantX) > epsilon || math.Abs(gotY-tc.wantY) > epsilon {
				t.Errorf("branchEnd = (%v, %v), want (%v, %v)", gotX, gotY, tc.wantX, tc.wantY)
			}
		})
	}
}

func TestTreeDepth(t *testing.T) {
	tests := []struct {
		name   string
		height float64
		want   int
	}{
		{name: "short viewport clamps to minimum", height: 300, want: 5},
		{name: "typical viewport", height: 800, want: 8},
		{name: "tall viewport clamps to maximum", height: 2000, want: 10},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := treeDepth(tc.height); got != tc.want {
				t.Errorf("treeDepth(%v) = %d, want %d", tc.height, got, tc.want)
			}
		})
	}
}

func TestAdvanceStar(t *testing.T) {
	const width, height = 800.0, 600.0
	tests := []struct {
		name        string
		star        star
		wantVisible bool
		wantRespawn bool
	}{
		{name: "drifts left and stays visible", star: star{x: 100, y: 50, speed: 0.1}, wantVisible: true},
		{name: "respawns at the right edge", star: star{x: 0.01, y: 50, speed: 10}, wantRespawn: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rng := rand.New(rand.NewSource(1))
			got, visible := advanceStar(tc.star, width, height, rng)
			if visible != tc.wantVisible {
				t.Fatalf("visible = %v, want %v", visible, tc.wantVisible)
			}
			if tc.wantRespawn && got.x != width {
				t.Errorf("respawned star x = %v, want %v", got.x, width)
			}
			if !tc.wantRespawn && got.x != tc.star.x-tc.star.speed {
				t.Errorf("moved star x = %v, want %v", got.x, tc.star.x-tc.star.speed)
			}
		})
	}
}

func TestAdvanceShootingStar(t *testing.T) {
	const width, height = 800.0, 600.0
	now := time.Unix(1000, 0)
	tests := []struct {
		name        string
		star        shootingStar
		wantVisible bool
		wantActive  bool
	}{
		{
			name:       "waits until its scheduled time",
			star:       shootingStar{waitUntil: now.Add(time.Second)},
			wantActive: false,
		},
		{
			name:       "wakes up after its scheduled time",
			star:       shootingStar{waitUntil: now.Add(-time.Second)},
			wantActive: true,
		},
		{
			name:        "streaks down-left while on canvas",
			star:        shootingStar{x: 400, y: 100, speed: 8, active: true},
			wantVisible: true,
			wantActive:  true,
		},
		{
			name:       "resets after leaving the canvas",
			star:       shootingStar{x: 2, y: 100, speed: 8, active: true},
			wantActive: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rng := rand.New(rand.NewSource(1))
			got, visible := advanceShootingStar(tc.star, width, height, now, rng)
			if visible != tc.wantVisible {
				t.Fatalf("visible = %v, want %v", visible, tc.wantVisible)
			}
			if got.active != tc.wantActive {
				t.Errorf("active = %v, want %v", got.active, tc.wantActive)
			}
		})
	}
}

func TestRandInt(t *testing.T) {
	tests := []struct {
		name     string
		min, max int
	}{
		{name: "branch arm range", min: 0, max: 20},
		{name: "branch angle range", min: 15, max: 20},
		{name: "single value", min: 3, max: 3},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rng := rand.New(rand.NewSource(1))
			for range 1000 {
				if got := randInt(rng, tc.min, tc.max); got < tc.min || got > tc.max {
					t.Fatalf("randInt(%d, %d) = %d out of range", tc.min, tc.max, got)
				}
			}
		})
	}
}

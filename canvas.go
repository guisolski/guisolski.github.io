package main

import (
	"math"
	"math/rand"
	"time"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

const nightSky = "#05004c"

type star struct {
	x, y, size, speed float64
}

func newStar(width, height float64, rng *rand.Rand) star {
	return star{
		x:     rng.Float64() * width,
		y:     rng.Float64() * height,
		size:  rng.Float64() * 2,
		speed: rng.Float64() * 0.1,
	}
}

func advanceStar(s star, width, height float64, rng *rand.Rand) (star, bool) {
	s.x -= s.speed
	if s.x < 0 {
		s = newStar(width, height, rng)
		s.x = width
		return s, false
	}
	return s, true
}

type shootingStar struct {
	x, y, length, speed, size float64
	active                    bool
	waitUntil                 time.Time
}

func newShootingStar(width float64, now time.Time, rng *rand.Rand) shootingStar {
	return shootingStar{
		x:         rng.Float64() * width,
		y:         0,
		length:    rng.Float64()*80 + 10,
		speed:     rng.Float64()*10 + 6,
		size:      rng.Float64() + 0.1,
		waitUntil: now.Add(time.Duration(rng.Float64()*3000+500) * time.Millisecond),
	}
}

func advanceShootingStar(s shootingStar, width, height float64, now time.Time, rng *rand.Rand) (shootingStar, bool) {
	if !s.active {
		if now.After(s.waitUntil) {
			s.active = true
		}
		return s, false
	}
	s.x -= s.speed
	s.y += s.speed
	if s.x < 0 || s.y >= height {
		return newShootingStar(width, now, rng), false
	}
	return s, true
}

type starfield struct {
	ctx2d         app.Value
	rng           *rand.Rand
	width, height float64
	stars         []star
	shooting      []shootingStar
}

func newStarfield(ctx2d app.Value, width, height float64, rng *rand.Rand) *starfield {
	f := &starfield{ctx2d: ctx2d, rng: rng, width: width, height: height}
	f.stars = make([]star, int(height))
	for i := range f.stars {
		f.stars[i] = newStar(width, height, rng)
	}
	f.shooting = make([]shootingStar, 4)
	for i := range f.shooting {
		f.shooting[i] = newShootingStar(width, time.Now(), rng)
	}
	return f
}

func (f *starfield) step() {
	f.ctx2d.Set("fillStyle", nightSky)
	f.ctx2d.Call("fillRect", 0, 0, f.width, f.height)
	f.ctx2d.Set("fillStyle", "#ffffff")
	f.ctx2d.Set("strokeStyle", "#ffffff")

	for i := range f.stars {
		next, visible := advanceStar(f.stars[i], f.width, f.height, f.rng)
		f.stars[i] = next
		if visible {
			f.ctx2d.Call("fillRect", next.x, next.y, next.size, next.size)
		}
	}

	now := time.Now()
	for i := range f.shooting {
		next, visible := advanceShootingStar(f.shooting[i], f.width, f.height, now, f.rng)
		f.shooting[i] = next
		if !visible {
			continue
		}
		f.ctx2d.Set("lineWidth", next.size)
		strokeLine(f.ctx2d, next.x, next.y, next.x+next.length, next.y-next.length)
	}
}

func strokeLine(ctx2d app.Value, x1, y1, x2, y2 float64) {
	ctx2d.Call("beginPath")
	ctx2d.Call("moveTo", x1, y1)
	ctx2d.Call("lineTo", x2, y2)
	ctx2d.Call("stroke")
}

func terrainPoints(width, height float64, rng *rand.Rand) []float64 {
	displacement := 50.0
	power := 1 << int(math.Ceil(math.Log2(width)))
	points := make([]float64, power+1)
	points[0] = height - height/2
	points[power] = height

	for i := 1; i < power; i *= 2 {
		step := power / i
		for j := step / 2; j < power; j += step {
			mid := (points[j-step/2] + points[j+step/2]) / 2
			points[j] = mid + math.Floor(rng.Float64()*-displacement+displacement)
		}
		displacement *= 0.6
	}
	return points
}

func drawTerrain(ctx2d app.Value, width, height float64, rng *rand.Rand) {
	points := terrainPoints(width, height, rng)
	ctx2d.Set("fillStyle", "#000000")
	ctx2d.Call("beginPath")
	ctx2d.Call("moveTo", 0, points[0])
	for i := 1; i <= int(width) && i < len(points); i++ {
		ctx2d.Call("lineTo", i, points[i])
	}
	ctx2d.Call("lineTo", width, height)
	ctx2d.Call("lineTo", 0, height)
	ctx2d.Call("closePath")
	ctx2d.Call("fill")
}

func treeDepth(height float64) int {
	return min(max(int(height*0.01), 5), 10)
}

func branchEnd(x, y, angle float64, depth int, armLength float64) (float64, float64) {
	rad := angle * math.Pi / 180
	return x + math.Cos(rad)*float64(depth)*armLength,
		y + math.Sin(rad)*float64(depth)*armLength
}

func drawTree(ctx2d app.Value, x, y, height float64, rng *rand.Rand) {
	branch(ctx2d, x, y, -90, treeDepth(height), rng)
}

func branch(ctx2d app.Value, x1, y1, angle float64, depth int, rng *rand.Rand) {
	if depth == 0 {
		ctx2d.Set("fillStyle", "#ffffff")
		ctx2d.Call("beginPath")
		ctx2d.Call("arc", x1, y1, float64(randInt(rng, 0, 10)), 0, 2*math.Pi)
		ctx2d.Call("fill")
		return
	}

	x2, y2 := branchEnd(x1, y1, angle, depth, float64(randInt(rng, 0, 20)))

	ctx2d.Set("strokeStyle", "#8b4513")
	ctx2d.Set("lineWidth", float64(depth)*1.5)
	strokeLine(ctx2d, x1, y1, x2, y2)

	branch(ctx2d, x2, y2, angle-float64(randInt(rng, 15, 20)), depth-1, rng)
	branch(ctx2d, x2, y2, angle+float64(randInt(rng, 15, 20)), depth-1, rng)
}

func randInt(rng *rand.Rand, min, max int) int {
	return min + rng.Intn(max+1-min)
}

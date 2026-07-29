# guisolski.github.io

My personal CV — https://guisolski.github.io

Written in **Go**, compiled to **WebAssembly** with
[go-app](https://github.com/maxence-charriere/go-app). No JavaScript
frameworks, no external requests: the whole UI (reactive career timeline,
cards, and a hidden canvas night scene) runs as a Go program in the browser.

## Development

```sh
make help    # list all targets
make serve   # build wasm + dev server on :8000
make test    # go vet + table-driven test suite
make dist    # production static site in dist/
```

Deployment is automated: every push to `master` runs
`.github/workflows/deploy.yml`, which builds the WASM binary, generates the
static site with prerendered HTML, and publishes it to GitHub Pages.

## Design notes

- **Determinism via explicit inputs (`canvas.go`)**: `advanceStar`,
  `advanceShootingStar`, and `terrainPoints` take `*rand.Rand`/`time.Time` as
  explicit parameters instead of calling `rand.Float64()`/`time.Now()`
  internally, keeping them pure and letting callers (and tests) control
  randomness and time deterministically.
- **Go owns the simulation, JS only draws (`canvas.go`)**: starfield/terrain/
  tree state advances entirely in Go through pure functions; only the final
  draw calls cross into JS.
- **`treeDepth` clamp (`canvas.go`)**: recursion depth is clamped to `[5, 10]`
  so branch count stays sane on very short or very tall viewports.
- **`/relax` staged reveal (`relax.go`)**: the hidden night scene reveals its
  layers (starfield, terrain, tree) in stages, matching the pacing of the
  site's original, pre-Go/WASM version.
- **Timeline "dawn" color story (`timeline.go`)**: the accent hue runs from
  night indigo (`dawnStartHue = 245`) at the first year, through dusk rose, to
  star amber (35°) at the most recent year — mirroring the night-to-day story
  of the hidden `/relax` scene.
- **Timeline opens on the latest year (`timeline.go`)**: `OnInit` selects the
  most recent entry so the current role reads first and the dawn line arrives
  fully revealed; it runs before the first render on both prerender and
  in-browser (go-app lifecycle).
- **Timeline keeps the active date in view (`timeline.go`)**: after a
  selection change, `move` defers a scroll-into-view of the active date dot so
  it stays visible when the rail overflows on narrow screens.
- **Timeline event-panel slide direction (`timeline.go`)**: the selected panel
  animates in from the side it was approached from.
- **Inline SVG icons (`cards.go`)**: icons render as inline `<svg>` markup
  rather than an icon font, so the site makes no external requests.
- **Terrain-drift test tolerance (`canvas_test.go`)**: midpoint displacement
  can drift the profile by at most ~50+30+18+...≈125px total, hence
  `TestTerrainPoints`'s `height+125` bound.
- **Timeline entries must stay chronological (`content.go`)**: `timelineEntries`
  is expected to remain chronologically ordered (enforced by
  `TestTimelineEntriesChronological`); add new milestones as one-line appends
  at the end.
- **`entryTime` test helper (`content_test.go`)**: bare-year dates (`"2023"`)
  are treated as sorting at the end of that year, so they don't spuriously
  sort before same-year entries that include a month.

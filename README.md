# guisolski.github.io

My personal CV — https://guisolski.github.io

Written in **Go**, compiled to **WebAssembly** with
[go-app](https://github.com/maxence-charriere/go-app). No JavaScript
frameworks, no external requests: the whole UI (interactive career rail, focus
areas, cards, and a hidden canvas night scene) runs as a Go program in the
browser. The machine-readable layer — prerendered HTML with `<time datetime>`
milestones, schema.org JSON-LD, `llms.txt`, `sitemap.xml` — is generated from
the same `content.go` the page renders from.

English CV download: [`assets/pdf/cv.pdf`](assets/pdf/cv.pdf) (moderncv).
`assets/pdf/resume.pdf` is kept as an identical copy so older bookmarks still
resolve. Portuguese: [`assets/pdf/Curriculum/portugues.pdf`](assets/pdf/Curriculum/portugues.pdf).

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

Files under `static/` (for example `robots.txt`) are copied to the site root by
`make dist`. `llms.txt` and `sitemap.xml` are not files at all — they are
generated from `content.go` at build time by `llms.go`, so there is no second
copy of the career facts to drift out of date. `make og` regenerates the
1200x630 share card at `assets/images/og.png`.

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
- **Timeline keeps the active date in view (`timeline.go`)**: the rail
  overflows on narrow screens, so `centerActiveDate` scrolls the rail itself by
  the active dot's offset. `scrollIntoView` walks every scrollable ancestor and
  settles short of the mark; and the initial centring is `"instant"`, because a
  smooth ten-year glide on mount is both unasked-for and, under a stalled
  compositor, liable never to land.
- **Timeline renders differently on the server (`timeline.go`)**: the rail shows
  one milestone at a time, which would ship fourteen of fifteen panels as
  `aria-hidden` to anything that only reads the HTML. Under `app.IsServer` the
  component renders the whole career as an `<ol>` of `<article>`s with
  `<time datetime>`, and WASM swaps in the rail once it boots. go-app replaces
  the body wholesale on load, so the two forms need not match.
- **Repeated years are unlabelled (`timeline.go`)**: `showsYear` blanks the year
  under the second and later dots of the same year, so the rail reads as a run
  of years rather than "2022 2022 2022 2022". Every dot stays.
- **Structured data mirrors the page (`jsonld.go`)**: the schema.org
  `ProfilePage`/`Person` block is built from the same constants the page
  renders. `json.Marshal` escapes `<`, `>` and `&`, so the payload cannot close
  the `<script>` element it is injected into.
- **`Domain` and `Icon` on the handler (`main.go`)**: without `Domain`, go-app
  emits `og:url` and `og:image` as the literal string `https://`; without
  `Icon.SVG`, the favicon is fetched from `raw.githubusercontent.com` — which
  is both the wrong logo and the site's only external request.
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
- **`entryTime` test helper (`content_test.go`)**: bare-year dates (`"2024"`)
  are treated as sorting at the end of that year, so they don't spuriously
  sort before same-year entries that include a month.
- **Colour is information, not decoration (`web/app.css`)**: hue always encodes
  a position in time. The career rail walks the dawn spectrum from night indigo
  (2015) to star amber (2025), and the focus areas and cards below take their
  stop from the same walk, so scrolling the page warms it the way scrubbing the
  rail does. Nothing is tinted for its own sake.
- **Dot contrast (`web/app.css`)**: the rail dot's border carries the contrast
  (>=3.5:1 at every hue) and its fill carries the colour, so the bright amber
  the design wants stays a legible UI component.

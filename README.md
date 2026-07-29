# guisolski.github.io

My personal CV — https://guisolski.github.io

Written in **Go**, compiled to **WebAssembly** with
[go-app](https://github.com/maxence-charriere/go-app). No JavaScript
frameworks, no external requests: the whole UI (reactive career timeline,
cards, and a hidden canvas night scene) runs as a Go program in the browser.

## Development

```sh
GOOS=js GOARCH=wasm go build -o web/app.wasm .   # build the browser binary
go run . -serve                                   # dev server on :8000
```

## Build the static site

```sh
GOOS=js GOARCH=wasm go build -trimpath -ldflags="-s -w" -o web/app.wasm .
go run . -out dist
cp -r assets dist/assets && cp -r static/. dist/
```

Deployment is automated: every push to `master` runs
`.github/workflows/deploy.yml`, which builds the WASM binary, generates the
static site with prerendered HTML, and publishes it to GitHub Pages.

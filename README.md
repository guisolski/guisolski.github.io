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

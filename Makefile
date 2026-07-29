# guisolski.github.io — Go + WebAssembly personal site
#
# Targets are self-documenting: lines with "##" feed `make help`.

.DEFAULT_GOAL := help

GO_WASM   := GOOS=js GOARCH=wasm go build
LDFLAGS   := -trimpath -ldflags="-s -w"
WASM_OUT  := web/app.wasm
DIST      := dist

.PHONY: help test wasm serve dist opt clean

help: ## Show this help
	@echo "Usage: make <target>"
	@echo ""
	@awk 'BEGIN {FS = ":.*##"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-8s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

test: ## Run go vet and the test suite
	go vet ./...
	go test ./...

wasm: ## Build the WebAssembly binary (development)
	$(GO_WASM) -o $(WASM_OUT) .

serve: wasm ## Run the dev server on http://localhost:8000
	go run . -serve

dist: test ## Build the production static site into dist/
	$(GO_WASM) $(LDFLAGS) -o $(WASM_OUT) .
	rm -rf $(DIST)
	go run . -out $(DIST)
	cp -r web/. $(DIST)/web/
	cp -r assets $(DIST)/assets
	cp -r static/. $(DIST)/
	cp $(DIST)/index.html $(DIST)/404.html

opt: ## Shrink web/app.wasm with wasm-opt (requires binaryen)
	wasm-opt -Oz --enable-bulk-memory --enable-sign-ext \
		--enable-nontrapping-float-to-int -o $(WASM_OUT).opt $(WASM_OUT)
	mv $(WASM_OUT).opt $(WASM_OUT)
	ls -lh $(WASM_OUT)

clean: ## Remove build artifacts (dist/ and web/app.wasm)
	rm -rf $(DIST) $(WASM_OUT)

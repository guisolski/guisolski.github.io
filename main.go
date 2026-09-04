package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type route struct {
	Path string
	New  func() app.Composer
}

var routes = []route{
	{"/", newHome},
	{"/relax", newRelax},
}

func newHome() app.Composer {
	return &Home{}
}

func newRelax() app.Composer {
	return &Relax{}
}

func registerRoutes() {
	for _, r := range routes {
		app.Route(r.Path, r.New)
	}
}

func newHandler() *app.Handler {
	return &app.Handler{
		Name:        personName,
		ShortName:   "Solski",
		Title:       personName + " — " + personJobTitle,
		Description: personTagline + " Career timeline, focus areas, contact, and CV.",
		Author:      personName,
		Lang:        "en",

		// Without Domain, go-app emits og:url and og:image as the literal
		// string "https://" and every share preview resolves to nothing.
		Domain:   "guisolski.github.io",
		Image:    "/assets/images/og.png",
		Keywords: []string{"Go", "backend engineer", "observability", "Curitiba", "remote"},

		// SVG defaults to go-app's own icon on raw.githubusercontent.com,
		// which is both the wrong logo and the site's only external request.
		Icon: app.Icon{
			Default: "/web/icon-192.png",
			Large:   "/web/icon-512.png",
			SVG:     "/web/icon.svg",
		},

		RawHeaders:      []string{personJSONLD()},
		BackgroundColor: "#ffffff",
		ThemeColor:      "#ffffff",
		LoadingLabel:    "Loading… {progress}%",
		Styles:          []string{"/web/app.css"},
		Version:         os.Getenv("GITHUB_SHA"),
	}
}

func parseFlags() (bool, string) {
	serve := flag.Bool("serve", false, "run local dev server instead of generating the static site")
	out := flag.String("out", "dist", "output directory for the generated static site")
	flag.Parse()
	return *serve, *out
}

func serveServiceWorker(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/service-worker.js")
}

func serveStaticRootFile(name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/"+name)
	}
}

func serveText(contentType string, body func() string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", contentType)
		if _, err := w.Write([]byte(body())); err != nil {
			log.Printf("serving %s: %v", r.URL.Path, err)
		}
	}
}

func today() string {
	return time.Now().UTC().Format("2006-01-02")
}

func newDevMux(h *app.Handler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.FileServer(http.Dir(".")))
	mux.HandleFunc("/service-worker.js", serveServiceWorker)
	mux.HandleFunc("/robots.txt", serveStaticRootFile("robots.txt"))
	mux.HandleFunc("/llms.txt", serveText("text/plain; charset=utf-8", llmsTxt))
	mux.HandleFunc("/sitemap.xml", serveText("application/xml; charset=utf-8", func() string {
		return sitemapXML(today())
	}))
	mux.Handle("/", h)
	return mux
}

func runDevServer(h *app.Handler) error {
	log.Println("dev server listening on http://localhost:8000")
	return http.ListenAndServe(":8000", newDevMux(h))
}

// generatedFiles are written into the static site after go-app has produced the
// HTML, so llms.txt and sitemap.xml are built from content.go rather than kept
// by hand alongside it.
func writeGeneratedFiles(out string) error {
	files := map[string]string{
		"llms.txt":    llmsTxt(),
		"sitemap.xml": sitemapXML(today()),
	}
	for name, body := range files {
		if err := os.WriteFile(filepath.Join(out, name), []byte(body), 0o644); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	registerRoutes()
	app.RunWhenOnBrowser()

	h := newHandler()
	serve, out := parseFlags()

	if serve {
		log.Fatal(runDevServer(h))
	}

	if err := app.GenerateStaticWebsite(out, h); err != nil {
		log.Fatal(err)
	}

	if err := writeGeneratedFiles(out); err != nil {
		log.Fatal(err)
	}
}

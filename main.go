package main

import (
	"flag"
	"log"
	"net/http"
	"os"

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
		Name:            "Guilherme Solski Alves",
		ShortName:       "Solski",
		Title:           "Guilherme Solski Alves — Software Developer",
		Description:     "Software developer at Mercado Libre. Career timeline, contact, and résumé.",
		Author:          "Guilherme Solski Alves",
		Lang:            "en",
		Icon:            app.Icon{Default: "/web/icon-192.png", Large: "/web/icon-512.png"},
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

func newDevMux(h *app.Handler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.FileServer(http.Dir(".")))
	mux.HandleFunc("/service-worker.js", serveServiceWorker)
	mux.Handle("/", h)
	return mux
}

func runDevServer(h *app.Handler) error {
	log.Println("dev server listening on http://localhost:8000")
	return http.ListenAndServe(":8000", newDevMux(h))
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
}

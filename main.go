package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

func main() {
	app.Route("/", func() app.Composer { return &Home{} })
	app.Route("/relax", func() app.Composer { return &Relax{} })
	app.RunWhenOnBrowser()

	h := &app.Handler{
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

	serve := flag.Bool("serve", false, "run local dev server instead of generating the static site")
	out := flag.String("out", "dist", "output directory for the generated static site")
	flag.Parse()

	if *serve {
		mux := http.NewServeMux()
		mux.Handle("/assets/", http.FileServer(http.Dir(".")))
		mux.HandleFunc("/service-worker.js", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "static/service-worker.js")
		})
		mux.Handle("/", h)
		log.Println("dev server listening on http://localhost:8000")
		log.Fatal(http.ListenAndServe(":8000", mux))
	}

	if err := app.GenerateStaticWebsite(*out, h); err != nil {
		log.Fatal(err)
	}
}

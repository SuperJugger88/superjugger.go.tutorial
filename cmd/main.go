package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"net/http"
	"superjugger.go.tutorial/internal/handlers"
	healthcheck "superjugger.go.tutorial/internal/utils"
	env "superjugger.go.tutorial/internal/utils/config"
	metrics "superjugger.go.tutorial/pkg"
)

func main() {
	mux := http.NewServeMux()

	prometheus.MustRegister(metrics.RequestDuration)

	fileServer := http.FileServer(http.Dir(env.StaticPath + "static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", handlers.GetHomePage)
	mux.HandleFunc("/snippet/view", handlers.GetSnippetByID)
	mux.HandleFunc("/snippet/create", handlers.CreateSnippet)
	mux.HandleFunc("/healthz", healthcheck.GetHealthCheckStatus)

	log.Printf("Starting server on :%s", env.Port)

	err := http.ListenAndServe(":"+env.Port, mux)
	if err != nil {
		log.Fatal(err)
	}
}

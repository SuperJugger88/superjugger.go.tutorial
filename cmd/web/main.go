package main

import (
	"log"
	"net/http"
	"os"
	"superjugger.go.tutorial/cmd/web/handlers"
)

var (
	port       = os.Getenv("APP_PORT")
	staticPath = os.Getenv("STATIC_PATH")
)

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(staticPath + "static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", handlers.GetHomePage)
	mux.HandleFunc("/snippet/view", handlers.GetSnippetByID)
	mux.HandleFunc("/snippet/create", handlers.CreateSnippet)

	log.Print("Starting server on :4000")

	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal(err)
	}
}

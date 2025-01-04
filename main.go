package main

import (
	"log"
	"net/http"
	"os"
	handlers "superjugger.go.tutorial/cmd/web"
)

var port = os.Getenv("APP_PORT")

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.GetHomePage)
	mux.HandleFunc("/snippet/create", handlers.CreateSnippet)
	mux.HandleFunc("/snippet", handlers.GetSnippetByID)

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		panic(err)
	}
}

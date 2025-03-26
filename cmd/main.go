package main

import (
	"log"
	"net/http"
	"superjugger.go.tutorial/internal/handlers"
	"superjugger.go.tutorial/internal/utils"
)

func main() {
	env := utils.SetEnvVars()

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(env.StaticPath + "static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", handlers.GetHomePage)
	mux.HandleFunc("/snippet/view", handlers.GetSnippetByID)
	mux.HandleFunc("/snippet/create", handlers.CreateSnippet)

	log.Printf("Starting server on :%s", env.Port)

	err := http.ListenAndServe(":"+env.Port, mux)
	if err != nil {
		log.Fatal(err)
	}
}

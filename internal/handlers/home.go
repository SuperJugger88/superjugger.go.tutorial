package handlers

import (
	"html/template"
	"log"
	"net/http"
	"superjugger.go.tutorial/internal/utils"
)

func GetHomePage(w http.ResponseWriter, r *http.Request) {
	env := utils.SetEnvVars()

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		env.StaticPath + "templates/base.tmpl",
		env.StaticPath + "templates/pages/home.tmpl",
		env.StaticPath + "templates/partials/nav.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

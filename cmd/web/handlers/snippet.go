package handlers

import (
	"log"
	"net/http"
	"strconv"
)

func CreateSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	_, err := w.Write([]byte("Create a new snippet..."))
	if err != nil {
		log.Println(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetSnippetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	if _, err := w.Write([]byte("Get snippet with id:" + strconv.Itoa(id))); err != nil {
		log.Println("Error displaying snippet:", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

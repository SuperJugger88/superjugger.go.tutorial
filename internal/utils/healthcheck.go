package utils

import (
	"log"
	"net/http"
)

func GetHealthCheckStatus(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/healthz" {
		http.NotFound(w, r)
		return
	}

	if _, err := w.Write([]byte("App is healthy")); err != nil {
		log.Println("Error displaying health check:", err.Error())
	}
	w.WriteHeader(http.StatusOK)
}

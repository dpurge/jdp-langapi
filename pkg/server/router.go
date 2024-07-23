package server

import (
	"net/http"

	"github.com/dpurge/lang-api/pkg/language"
)

type Handler struct{}

func loadRoutes(router *http.ServeMux) {
	handler := &language.Handler{}

	router.HandleFunc("GET /lang/{id}", handler.GetByID)
}

package language

import (
	"encoding/json"
	"log"
	"net/http"
)

type Handler struct{}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	log.Println("handling READ request - Method:", r.Method)
	language, exists := loadLanguages()[r.PathValue("id")]
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(language)
}

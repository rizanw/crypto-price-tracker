package module

import (
	"log"
	"net/http"
)

func (h *handler) Signup(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	log.Print("signup handler")
}

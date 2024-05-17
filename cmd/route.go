package main

import (
	hAuth "crypto-tracker/internal/handler/http/auth/module"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func newRoutes(uc Usecase) *mux.Router {
	router := mux.NewRouter()

	// auth routes
	handlerAuth := hAuth.New()
	router.HandleFunc("/signup", handlerAuth.Signup).Methods(http.MethodGet)

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		log.Println("server OK!")
		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodGet)

	return router
}

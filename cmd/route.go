package main

import (
	hAuth "crypto-tracker/internal/handler/http/auth/module"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func newRoutes(uc UseCase) *mux.Router {
	router := mux.NewRouter()

	// auth routes
	handlerAuth := hAuth.New(uc.Auth)
	router.HandleFunc("/signup", handlerAuth.SignUp).Methods(http.MethodPost)
	router.HandleFunc("/signin", handlerAuth.SignIn).Methods(http.MethodPost)

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		log.Println("server OK!")
		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodGet)

	return router
}

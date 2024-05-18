package main

import (
	"crypto-tracker/internal/common/middleware"
	"crypto-tracker/internal/config"
	hAuth "crypto-tracker/internal/handler/http/auth/module"
	hCoin "crypto-tracker/internal/handler/http/coin/module"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func newRoutes(uc UseCase, conf *config.JWTConfig) *mux.Router {
	router := mux.NewRouter()

	// auth routes
	handlerAuth := hAuth.New(uc.Auth)
	router.HandleFunc("/signup", handlerAuth.SignUp).Methods(http.MethodPost)
	router.HandleFunc("/signin", handlerAuth.SignIn).Methods(http.MethodPost)
	router.Handle("/signout", middleware.VerifyAuth(conf, http.HandlerFunc(handlerAuth.SignOut))).Methods(http.MethodGet)

	// coin routes
	handlerCoin := hCoin.New(uc.Coin)
	router.Handle("/coins", middleware.VerifyAuth(conf, http.HandlerFunc(handlerCoin.GetCoins))).Methods(http.MethodGet)
	router.Handle("/coin/add", middleware.VerifyAuth(conf, http.HandlerFunc(handlerCoin.AddCoin))).Methods(http.MethodPost)
	router.Handle("/coin/remove", middleware.VerifyAuth(conf, http.HandlerFunc(handlerCoin.RemoveCoin))).Methods(http.MethodPut)

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		log.Println("server OK!")
		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodGet)

	return router
}

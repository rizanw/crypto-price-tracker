package main

import (
	"crypto-tracker/internal/config"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	appName = "crypto-tracker"
)

func main() {
	var (
		address string
	)

	conf, err := config.New(appName)
	if err != nil {
		return
	}

	repo := newRepo(conf)
	uc := newUseCase(conf, repo)
	router := newRoutes(uc, &conf.JWT)

	address = fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port)
	srv := http.Server{
		Addr:         address,
		ReadTimeout:  time.Duration(conf.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(conf.Server.WriteTimeout) * time.Second,
		Handler:      router,
	}

	log.Println("app starting on ", address)
	err = srv.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

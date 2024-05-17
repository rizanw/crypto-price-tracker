package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

func main() {
	repo := newRepo(dbPath)
	uc := newUseCase(repo)
	router := newRoutes(uc)

	srv := http.Server{
		Addr:         address,
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
		Handler:      router,
	}

	err := srv.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

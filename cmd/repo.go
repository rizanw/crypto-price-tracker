package main

import (
	rCoincapHttp "crypto-tracker/internal/repo/coincap/http"
	coincapHttp "crypto-tracker/internal/repo/coincap/http/module"
	rDB "crypto-tracker/internal/repo/sqlite"
	sqlite "crypto-tracker/internal/repo/sqlite/module"
	"log"
)

type Repo struct {
	db          rDB.Sqlite
	coincapHttp rCoincapHttp.Repo
}

func newRepo(dbfile, coincapURL string) *Repo {
	db, err := sqlite.New(dbfile)
	if err != nil {
		log.Println("!Error init db:", err)
		return nil
	}

	return &Repo{
		db:          db,
		coincapHttp: coincapHttp.New(coincapURL),
	}
}

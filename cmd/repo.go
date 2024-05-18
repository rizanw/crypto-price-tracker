package main

import (
	"crypto-tracker/internal/config"
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

func newRepo(conf *config.Config) *Repo {
	db, err := sqlite.New(conf.Database)
	if err != nil {
		log.Println("!Error init db:", err)
		return nil
	}

	return &Repo{
		db:          db,
		coincapHttp: coincapHttp.New(conf.HTTPs.CoinCap),
	}
}

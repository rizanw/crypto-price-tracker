package main

import (
	rDB "crypto-tracker/internal/repo/sqlite"
	sqlite "crypto-tracker/internal/repo/sqlite/module"
	"log"
)

type Repo struct {
	db rDB.Sqlite
}

func newRepo(dbfile string) *Repo {
	db, err := sqlite.New(dbfile)
	if err != nil {
		log.Println("!Error init db:", err)
		return nil
	}

	return &Repo{
		db: db,
	}
}

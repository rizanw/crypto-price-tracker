package module

import (
	"crypto-tracker/internal/config"
	rDB "crypto-tracker/internal/repo/sqlite"
	"database/sql"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

type sqlite struct {
	mu sync.Mutex
	db *sql.DB
}

func New(conf config.SqliteConfig) (rDB.Sqlite, error) {
	db, err := sql.Open("sqlite3", conf.Path)
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(createUsersTable); err != nil {
		return nil, err
	}
	if _, err := db.Exec(createCoinsTable); err != nil {
		return nil, err
	}

	return &sqlite{
		db: db,
	}, nil
}

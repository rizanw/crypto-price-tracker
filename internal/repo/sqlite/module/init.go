package module

import (
	rDB "crypto-tracker/internal/repo/sqlite"
	"database/sql"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

type sqlite struct {
	mu sync.Mutex
	db *sql.DB
}

func New(file string) (rDB.Sqlite, error) {
	db, err := sql.Open("sqlite3", file)
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

package module

import (
	rDB "crypto-tracker/internal/repo/sqlite"
	ucAuth "crypto-tracker/internal/usecase/auth"
)

type usecase struct {
	rDB rDB.Sqlite
}

func New(rDB rDB.Sqlite) ucAuth.UseCase {
	return &usecase{
		rDB: rDB,
	}
}

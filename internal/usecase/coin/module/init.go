package module

import (
	rCoincapHttp "crypto-tracker/internal/repo/coincap/http"
	rDB "crypto-tracker/internal/repo/sqlite"
	ucCoin "crypto-tracker/internal/usecase/coin"
)

type usecase struct {
	rDB      rDB.Sqlite
	rCoincap rCoincapHttp.Repo
}

func New(rDB rDB.Sqlite, rCoincap rCoincapHttp.Repo) ucCoin.UseCase {
	return &usecase{
		rDB:      rDB,
		rCoincap: rCoincap,
	}
}

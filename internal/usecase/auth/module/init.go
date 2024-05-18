package module

import (
	"crypto-tracker/internal/config"
	rDB "crypto-tracker/internal/repo/sqlite"
	ucAuth "crypto-tracker/internal/usecase/auth"
)

type usecase struct {
	rDB     rDB.Sqlite
	confJWT config.JWTConfig
}

func New(rDB rDB.Sqlite, confJWT config.JWTConfig) ucAuth.UseCase {
	return &usecase{
		rDB:     rDB,
		confJWT: confJWT,
	}
}

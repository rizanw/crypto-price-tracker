package main

import (
	"crypto-tracker/internal/config"
	"crypto-tracker/internal/usecase/auth"
	ucAuth "crypto-tracker/internal/usecase/auth/module"
	"crypto-tracker/internal/usecase/coin"
	ucCoin "crypto-tracker/internal/usecase/coin/module"
)

type UseCase struct {
	Auth auth.UseCase
	Coin coin.UseCase
}

func newUseCase(conf *config.Config, repo *Repo) UseCase {
	return UseCase{
		Auth: ucAuth.New(repo.db, conf.JWT),
		Coin: ucCoin.New(repo.db, repo.coincapHttp),
	}
}

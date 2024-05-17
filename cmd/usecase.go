package main

import (
	"crypto-tracker/internal/usecase/auth"
	ucAuth "crypto-tracker/internal/usecase/auth/module"
)

type UseCase struct {
	Auth auth.UseCase
}

func newUseCase(repo *Repo) UseCase {
	return UseCase{
		Auth: ucAuth.New(repo.db),
	}
}

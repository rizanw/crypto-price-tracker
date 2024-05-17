package module

import (
	hAuth "crypto-tracker/internal/handler/http/auth"
	"crypto-tracker/internal/usecase/auth"
)

type handler struct {
	ucAuth auth.UseCase
}

func New(ucAuth auth.UseCase) hAuth.Handler {
	return &handler{
		ucAuth: ucAuth,
	}
}

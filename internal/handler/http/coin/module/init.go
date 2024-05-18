package module

import (
	hCoin "crypto-tracker/internal/handler/http/coin"
	ucCoin "crypto-tracker/internal/usecase/coin"
)

type handler struct {
	ucCoin ucCoin.UseCase
}

func New(ucCoin ucCoin.UseCase) hCoin.Handler {
	return &handler{
		ucCoin: ucCoin,
	}
}

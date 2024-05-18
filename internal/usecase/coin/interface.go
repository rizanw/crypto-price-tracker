package coin

import mCoin "crypto-tracker/internal/model/coin"

//go:generate mockgen -package=mock -source=interface.go -destination=./_mock/mock.go
type UseCase interface {
	AddCoin(userID int64, coin string) error
	RemoveCoin(userID int64, coin string) error
	GetCoins(userID int64) (coins []mCoin.Coin, err error)
}

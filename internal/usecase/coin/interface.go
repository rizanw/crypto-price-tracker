package coin

import mCoin "crypto-tracker/internal/model/coin"

type UseCase interface {
	AddCoin(userID int64, coin string) error
	RemoveCoin(userID int64, coin string) error
	GetCoins(userID int64) (coins []mCoin.Coin, err error)
}

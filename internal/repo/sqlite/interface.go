package sqlite

import (
	mCoin "crypto-tracker/internal/model/coin"
	mUser "crypto-tracker/internal/model/user"
)

//go:generate mockgen -package=mock -source=interface.go -destination=./_mock/mock.go
type Sqlite interface {
	InsertUser(email, password string) (userID int64, err error)
	FindUser(email string) (mUser.User, error)
	GetCoins(userID int64) ([]mCoin.CoinDB, error)
	InsertCoin(userID int64, coin string) error
	DeleteCoin(userID int64, coin string) error
}

package coin

type UseCase interface {
	AddCoin(userID int64, coin string) error
	RemoveCoin(userID int64, coin string) error
}

package coin

type UseCase interface {
	AddCoin(userID int64, coin string) error
}

package module

import "errors"

func (u *usecase) AddCoin(userID int64, coin string) error {
	rate, err := u.rCoincap.FindRate(coin)
	if err != nil {
		return err
	}
	if rate.CoinID != coin {
		return errors.New("coin not exist")
	}

	return u.rDB.InsertCoin(userID, coin)
}

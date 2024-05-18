package module

import "errors"

func (u *usecase) AddCoin(userID int64, coin string) error {
	// validate coin rate is exist
	rate, err := u.rCoincap.FindRate(coin)
	if err != nil {
		return err
	}
	if rate.CoinID != coin {
		return errors.New("coin not exist")
	}

	// don't track duplicated coins for a user
	if u.isUserCoinExist(userID, coin) {
		return errors.New("user has coin already")
	}

	// track coin for a user
	return u.rDB.InsertCoin(userID, coin)
}

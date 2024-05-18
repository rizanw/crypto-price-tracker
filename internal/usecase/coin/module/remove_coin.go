package module

import "errors"

func (u *usecase) RemoveCoin(userID int64, coin string) error {
	if !u.isUserCoinExist(userID, coin) {
		return errors.New("user not found coin")
	}

	return u.rDB.DeleteCoin(userID, coin)
}

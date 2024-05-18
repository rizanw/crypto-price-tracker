package module

import (
	"log"
)

func (u *usecase) isUserCoinExist(userID int64, coinName string) bool {
	coins, err := u.rDB.GetCoins(userID)
	if err != nil {
		log.Println("[isUserCoinExist]:", err)
		return false
	}

	for _, coin := range coins {
		if coin.CoindID == coinName {
			return true
		}
	}

	return false
}

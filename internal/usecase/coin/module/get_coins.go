package module

import mCoin "crypto-tracker/internal/model/coin"

func (u *usecase) GetCoins(userID int64) (coins []mCoin.Coin, err error) {
	idrRate, err := u.rCoincap.FindRate("indonesian-rupiah")
	if err != nil {
		return []mCoin.Coin{}, err
	}

	dbCoins, err := u.rDB.GetCoins(userID)
	if err != nil {
		return []mCoin.Coin{}, err
	}

	for _, coin := range dbCoins {
		rate, err := u.rCoincap.FindRate(coin.CoindID)
		if err != nil {
			return []mCoin.Coin{}, err
		}

		coins = append(coins, mCoin.Coin{
			Name:    rate.CoinID,
			RateIdr: rate.RateUsd / idrRate.RateUsd,
		})
	}

	return coins, nil
}

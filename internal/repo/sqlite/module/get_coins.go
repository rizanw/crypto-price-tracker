package module

import mCoin "crypto-tracker/internal/model/coin"

func (r *sqlite) GetCoins(userID int64) ([]mCoin.CoinDB, error) {
	var (
		coins = make([]mCoin.CoinDB, 0)
		err   error
	)

	row, err := r.db.Query(qGetCoins, userID)
	defer row.Close()
	if err != nil {
		return nil, err
	}

	for row.Next() {
		var coin mCoin.CoinDB

		err = row.Scan(
			&coin.ID,
			&coin.CoindID,
			&coin.UserID,
		)
		if err != nil {
			return nil, err
		}

		coins = append(coins, coin)
	}

	return coins, nil
}

package module

func (r *sqlite) DeleteCoin(userID int64, coin string) error {
	_, err := r.db.Exec(qDeleteCoin, userID, coin)
	if err != nil {
		return err
	}

	return nil
}

package module

func (r *sqlite) InsertCoin(userID int64, coin string) error {
	_, err := r.db.Exec(qInsertCoin, coin, userID)
	if err != nil {
		return err
	}

	return nil
}

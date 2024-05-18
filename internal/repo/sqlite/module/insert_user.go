package module

func (r *sqlite) InsertUser(email, password string) (userID int64, err error) {
	err = r.db.QueryRow(qInsertUser, email, password).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

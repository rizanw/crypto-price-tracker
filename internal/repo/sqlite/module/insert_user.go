package module

func (r *sqlite) InsertUser(email, password string) error {
	_, err := r.db.Exec(qInsertUser, email, password)
	if err != nil {
		return err
	}

	return nil
}

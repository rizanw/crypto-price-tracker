package module

import mUser "crypto-tracker/internal/model/user"

func (r *sqlite) FindUser(email string) (mUser.User, error) {
	var (
		user mUser.User
		err  error
	)

	row := r.db.QueryRow(qFindUser, email)

	if err = row.Scan(&user.UserID, &user.Email, &user.Password); err != nil {
		return mUser.User{}, err
	}

	return user, nil
}

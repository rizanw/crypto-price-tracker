package module

import (
	mUser "crypto-tracker/internal/model/user"
	"database/sql"
	"errors"
)

func (r *sqlite) FindUser(email string) (mUser.User, error) {
	var (
		user mUser.User
		err  error
	)

	row := r.db.QueryRow(qFindUser, email)

	if err = row.Scan(&user.UserID, &user.Email, &user.Password); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return mUser.User{}, err
	}

	return user, nil
}

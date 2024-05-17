package sqlite

import (
	mUser "crypto-tracker/internal/model/user"
)

type Sqlite interface {
	InsertUser(email, password string) error
	FindUser(email string) (mUser.User, error)
}

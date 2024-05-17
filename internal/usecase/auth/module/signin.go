package module

import (
	mAuth "crypto-tracker/internal/model/auth"
	"errors"
	"time"
)

func (u *usecase) SignIn(in mAuth.AuthRequest) (mAuth.AuthResponse, error) {
	var (
		res mAuth.AuthResponse
		err error
	)

	user, err := u.rDB.FindUser(in.Email)
	if err != nil {
		return res, err
	}

	// TODO: encrypt password
	if user.Password != in.Password {
		return res, errors.New("invalid password")
	}

	// TODO: generate jwt token
	res.Email = user.Email
	res.Time = time.Now().Unix()
	res.Token = ""

	return res, nil
}

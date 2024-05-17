package module

import (
	mAuth "crypto-tracker/internal/model/auth"
	"errors"
	"time"
)

func (u *usecase) SignUp(in mAuth.AuthRequest) (mAuth.AuthResponse, error) {
	var (
		res mAuth.AuthResponse
		err error
	)

	if err = u.isUserExist(in.Email); err != nil {
		return res, err
	}

	// TODO: encrypt password
	pwd := in.Password

	err = u.rDB.InsertUser(in.Email, pwd)
	if err != nil {
		return res, err
	}

	// TODO: generate jwt token
	res.Email = in.Email
	res.Time = time.Now().Unix()
	res.Token = ""

	return res, nil
}

func (u *usecase) isUserExist(email string) error {
	user, err := u.rDB.FindUser(email)
	if err != nil {
		return err
	}

	if user.UserID != 0 {
		return errors.New("user is already registered")
	}

	return nil
}

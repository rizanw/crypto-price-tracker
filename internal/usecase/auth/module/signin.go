package module

import (
	"crypto-tracker/internal/common/session"
	mAuth "crypto-tracker/internal/model/auth"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (u *usecase) SignIn(in mAuth.AuthRequest) (mAuth.AuthResponse, error) {
	var (
		res mAuth.AuthResponse
		err error
		now = time.Now()
	)

	user, err := u.rDB.FindUser(in.Email)
	if err != nil {
		return res, err
	}
	if user.UserID == 0 {
		return res, errors.New("user not register")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password)); err != nil {
		return res, errors.New("invalid password")
	}

	token, err := u.generateToken(session.Session{
		UserID: user.UserID,
		Email:  user.Email,
		Expiry: now.Add(24 * time.Hour).Unix(),
	})
	if err != nil {
		return res, err
	}

	res.Email = user.Email
	res.Time = now.Unix()
	res.Token = token

	return res, nil
}

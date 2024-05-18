package module

import (
	"crypto-tracker/internal/common/session"
	mAuth "crypto-tracker/internal/model/auth"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (u *usecase) SignUp(in mAuth.AuthRequest) (mAuth.AuthResponse, error) {
	var (
		res mAuth.AuthResponse
		err error
		now = time.Now()
	)

	if err = u.isUserExist(in.Email); err != nil {
		return res, err
	}

	pwd, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return res, err
	}

	userID, err := u.rDB.InsertUser(in.Email, string(pwd))
	if err != nil {
		return res, err
	}

	token, err := u.generateToken(session.Session{
		UserID: userID,
		Email:  in.Email,
		Expiry: now.Add(24 * time.Hour).Unix(),
	})
	if err != nil {
		return res, err
	}

	res.Email = in.Email
	res.Time = now.Unix()
	res.Token = token

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

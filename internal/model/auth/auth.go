package auth

import (
	"errors"
	"net/mail"
	"unicode"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
	Time  int64  `json:"time"`
}

func (r AuthRequest) Validate(isSignup bool, pwdConfirm string) error {
	if r.Email == "" {
		return errors.New("email is required")
	}
	if _, err := mail.ParseAddress(r.Email); err != nil {
		return errors.New("invalid email")
	}

	if r.Password == "" {
		return errors.New("password is required")
	}
	count := 0
	isContainNumber, isContainLetter, isContainSpecial := false, false, false
	for _, c := range r.Password {
		count++
		switch {
		case unicode.IsNumber(c):
			isContainNumber = true
		case unicode.IsLetter(c):
			isContainLetter = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			isContainSpecial = true
		default:
		}
	}
	if !isContainNumber || !isContainLetter || !isContainSpecial || count < 8 {
		return errors.New("invalid password")
	}

	if isSignup && pwdConfirm != r.Password {
		return errors.New("invalid password confirmation")
	}

	return nil
}

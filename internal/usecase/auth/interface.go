package auth

import (
	mAuth "crypto-tracker/internal/model/auth"
)

//go:generate mockgen -package=mock -source=interface.go -destination=./_mock/mock.go
type UseCase interface {
	SignUp(in mAuth.AuthRequest) (mAuth.AuthResponse, error)
	SignIn(in mAuth.AuthRequest) (mAuth.AuthResponse, error)
	SignOut(sessionKey string) error
}

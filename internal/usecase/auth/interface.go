package auth

import mAuth "crypto-tracker/internal/model/auth"

type UseCase interface {
	SignUp(in mAuth.AuthRequest) (mAuth.AuthResponse, error)
}

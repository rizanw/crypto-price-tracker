package module

import (
	"crypto-tracker/internal/common/constants"
	"crypto-tracker/internal/common/session"

	"github.com/golang-jwt/jwt"
)

func generateToken(s session.Session) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": s.UserID,
		"email":   s.Email,
		"exp":     s.Expiry,
	})

	signedToken, err := token.SignedString(constants.JWTSecretKey)
	if err != nil {
		return "", err
	}

	session.Sessions[signedToken] = s
	return signedToken, nil
}

package module

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("secret-key")

func generateToken(email string, tm time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   tm.Add(time.Hour * 24).Unix(),
	})

	return token.SignedString(secretKey)
}

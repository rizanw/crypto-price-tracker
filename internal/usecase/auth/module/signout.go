package module

import (
	"crypto-tracker/internal/common/session"
	"fmt"
)

func (u *usecase) SignOut(sessionKey string) error {
	if _, exists := session.Sessions[sessionKey]; !exists {
		return fmt.Errorf("invalid token")
	}

	delete(session.Sessions, sessionKey)

	return nil
}

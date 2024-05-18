package module

import (
	"crypto-tracker/internal/common/session"
	"net/http"
)

func (h *handler) RemoveCoin(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		ctx = r.Context()
	)

	coin := r.URL.Query().Get("coin")
	if coin == "" {
		http.Error(w, "coin param must be filled", http.StatusBadRequest)
		return
	}
	ses := ctx.Value("session").(session.Session)
	if ses.UserID == 0 {
		http.Error(w, "user not login", http.StatusUnauthorized)
	}

	err = h.ucCoin.RemoveCoin(ses.UserID, coin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

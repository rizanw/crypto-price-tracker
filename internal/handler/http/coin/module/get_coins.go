package module

import (
	"crypto-tracker/internal/common/session"
	"crypto-tracker/internal/model/coin"
	"encoding/json"
	"net/http"
	"time"
)

func (h *handler) GetCoins(w http.ResponseWriter, r *http.Request) {
	var (
		ctx  = r.Context()
		err  error
		now  = time.Now()
		resp struct {
			Data      []coin.Coin `json:"data,omitempty"`
			Timestamp int64       `json:"timestamp,omitempty"`
		}
	)

	ses, ok := ctx.Value("session").(session.Session)
	if ses.UserID == 0 || !ok {
		http.Error(w, "user not login", http.StatusUnauthorized)
	}

	res, err := h.ucCoin.GetCoins(ses.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp.Timestamp = now.Unix()
	resp.Data = res
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

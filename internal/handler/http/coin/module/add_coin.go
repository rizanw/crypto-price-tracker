package module

import (
	"net/http"
)

func (h *handler) AddCoin(w http.ResponseWriter, r *http.Request) {
	var (
		err error
	)

	coin := r.URL.Query().Get("coin")
	if coin == "" {
		http.Error(w, "coin param must be filled", http.StatusBadRequest)
		return
	}

	// TODO get userID from token
	err = h.ucCoin.AddCoin(0, coin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

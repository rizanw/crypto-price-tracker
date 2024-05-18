package coin

import "net/http"

type Handler interface {
	AddCoin(w http.ResponseWriter, r *http.Request)
	RemoveCoin(w http.ResponseWriter, r *http.Request)
	GetCoins(w http.ResponseWriter, r *http.Request)
}

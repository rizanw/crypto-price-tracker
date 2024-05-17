package auth

import "net/http"

type Handler interface {
	Signup(w http.ResponseWriter, r *http.Request)
}

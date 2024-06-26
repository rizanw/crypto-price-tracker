package module

import (
	mAuth "crypto-tracker/internal/model/auth"
	"encoding/json"
	"net/http"
)

func (h *handler) SignUp(w http.ResponseWriter, r *http.Request) {
	var (
		req struct {
			Email                string `json:"email"`
			Password             string `json:"password"`
			PasswordConfirmation string `json:"password_confirmation"`
		}
		authReq mAuth.AuthRequest
		err     error
	)

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	authReq.Email = req.Email
	authReq.Password = req.Password
	if err = authReq.Validate(true, req.PasswordConfirmation); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.ucAuth.SignUp(authReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

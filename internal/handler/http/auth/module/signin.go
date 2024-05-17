package module

import (
	mAuth "crypto-tracker/internal/model/auth"
	"encoding/json"
	"net/http"
)

func (h *handler) SignIn(w http.ResponseWriter, r *http.Request) {
	var (
		req mAuth.AuthRequest
		err error
	)

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = req.Validate(false, ""); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.ucAuth.SignIn(req)
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

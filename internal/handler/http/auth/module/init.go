package module

import hAuth "crypto-tracker/internal/handler/http/auth"

type handler struct {
}

func New() hAuth.Handler {
	return &handler{}
}

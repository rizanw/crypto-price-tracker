package http

import "crypto-tracker/internal/model/coincap"

//go:generate mockgen -package=mock -source=interface.go -destination=./_mock/mock.go
type Repo interface {
	FindRate(id string) (coincap.Rate, error)
}

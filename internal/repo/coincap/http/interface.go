package http

import "crypto-tracker/internal/model/coincap"

type Repo interface {
	FindRates(id string) (coincap.Rate, error)
}

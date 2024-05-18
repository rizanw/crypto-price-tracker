package http

import "crypto-tracker/internal/model/coincap"

type Repo interface {
	FindRate(id string) (coincap.Rate, error)
}

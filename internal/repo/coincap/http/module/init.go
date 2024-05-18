package module

import (
	"crypto-tracker/internal/config"
	coincapHttp "crypto-tracker/internal/repo/coincap/http"
	"fmt"
)

type repo struct {
	url  string
	conf config.HTTP
}

func New(conf config.HTTP) coincapHttp.Repo {
	return &repo{
		url:  fmt.Sprintf("%s/v2", conf.Address),
		conf: conf,
	}
}

package module

import coincapHttp "crypto-tracker/internal/repo/coincap/http"

type repo struct {
	url string
}

func New(url string) coincapHttp.Repo {
	return &repo{
		url: url,
	}
}

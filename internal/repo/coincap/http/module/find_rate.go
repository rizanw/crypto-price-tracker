package module

import (
	"crypto-tracker/internal/model/coincap"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func (r *repo) FindRate(id string) (coincap.Rate, error) {
	var (
		res        coincap.RateResponse
		err        error
		requestURL = fmt.Sprintf("%s/rates/%s", r.url, id)
	)

	resp, err := http.Get(requestURL)
	if err != nil {
		return coincap.Rate{}, err
	}

	defer resp.Body.Close()
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return coincap.Rate{}, err
	}

	err = json.Unmarshal(resBody, &res)
	if err != nil {
		return coincap.Rate{}, err
	}

	reteUSD, _ := strconv.ParseFloat(res.Data.RateUsd, 64)
	return coincap.Rate{
		CoinID:         res.Data.CoinID,
		Symbol:         res.Data.Symbol,
		CurrencySymbol: res.Data.CurrencySymbol,
		Type:           res.Data.Type,
		RateUsd:        reteUSD,
	}, nil
}

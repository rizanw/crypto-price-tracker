package coincap

type Rate struct {
	CoinID         string  `json:"id"`
	Symbol         string  `json:"symbol"`
	CurrencySymbol string  `json:"currencySymbol"`
	Type           string  `json:"type"`
	RateUsd        float64 `json:"rateUsd"`
}

type RateData struct {
	CoinID         string `json:"id"`
	Symbol         string `json:"symbol"`
	CurrencySymbol string `json:"currencySymbol"`
	Type           string `json:"type"`
	RateUsd        string `json:"rateUsd"`
}

type RateResponse struct {
	Data      RateData `json:"data"`
	Timestamp int64    `json:"timestamp"`
}

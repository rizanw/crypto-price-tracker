package coincap

type Rate struct {
	CoinID         string `json:"id"`
	Symbol         string `json:"symbol"`
	CurrencySymbol string `json:"currencySymbol"`
	Type           string `json:"type"`
	RateUsd        string `json:"rateUsd"`
}

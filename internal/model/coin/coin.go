package coin

type CoinDB struct {
	ID      int64
	UserID  int64
	CoindID string
}

type Coin struct {
	Name    string  `json:"name"`
	RateIdr float64 `json:"rate_idr"`
}

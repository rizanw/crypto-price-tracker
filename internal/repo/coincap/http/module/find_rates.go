package module

import (
	"crypto-tracker/internal/model/coincap"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (r *repo) FindRates(id string) (coincap.Rate, error) {
	var (
		res        coincap.Rate
		err        error
		requestURL = fmt.Sprintf("%s/rates/%s", r.url, id)
	)

	resp, err := http.Get(requestURL)
	if err != nil {
		return res, err
	}

	defer resp.Body.Close()
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}

	fmt.Println(string(resBody))
	err = json.Unmarshal(resBody, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

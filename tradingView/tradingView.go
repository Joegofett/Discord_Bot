package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type StockData struct {
	Open   float32 `json:"open"`
	High   float32 `json:"high"`
	Low    float32 `json:"low"`
	Close  float32 `json:"close"`
	Volume float32 `json:"volume"`
	Date   string  `json:"date"`
	Symbol string  `json:"symbol"`
}

type Response struct {
	Data []StockData `json:"data"`
}

func tradingView(stonk string) final {
	
	httpClient := http.Client{}

	req, err := http.NewRequest("GET", "https://api.marketstack.com/v1/tickers/" + stonk + "/eod", nil)
	if err != nil {
		panic(err)
	}

	q := req.URL.Query()
	q.Add("access_key", "")
	req.URL.RawQuery = q.Encode()

	res, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var apiResponse Response
	json.NewDecoder(res.Body).Decode(&apiResponse)

	for _, stockData := range apiResponse.Data {
		final := ("Ticker %s has a day high of %v on %s",
			stockData.Symbol,
			stockData.High,
			stockData.Date)
		return final
	}
	return final
}

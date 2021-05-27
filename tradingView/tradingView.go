package tradingView

import (
	"bytes"
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"
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

func Message(s *discordgo.Session, m *discordgo.MessageCreate) {
	stonk := strings.SplitAfter(m.Content, "$")

	httpClient := http.Client{}

	req, _ := http.NewRequest("GET", "http://api.marketstack.com/v1/tickers/"+stonk[1]+"/eod/latest", nil)

	q := req.URL.Query()
	q.Add("access_key", "5bad3490144c7b35a14f11b47cfcbc1b")
	req.URL.RawQuery = q.Encode()

	res, _ := httpClient.Do(req)

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	newStr := buf.String()
	split := strings.SplitAfter(newStr, ",")
	cleanup := strings.SplitAfter(split[3], ":")
	close := strings.Split(cleanup[1], ",")
	//stockDate := strings.SplitAfter(split[13], ":")

	s.ChannelMessageSend(m.ChannelID, "The most recent close for "+stonk[1]+" is "+close[0])
}

//func Stock(stonk string) (high string, symbol string, stockDate string) {

//}

// req, err := http.NewRequest("GET", "https://api.marketstack.com/v1/tickers/"+stonk, nil)
// if err != nil {
// 	empty := "empty"
// 	discard := "discard"
// 	return err.Error(), empty, discard
// }

// q := req.URL.Query()
// q.Add("access_key", "5bad3490144c7b35a14f11b47cfcbc1b")
// req.URL.RawQuery = q.Encode()

// res, err := httpClient.Do(req)
// if err != nil {
// 	empty := "empty"
// 	discard := "discard"
// 	return err.Error(), empty, discard
// }
// //defer res.Body.Close()

// var apiResponse Response
// json.NewDecoder(res.Body).Decode(&apiResponse)

// for _, stockData := range apiResponse.Data {
// 	symbol = (stockData.Symbol)
// 	stockDate = (stockData.Date)
// 	high = strconv.FormatFloat(float64(stockData.High), 'E', -1, 32)

// }
// return high, symbol, stockDate

package tradingView

import (
	//"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type TickerData struct {
	Open   float64 `json:"open"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Close  float64 `json:"close"`
	Volume float64 `json:"volume"`
	Date   string  `json:"date"`
	Symbol string  `json:"symbol"`
}
type CryptoData struct {
}

func Message(s *discordgo.Session, m *discordgo.MessageCreate) {
	stonk := strings.SplitAfter(m.Content, "$")

	httpClient := http.Client{}

	req, _ := http.NewRequest("GET", "http://api.marketstack.com/v1/tickers/"+stonk[1]+"/eod/latest", nil)

	q := req.URL.Query()
	q.Add("access_key", "78ab56e11bab9073b2681d3c1baa49a7")
	req.URL.RawQuery = q.Encode()

	res, _ := httpClient.Do(req)

	data, _ := io.ReadAll(res.Body)

	var tickerData TickerData
	err := json.Unmarshal(data, &tickerData)
	if err != nil {
		panic(err)
	}
	close := strconv.FormatFloat(tickerData.Close, 'f', -1, 32)
	s.ChannelMessageSend(m.ChannelID, "The most recent close for "+stonk[1]+" is "+close)
}

func Crypto(s *discordgo.Session, m *discordgo.MessageCreate) {
	crypto := strings.SplitAfter(m.Content, "%")

	httpClient := http.Client{}

	req, problem := http.NewRequest("GET", "https://api.coingecko.com/api/v3/simple/price?ids="+crypto[1]+"&vs_currencies=USD", nil)

	if problem != nil {
		s.ChannelMessageSend(m.ChannelID, "That doesn't seem right")
		panic(problem)
	}
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()

	res, errs := httpClient.Do(req)
	if errs != nil {
		panic(errs)
	}

	data, _ := io.ReadAll(res.Body)
	// buf := new(bytes.Buffer)
	// buf.ReadFrom(res.Body)
	// newStr := buf.String()
	// new_ish := strings.SplitAfter(newStr, ":")

	var CryptoData CryptoData
	err := json.Unmarshal(data, &CryptoData)
	if err != nil {
		panic(err)
	}

	s.ChannelMessageSend(m.ChannelID, "   ")
}

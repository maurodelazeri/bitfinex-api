package bitfinex

import (
	"fmt"

	"github.com/ssut/cryptoticker"
)

type Ticks []*Tick

type Tick struct {
	Symbol          string
	Bid             float64
	BidPeriod       int64
	BidSize         float64
	Ask             float64
	AskPeriod       int64
	AskSize         float64
	DailyChange     float64
	DailyChangePerc float64
	LastPrice       float64
	Volume          float64
	High            float64
	Low             float64
}

// bitfinex API implementation of Ticker endpoint.
//
// Endpoint:  tickers?symbols=
// Method: GET
//
// Example: https://api.bitfinex.com/v2/tickers?symbols=
//
// Sample Response:
//
/*
[
  {
  [
    SYMBOL,
    FRR,
    BID,
    BID_SIZE,
    BID_PERIOD,
    ASK,
    ASK_SIZE,
    ASK_PERIOD,
    DAILY_CHANGE,
    DAILY_CHANGE_PERC,
    LAST_PRICE,
    VOLUME,
    HIGH,
    LOW
  ],
  },
  ]
*/

func (client *Client) GetTickers() (Ticks, error) {

	parser := cryptoticker.NewParser(cryptoticker.BitfinexTicker)
	parsed, err := parser.Parse()
	if err != nil {
		return nil, fmt.Errorf("parser: %v", err)
	}

	tickers, err := parsed.Tickers()
	if err != nil {
		return nil, fmt.Errorf("parser: %v", err)
	}

	res := make(Ticks, 0)

	for _, ticker := range tickers {
		// ticker.Currency (the same object as parsed.Coins())
		// ticker.Volume, ticker.Last, ticker.High, ticker.Low, ticker.First
		data := &Tick{}
		data.Symbol = ticker.Currency.String()
		/*
			data.Volume = ticker.Volume
			data.LastPrice = ticker.Last
			data.High = ticker.High
			data.Low = ticker.Low
			data.First = ticker.First
		*/

	}

	return res, nil
}

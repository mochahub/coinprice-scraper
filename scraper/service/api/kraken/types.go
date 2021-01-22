package kraken

import (
	"fmt"
	"strconv"
	"time"
)

type header struct {
	key   string
	value string
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////
// https://docs.pro.coinbase.com/?python#get-historic-rates
//////////////////////////////////////////////////////////////////////////////////////////////////////////
//type CandleStickDataResponse struct {
//	Result map[string][]CandleStickData `json:"result"`
//}

type KrakenResponse struct {
	Error  []string    `json:"error"`
	Result interface{} `json:"result"`
}

//https://github.com/beldur/kraken-go-api-client/blob/master/types.go
// OHLCResponse represents the OHLC's response
//type CandleStickDataResponse struct {
//	Error []interface{} `json:"error"`
//	Pair  string        `json:"pair"`
//	OHLC  []*OHLC       `json:"OHLC"`
//	Last  float64       `json:"last"`
//}

// NewOHLC constructor for OHLC
func NewOHLC(input []interface{}) (*OHLC, error) {
	if len(input) != 8 {
		return nil, fmt.Errorf("the length is not 8 but %d", len(input))
	}

	tmp := new(OHLC)
	tmp.Time = time.Unix(int64(input[0].(float64)), 0)
	tmp.Open, _ = strconv.ParseFloat(input[1].(string), 64)
	tmp.High, _ = strconv.ParseFloat(input[2].(string), 64)
	tmp.Low, _ = strconv.ParseFloat(input[3].(string), 64)
	tmp.Close, _ = strconv.ParseFloat(input[4].(string), 64)
	tmp.Vwap, _ = strconv.ParseFloat(input[5].(string), 64)
	tmp.Volume, _ = strconv.ParseFloat(input[6].(string), 64)
	tmp.Count = int(input[7].(float64))

	return tmp, nil
}

// OHLC represents the "Open-high-low-close chart"
type OHLC struct {
	Time   time.Time `json:"time"`
	Open   float64   `json:"open"`
	High   float64   `json:"high"`
	Low    float64   `json:"low"`
	Close  float64   `json:"close"`
	Vwap   float64   `json:"vwap"`
	Volume float64   `json:"volume"`
	Count  int       `json:"count"`
}

// OHLCResponse represents the OHLC's response
type OHLCResponse struct {
	Pair string  `json:"pair"`
	OHLC []*OHLC `json:"OHLC"`
	Last float64 `json:"last"`
}

//// OHLC represents the "Open-high-low-close chart"
//type OHLC struct {
//	Time   time.Time `json:"time"`
//	Open   float64   `json:"open"`
//	High   float64   `json:"high"`
//	Low    float64   `json:"low"`
//	Close  float64   `json:"close"`
//	Vwap   float64   `json:"vwap"`
//	Volume float64   `json:"volume"`
//	Count  int       `json:"count"`
//}
//
////[
////1611293340,
////"220.03",
////"220.03",
////"220.03",
////"220.03",
////"0.00",
////"0.00000000",
////0
////]
////<time>, <open>, <high>, <low>, <close>, <vwap>, <volume>, <count>
//type CandleStickData struct {
//	OpenTime   float64
//	OpenPrice  float64
//	HighPrice  float64
//	LowPrice   float64
//	ClosePrice float64
//	Volume     float64
//}
//
//func (candleStickResponse *CandleStickData) UnmarshalJSON(
//	data []byte,
//) (err error) {
//	var responseSlice []interface{}
//	if err := json.Unmarshal(data, &responseSlice); err != nil {
//		return err
//	}
//	candleStickResponse.OpenTime = responseSlice[0].(float64)
//	candleStickResponse.OpenPrice, err = strconv.ParseFloat(responseSlice[1].(string), 64)
//	if err != nil {
//		return err
//	}
//	candleStickResponse.HighPrice, err = strconv.ParseFloat(responseSlice[2].(string), 64)
//	if err != nil {
//		return err
//	}
//	candleStickResponse.LowPrice, err = strconv.ParseFloat(responseSlice[3].(string), 64)
//	if err != nil {
//		return err
//	}
//	candleStickResponse.ClosePrice, err = strconv.ParseFloat(responseSlice[4].(string), 64)
//	if err != nil {
//		return err
//	}
//	candleStickResponse.Volume, err = strconv.ParseFloat(responseSlice[5].(string), 64)
//	if err != nil {
//		return err
//	}
//	return nil
//}

//////////////////////////////////////////////////////////////////////////////////////////////////////////
// https://docs.pro.coinbase.com/?python#get-products
//////////////////////////////////////////////////////////////////////////////////////////////////////////
type AssetPairs struct {
	Error  []interface{}        `json:"error"`
	Result map[string]AssetPair `json:"result"`
}

type AssetPair struct {
	Altname       string `json:"altname"`
	Wsname        string `json:"wsname"`
	AclassBase    string `json:"aclass_base"`
	Base          string `json:"base"`
	AclassQuote   string `json:"aclass_quote"`
	Quote         string `json:"quote"`
	Lot           string `json:"lot"`
	PairDecimals  int    `json:"pair_decimals"`
	LotDecimals   int    `json:"lot_decimals"`
	LotMultiplier int    `json:"lot_multiplier"`
}

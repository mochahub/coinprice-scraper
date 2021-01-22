package kraken

import (
	"fmt"
	"github.com/mochahub/coinprice-scraper/scraper/models"
	"github.com/mochahub/coinprice-scraper/scraper/service/api/common"
	"strings"
	"time"
)

func (apiClient *ApiClient) GetSupportedPairs() ([]*models.Symbol, error) {
	products, err := apiClient.getTraceableAssetPairs()
	if err != nil {
		return nil, err
	}
	result := []*models.Symbol{}

	for key := range products.Result {
		product := products.Result[key]
		quote := product.Quote
		normalizedQuote := GetCoinpriceSymbolFromKraken(quote)
		base := product.Base
		normalizedBase := GetCoinpriceSymbolFromKraken(base)
		newPair := &models.Symbol{
			ProductID:       product.Altname,
			RawBase:         base,
			NormalizedBase:  strings.ToUpper(normalizedBase),
			RawQuote:        quote,
			NormalizedQuote: strings.ToUpper(normalizedQuote),
		}

		result = append(result, newPair)
	}
	return common.FilterSupportedAssets(result), nil
}

// Get CandleStick data from [startTime, endTime) with pagination
func (apiClient *ApiClient) GetAllOHLCMarketData(
	symbol models.Symbol,
	interval common.Interval,
	startTime time.Time,
	endTime time.Time,
) ([]*models.OHLCMarketData, error) {
	var durationFromInterval time.Duration
	switch interval {
	case common.Day:
		durationFromInterval = time.Hour * 24
	case common.Hour:
		durationFromInterval = time.Hour
	case common.Minute:
		durationFromInterval = time.Minute
	default:
		return nil, fmt.Errorf("unknown interval: %s", interval)
	}
	if endTime.IsZero() || endTime.After(time.Now()) {
		endTime = time.Now()
	}
	result := []*models.OHLCMarketData{}
	ohlcMarketData, last, err := apiClient.GetOHLCMarketData(
		symbol,
		durationFromInterval,
		startTime,
		0)
	if err != nil {
		return nil, err
	}
	result = append(result, ohlcMarketData...)
	for result[len(result)-1].EndTime.Before(endTime) {
		var newLast int
		ohlcMarketData, newLast, err = apiClient.GetOHLCMarketData(
			symbol,
			durationFromInterval,
			time.Time{},
			last)
		if newLast == last {
			break
		}
		result = append(result, ohlcMarketData...)
		last = newLast
	}

	//last := 0
	//ohlcMarketData := []*models.OHLCMarketData{}
	//var err error
	//for startTime.Before(endTime) || startTime.Equal(endTime) {
	//	//newEndTime := startTime.Add(maxLimit * durationFromInterval)
	//	//if newEndTime.After(endTime) {
	//	//	newEndTime = endTime
	//	//}
	//	ohlcMarketData, last, err = apiClient.GetOHLCMarketData(
	//		symbol,
	//		durationFromInterval,
	//		startTime,
	//		last)
	//	if err != nil {
	//		return nil, err
	//	}
	//	result = append(result, ohlcMarketData...)
	//}
	return result, nil
}

//func reverse(s []*models.OHLCMarketData) []*models.OHLCMarketData {
//	a := make([]*models.OHLCMarketData, len(s))
//	copy(a, s)
//
//	for i := len(a)/2 - 1; i >= 0; i-- {
//		opp := len(a) - 1 - i
//		a[i], a[opp] = a[opp], a[i]
//	}
//
//	return a
//}

func (apiClient *ApiClient) GetRawMarketData() ([]*models.RawMarketData, error) {
	return nil, fmt.Errorf("not implemented")
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////
//// Helpers
////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Get CandleStick data from [startTime, endTime]
func (apiClient *ApiClient) GetOHLCMarketData(
	symbol models.Symbol,
	durationInterval time.Duration,
	startTime time.Time,
	last int,
) ([]*models.OHLCMarketData, int, error) {
	candleStickData, err := apiClient.getCandleStickData(
		int(durationInterval.Minutes()),
		startTime,
		last,
		fmt.Sprintf(symbol.ProductID))
	if err != nil {
		return nil, 0, err
	}
	result := []*models.OHLCMarketData{}
	for i := range candleStickData.OHLC {
		ohlc := candleStickData.OHLC[i]
		result = append(result, &models.OHLCMarketData{
			MarketData: models.MarketData{
				Source:        apiClient.GetExchangeIdentifier(),
				BaseCurrency:  symbol.NormalizedBase,
				QuoteCurrency: symbol.NormalizedQuote,
			},
			StartTime:  ohlc.Time,
			EndTime:    ohlc.Time.Add(durationInterval),
			OpenPrice:  ohlc.Open,
			HighPrice:  ohlc.Close,
			LowPrice:   ohlc.Low,
			ClosePrice: ohlc.Close,
			Volume:     ohlc.Volume,
		})
	}
	return result, int(candleStickData.Last), nil
}

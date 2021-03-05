package okex

import (
	"github.com/mochahub/coinprice-scraper/scraper/service/api/common"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func TestOkexClient(t *testing.T) {
	// TODO: Use DI instead of calling GetSecrets directly
	exchangeClient := NewOkexAPIClient()
	pass := true
	// Get Candles from [start, end]
	pass = t.Run("TestGetCandleStickData", func(t *testing.T) {
		expectedLength := 50 * time.Minute
		startTime := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
		endTime := startTime.Add(expectedLength)
		candleStickResponse, err := exchangeClient.getInstrumentCandles(
			"BTC-USDT",
			time.Minute,
			startTime,
			endTime,
			int64(expectedLength.Minutes()),
		)
		assert.NoError(t, err)
		assert.NotEmpty(t, candleStickResponse)
	}) && pass
	pass = t.Run("TestGetSymbols", func(t *testing.T) {
		exchangeInfo, err := exchangeClient.getInstruments()
		assert.NoError(t, err)
		assert.NotNil(t, exchangeInfo)
		//fmt.Print(utils.PrettyJSON(exchangeInfo))
	}) && pass
	// Interface Methods
	pass = t.Run("TestGetSupportedPairs", func(t *testing.T) {
		pairs, err := exchangeClient.GetSupportedPairs()
		assert.Nil(t, err)
		assert.NotEmpty(t, pairs)
		//fmt.Print(utils.PrettyJSON(pairs))
		assert.Equal(t, 3, len(pairs))
	}) && pass
	pass = t.Run("TestGetAllOHLCMarketData", func(t *testing.T) {
		expectedLength := 20000 * time.Minute
		startTime := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
		endTime := startTime.Add(expectedLength - time.Minute)
		candleStickData, err := exchangeClient.GetAllOHLCMarketData(
			"BTC",
			"USDT",
			common.Minute,
			startTime,
			endTime,
		)
		log.Println(candleStickData[len(candleStickData)-1].StartTime.String())
		log.Println(candleStickData[0].StartTime.String())
		assert.NoError(t, err)
		assert.NotEmpty(t, candleStickData)
		assert.Equal(t, int(expectedLength.Minutes()), len(candleStickData))
	}) && pass

	assert.Equal(t, true, pass)
}

package kraken

import (
	"fmt"
	"github.com/mochahub/coinprice-scraper/config"
	"github.com/mochahub/coinprice-scraper/scraper/models"
	"github.com/mochahub/coinprice-scraper/scraper/service/api/common"
	"github.com/mochahub/coinprice-scraper/scraper/utils"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func TestCoinbaseProClient(t *testing.T) {
	// TODO: Use DI instead of calling GetSecrets directly
	config.LoadEnv()
	exchangeClient := NewKrakenAPIClient()
	pass := true
	pass = t.Run("TestGetProducts", func(t *testing.T) {
		exchangeInfo, err := exchangeClient.getTraceableAssetPairs()
		assert.NoError(t, err)
		assert.NotNil(t, exchangeInfo)
		//fmt.Print(utils.PrettyJSON(exchangeInfo))
	}) && pass
	pass = t.Run("TestGetSupportedPairs", func(t *testing.T) {
		pairs, err := exchangeClient.GetSupportedPairs()
		assert.Nil(t, err)
		assert.NotEmpty(t, pairs)
		//fmt.Print(utils.PrettyJSON(pairs))
		assert.Equal(t, 4, len(pairs))
	}) && pass
	pass = t.Run("TestGetCandleStickData", func(t *testing.T) {
		//expectedLength := maxLimit * time.Minute
		startTime := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
		//endTime := startTime.Add(expectedLength - time.Minute)
		candleStickData, err := exchangeClient.getCandleStickData(
			1,
			startTime,
			0,
			"XBTUSDT",
		)
		fmt.Print(utils.PrettyJSON(candleStickData))
		assert.NoError(t, err)
		assert.NotEmpty(t, candleStickData)
		//assert.Equal(t, int(expectedLength.Minutes()), len(candleStickData))
	}) && pass

	// Interface Methods
	// TODO(Zahin): Do we even need this? exhange_clients_test will test it as well...
	pass = t.Run("TestGetAllOHLCMarketData", func(t *testing.T) {
		expectedLength := 1000 * time.Minute
		startTime := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
		endTime := startTime.Add(expectedLength - time.Minute)
		candleStickData, err := exchangeClient.GetAllOHLCMarketData(
			models.Symbol{
				ProductID:       "XBTUSDT",
				RawBase:         "XXBT",
				NormalizedBase:  "BTC",
				RawQuote:        "USDT",
				NormalizedQuote: "USDT",
			},
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
	////

	assert.Equal(t, true, pass)
}

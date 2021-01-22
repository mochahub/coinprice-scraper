package kraken

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mochahub/coinprice-scraper/scraper/service/api/common"
	"golang.org/x/time/rate"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

type ApiClient struct {
	*retryablehttp.Client
	*rate.Limiter
}

func NewKrakenAPIClient() *ApiClient {
	// 1 callPerSecond
	rateLimiter := rate.NewLimiter(rate.Every(time.Second/rateLimit), 2)
	httpClient := retryablehttp.NewClient()
	httpClient.CheckRetry = common.DefaultCheckRetry
	httpClient.RetryWaitMin = common.DefaultRetryMin
	httpClient.RetryMax = common.MaxRetries
	apiClient := ApiClient{
		Client:  httpClient,
		Limiter: rateLimiter,
	}
	apiClient.RequestLogHook = func(logger retryablehttp.Logger, req *http.Request, retry int) {
		if err := apiClient.Limiter.Wait(context.Background()); err != nil {
			log.Printf("ERROR WAITING FOR LIMIT: %s\n", err.Error())
			return
		}
	}
	return &apiClient
}
func (apiClient *ApiClient) GetExchangeIdentifier() string {
	return KRAKEN
}

func (apiClient *ApiClient) getTraceableAssetPairs() (assetPairsResponse *AssetPairs, err error) {
	urlString := fmt.Sprintf("%s%s", baseURL, assetPairs)
	resp, err := apiClient.sendGetRequest(urlString)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err = json.Unmarshal(body, &assetPairsResponse); err != nil {
		log.Println(string(body))
		return nil, err
	}
	return assetPairsResponse, nil
}

// Get CandleStick data from [startTime, endTime]
func (apiClient *ApiClient) getCandleStickData(
	granularity int,
	since time.Time,
	last int,
	productID string,
) (candleStickResponse *OHLCResponse, err error) {
	params := url.Values{}
	if !since.IsZero() {
		params.Add("since", strconv.FormatInt(since.UnixNano(), 10))
	}
	if last != 0 {
		params.Add("last", strconv.Itoa(last))
	}
	params.Add("pair", productID)
	params.Add("interval", strconv.Itoa(granularity))

	urlString := fmt.Sprintf("%s%s?%s", baseURL, ohlc, params.Encode())
	resp, err := apiClient.sendGetRequest(urlString)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	krakenResponse := KrakenResponse{}
	if err = json.Unmarshal(body, &krakenResponse); err != nil {
		log.Println(string(body))
		return nil, err
	}

	// https://github.com/beldur/kraken-go-api-client/blob/master/krakenapi.go#L169
	mapResponse := krakenResponse.Result.(map[string]interface{})
	// Extracts the list of OHLC from the map to build a slice of interfaces
	OHLCsUnstructured := mapResponse[productID].([]interface{})

	ret := new(OHLCResponse)
	for _, OHLCInterfaceSlice := range OHLCsUnstructured {
		OHLCObj, OHLCErr := NewOHLC(OHLCInterfaceSlice.([]interface{}))
		if OHLCErr != nil {
			return nil, OHLCErr
		}
		ret.OHLC = append(ret.OHLC, OHLCObj)
	}

	ret.Pair = productID
	ret.Last = mapResponse["last"].(float64)
	return ret, nil
}

func (apiClient *ApiClient) sendGetRequest(
	urlString string,
) (*http.Response, error) {
	httpReq, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return nil, err
	}
	retryableRequest, err := retryablehttp.FromRequest(httpReq)
	if err != nil {
		return nil, err
	}
	return apiClient.Do(retryableRequest)
}

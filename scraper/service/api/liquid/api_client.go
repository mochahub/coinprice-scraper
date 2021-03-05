package liquid

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mochahub/coinprice-scraper/config"
	"github.com/mochahub/coinprice-scraper/scraper/service/api/common"
	"golang.org/x/time/rate"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

type ApiClient struct {
	*retryablehttp.Client
	*rate.Limiter
	apiKey        string
	apiSecret     string
	apiPassphrase string
}

func NewLiquidClient(
	secrets *config.Secrets,
) *ApiClient {
	rateLimiter := rate.NewLimiter(rate.Every(time.Minute/callsPerMinute), 1)
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
	return KUCOIN
}

//Get CandleStick data from [startTime, endTime)
//func (apiClient *ApiClient) getKlines(
//	symbol string,
//	interval time.Duration,
//	startTime time.Time,
//	endTime time.Time,
//) (candleStickResponse *CandleStickResponse, err error) {
//	if endTime.IsZero() {
//		endTime = time.Now()
//	}
//	params := url.Values{}
//	params.Add("symbol", symbol)
//	params.Add("startAt", strconv.FormatInt(startTime.Unix(), 10))
//	params.Add("endAt", strconv.FormatInt(endTime.Unix(), 10))
//	params.Add("type", apiClient.intervalQueryParamFromDuration(interval))
//	urlString := fmt.Sprintf("%s%s?%s", baseUrl, getKlines, params.Encode())
//	resp, err := apiClient.sendUnAuthenticatedGetRequest(urlString)
//	if err != nil {
//		return nil, err
//	}
//	defer resp.Body.Close()
//	body, err := ioutil.ReadAll(resp.Body)
//	if err = json.Unmarshal(body, &candleStickResponse); err != nil {
//		return nil, err
//	}
//	return candleStickResponse, nil
//}

// Get ExchangeInfo (supported pairs, precision, etc)
func (apiClient *ApiClient) intervalQueryParamFromDuration(intervalDuration time.Duration) (interval string) {
	return fmt.Sprintf("%dmin", int(intervalDuration.Minutes()))
}

// Get ExchangeInfo (supported pairs, precision, etc)
func (apiClient *ApiClient) getProducts() (products []*Product, err error) {
	urlString := fmt.Sprintf("%s%s", baseUrl, getProducts)
	resp, err := apiClient.sendUnAuthenticatedGetRequest(urlString)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err = json.Unmarshal(body, &products); err != nil {
		return nil, err
	}
	return products, nil
}

func (apiClient *ApiClient) sendUnAuthenticatedGetRequest(
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

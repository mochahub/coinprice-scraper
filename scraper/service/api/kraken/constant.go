package kraken

type Interval string

const (
	KRAKEN = "KRAKEN"
	// calls per second
	rateLimit = 1

	baseURL = "https://api.kraken.com/0/public"
	// Endpoints
	assetPairs = "/AssetPairs"
	ohlc       = "/OHLC" //pair, interval (in minutes) , since (unix timestamp) , last (for pagination)
)

// https://support.kraken.com/hc/en-us/articles/360001185506-How-to-interpret-asset-codes
var krakenToCoinprice = map[string]string{
	"XXBT": "BTC",
	"XETH": "ETH",
}

var coinpriceToKraken = map[string]string{
	"BTC": "XBT",
	"ETH": "XETH",
}

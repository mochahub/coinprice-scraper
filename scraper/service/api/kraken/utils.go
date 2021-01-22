package kraken

import (
	"strings"
)

func GetKrakenSymbolFromCoinprice(
	coinpriceSymbol string,
) string {
	if symbol, ok := coinpriceToKraken[coinpriceSymbol]; ok {
		return symbol
	}
	return coinpriceSymbol
}
func GetCoinpriceSymbolFromKraken(
	coinbaseSymbol string,
) string {
	if symbol, ok := krakenToCoinprice[coinbaseSymbol]; ok {
		return symbol
	}
	return strings.ToUpper(coinbaseSymbol)
}

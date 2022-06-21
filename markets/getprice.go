package markets

import (
	log "github.com/sirupsen/logrus"
	"math"
)

// GetPrice Gets current TON price via CoinMarketCap API
func GetPrice(fiatcurrency string) (float64, error) {
	quote, err := getQuote(fiatcurrency)
	if err != nil {
		log.Error("Can't get information from CoinMarketCap: ", err)
		return 0, err
	}
	return math.Round(quote.Price*100) / 100, nil
}

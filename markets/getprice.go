package markets

import (
	"fmt"
	"math"
)

// GetPrice Gets current TON price via CoinMarketCap API
func GetPrice() float64 {
	quote, err := getQuote()
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return math.Round(quote.Price*100) / 100
}

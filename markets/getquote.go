package markets

import (
	cmc "github.com/miguelmota/go-coinmarketcap/pro/v1"
	"os"
)

func getQuote() (*cmc.Quote, error) {
	client := cmc.NewClient(&cmc.Config{
		ProAPIKey: os.Getenv("APIKEY"),
	})

	results, err := client.Cryptocurrency.LatestQuotes(&cmc.QuoteOptions{
		Convert: "USD",
		Symbol:  "TON",
	})
	if err != nil {
		return nil, err
	}
	return results[0].Quote["USD"], nil
}

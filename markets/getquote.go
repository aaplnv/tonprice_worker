package markets

import (
	cmc "github.com/miguelmota/go-coinmarketcap/pro/v1"
	"os"
)

func getQuote(fiatcurrency string) (*cmc.Quote, error) {
	client := cmc.NewClient(&cmc.Config{
		ProAPIKey: os.Getenv("COINMARKETCAP_APIKEY"),
	})

	results, err := client.Cryptocurrency.LatestQuotes(&cmc.QuoteOptions{
		Convert: fiatcurrency,
		Symbol:  "TON",
	})
	if err != nil {
		return nil, err
	}
	return results[0].Quote[fiatcurrency], nil
}

package handlers

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/layout"
	"main/cache"
	"main/currsettings"
	"main/database"
)

func OnPriceRequest(c telebot.Context) error {
	// Layout temporary initiates here
	lt, err := layout.NewDefault("settings.yml", "default")
	if err != nil {
		log.Fatal(err)
	}

	// Log this event
	log.WithFields(log.Fields{
		"ID":       c.Sender().ID,
		"Username": c.Sender().Username,
	}).Info("New price request")

	// Get the user's settings
	user := database.GetUser(c.Sender().ID)
	stables := currsettings.Create(user)

	// If there is no stables selected, show the list of stables
	if len(stables.All) == 0 {
		return OnSelect(c)
	}

	answer := buildPriceRow(c.Callback().Data, lt) + "\n\n" + lt.Text("exchanges_row") + "\n\n" + lt.Text("ad_row")

	if len(stables.All) < 2 {
		c.EditOrSend(answer, lt.Markup("other"))
	} else {
		c.EditOrSend(answer, lt.Markup("swap", stables.NextTicker(c.Callback().Data), stables.NextTicker(c.Callback().Data)))
	}
	c.Respond()
	return nil
}

func buildPriceRow(currency string, lt *layout.DefaultLayout) string {
	var price float64
	switch currency {
	case "AED":
		price = cache.AEDLatest.Price
	case "ARS":
		price = cache.ARSLatest.Price
	case "AUD":
		price = cache.AUDLatest.Price
	case "BHD":
		price = cache.BHDLatest.Price
	case "BRL":
		price = cache.BRLLatest.Price
	case "BTC":
		price = cache.BTCLatest.Price
	case "CAD":
		price = cache.CADLatest.Price
	case "CHF":
		price = cache.CHFLatest.Price
	case "CLP":
		price = cache.CLPLatest.Price
	case "CNY":
		price = cache.CNYLatest.Price
	case "CZK":
		price = cache.CZKLatest.Price
	case "EUR":
		price = cache.EUROLatest.Price
	case "GBP":
		price = cache.GBPLatest.Price
	case "HKD":
		price = cache.HKDLatest.Price
	case "HUF":
		price = cache.HUFLatest.Price
	case "IDR":
		price = cache.IDRLatest.Price
	case "ILS":
		price = cache.ILSLatest.Price
	case "INR":
		price = cache.INRLatest.Price
	case "JPY":
		price = cache.JPYLatest.Price
	case "MXN":
		price = cache.MXNLatest.Price
	case "NOK":
		price = cache.NOKLatest.Price
	case "NZD":
		price = cache.NZDLatest.Price
	case "PKR":
		price = cache.PKRLatest.Price
	case "PLN":
		price = cache.PLNLatest.Price
	case "RUB":
		price = cache.RUBLatest.Price
	case "SAR":
		price = cache.SARLatest.Price
	case "SEK":
		price = cache.SEKLatest.Price
	case "TRY":
		price = cache.TRYLatest.Price
	case "TWD":
		price = cache.TWDLatest.Price
	case "UAH":
		price = cache.UAHLatest.Price
	case "USD":
		price = cache.USDLatest.Price
	case "ZAR":
		price = cache.ZARLatest.Price
	default:
		log.Error("Unknown currency: " + currency)
	}
	return fmt.Sprint(lt.Text("price_row"), " ", price, " ", currency)
}

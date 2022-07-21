package handlers

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/layout"
	"main/cache"
	"main/currencies"
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
	stables := currencies.Create(user)

	// If there is no stables selected, show the list of stables
	if len(stables.All) == 0 {
		return OnSelect(c)
	}

	var activeStable string
	if c.Callback() != nil {
		activeStable = c.Callback().Data
		if activeStable == "" {
			activeStable = stables.All[0]
		}
	} else {
		activeStable = stables.All[0]
	}

	answer := buildPriceRow(activeStable, lt) + "\n\n" + lt.Text("exchanges_row") + "\n\n" + lt.Text("ad_row")

	if len(stables.All) < 2 {
		c.EditOrSend(answer, lt.Markup("other"))
	} else {
		c.EditOrSend(answer, lt.Markup("swap", stables.NextTicker(activeStable), stables.NextTicker(activeStable)))
	}
	c.Respond()
	return nil
}

func buildPriceRow(currency string, lt *layout.DefaultLayout) string {
	price := cache.LatestQuotes[currency].Price
	return fmt.Sprint(lt.Text("price_row"), " ", price, " ", currency)
}

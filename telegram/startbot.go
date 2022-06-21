package telegram

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	botapi "gopkg.in/telebot.v3"
	"main/markets"
	"os"
	"time"
)

func StartBot() {
	pref := botapi.Settings{
		Token:  os.Getenv("TELEGRAM_BOT_APIKEY"),
		Poller: &botapi.LongPoller{Timeout: 10 * time.Second},
	}

	log.Info("Telegram trying to login...")
	bot, err := botapi.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Info("Logged in!")

	bot.Handle("/start", func(c botapi.Context) error {
		log.WithFields(log.Fields{
			"ID":       c.Message().Sender.ID,
			"Username": c.Message().Sender.Username,
		}).Info("New usdprice request")
		usdprice, err := markets.GetPrice("USD")
		if err != nil {
			return c.Send(fmt.Sprintln("Sorry, failed to get usdprice from CoinMarketCap"))
		}
		rubprice, err := markets.GetPrice("RUB")
		if err != nil {
			return c.Send(fmt.Sprintln("Sorry, failed to get usdprice from CoinMarketCap"))
		}
		return c.Send(fmt.Sprintln("TON price:", usdprice, "USD", "or", rubprice, "RUB"))
	})

	log.Info("Starting a bot...")
	bot.Start()
}

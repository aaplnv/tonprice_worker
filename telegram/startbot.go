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
		Token:  os.Getenv("TOKEN"),
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
		}).Info("New price request")
		price, err := markets.GetPrice()
		if err != nil {
			return c.Send(fmt.Sprintln("Sorry, failed to get price from CoinMarketCap"))
		}
		return c.Send(fmt.Sprintln("TON price:", price, "USD"))
	})

	log.Info("Starting a bot...")
	bot.Start()
}

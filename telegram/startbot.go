package telegram

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	botapi "gopkg.in/telebot.v3"
	"main/cache"
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
		}).Info("New price request")

		price := cache.USDLatest.Price

		answer := fmt.Sprint(os.Getenv("PRICE_ROW"), " ", price, " ", "USD") + "\n\n" + os.Getenv("EXCHANGES_ROW") + "\n\n" + os.Getenv("AD_ROW")
		return c.Send(answer)
	})

	log.Info("Starting a bot...")
	bot.Start()
}

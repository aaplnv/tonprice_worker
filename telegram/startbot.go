package telegram

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/layout"
	"main/handlers"
)

func StartBot() {
	lt, err := layout.NewDefault("settings.yml", "default")
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Logging in to Telegram")

	bot, err := telebot.NewBot(lt.Settings())
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Info("Logged in!")

	bot.Handle("/settings", handlers.OnSelect)

	bot.Handle("/start", handlers.OnPriceRequest)

	bot.Handle(lt.Callback("selector"), handlers.OnSelect)

	bot.Handle(lt.Callback("swap"), handlers.OnPriceRequest)

	bot.Handle(lt.Callback("done"), handlers.OnPriceRequest)

	bot.Handle(lt.Callback("other_currency"), handlers.OnSelect)

	log.Info("Starting a bot...")
	bot.Start()
}

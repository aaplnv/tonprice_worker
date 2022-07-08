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

	log.Info("Telegram trying to login...")

	bot, err := telebot.NewBot(lt.Settings())
	if err != nil {
		log.Fatal(err)
		return
	}

	bot.Handle("/select", handlers.OnSelect)

	log.Info("Logged in!")

	bot.Handle("/start", handlers.OnPriceRequest)

	bot.Handle(lt.Callback("select"), handlers.OnSelect)

	bot.Handle(lt.Callback("swap"), handlers.OnPriceRequest)

	log.Info("Starting a bot...")
	bot.Start()
}

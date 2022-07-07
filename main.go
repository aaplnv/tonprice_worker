package main

import (
	"github.com/jasonlvhit/gocron"
	"gopkg.in/telebot.v3/layout"
	"main/cache"
	"main/telegram"
	"sync"

	log "github.com/sirupsen/logrus"
)

var wg sync.WaitGroup

func init() {
	// Set up logging level here
	log.SetLevel(log.InfoLevel)

	lt, err := layout.NewDefault("settings.yml", "default")
	if err != nil {
		log.Panic("Failed to initialize layout: ", err)
	}

	if lt.Settings().Token == "" {
		log.Panic("Telegram token is not set")
	}

	if lt.String("mariadb.host") == "" {
		log.Panic("Database host is not set")
	}

	if lt.String("mariadb.login") == "" {
		log.Panic("Database login is not set")
	}

	if lt.String("mariadb.password") == "" {
		log.Panic("Database password is not set")
	}

	if lt.String("mariadb.port") == "" {
		log.Panic("Database port is not set")
	}
}

func main() {
	log.Info("Starting application...")
	cache.UpdateQuotes()

	wg.Add(2)
	go func() {
		defer wg.Done()
		telegram.StartBot()
	}()
	gocron.Every(10).Seconds().Do(cache.UpdateQuotes)
	go gocron.Start()
	wg.Wait()
}

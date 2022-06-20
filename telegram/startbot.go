package telegram

import (
	"fmt"
	botapi "gopkg.in/telebot.v3"
	"log"
	"main/markets"
	"os"
	"time"
)

func StartBot() {
	pref := botapi.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &botapi.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := botapi.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", func(c botapi.Context) error {
		return c.Send(fmt.Sprintln("TON price:", markets.GetPrice(), "USD"))
	})

	b.Start()
}

package handlers

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/layout"
	"main/database"
	"strings"
)

func OnSelect(c telebot.Context) error {
	lt, err := layout.NewDefault("settings.yml", "default")
	if err != nil {
		log.Fatal(err)
	}
	log.WithFields(log.Fields{
		"ID":       c.Message().Sender.ID,
		"Username": c.Message().Sender.Username,
	}).Info("New price request")

	user := database.GetUser(c.Sender().ID)

	items := strings.Split(user.ActiveStable, " ")
	btns := make([]telebot.Btn, len(items))

	for i, item := range items {
		btns[i] = *lt.Button("select", struct {
			Ticker string
			Text   string
		}{
			Ticker: item,
			Text:   item,
		})
	}

	m := c.Bot().NewMarkup()
	m.Inline(m.Row(btns...))

	return c.Send("Select currency", m)
}

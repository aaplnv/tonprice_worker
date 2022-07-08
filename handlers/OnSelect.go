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

	if c.Callback() != nil && c.Callback().Data != "" {

		if strings.Contains(user.AllStables, c.Callback().Data) {
			database.SetAllStables(strings.ReplaceAll(user.AllStables, c.Callback().Data, ""), user.TelegramId)
			user = database.GetUser(c.Sender().ID)
		} else {
			database.SetAllStables(user.AllStables+" "+c.Callback().Data, user.TelegramId)
			user = database.GetUser(c.Sender().ID)
		}
	}

	//allbtns := strings.Split(user.ActiveStable, " ")
	items := lt.Strings("currencies")
	btns := make([]telebot.Btn, len(items))

	for i, item := range items {
		if item == "" {
			continue
		}
		text := ""
		if strings.Contains(user.AllStables, item) {
			text = "✅" + " " + item
		} else {
			text = item
		}
		btns[i] = *lt.Button("select", struct {
			Ticker string
			Text   string
		}{
			Ticker: item,
			Text:   text,
		})
	}

	m := c.Bot().NewMarkup()

	if user.AllStables != "" {
		m.Inline(m.Row(btns...), m.Row(*lt.Button("done")))
	} else {
		m.Inline(m.Row(btns...))
	}

	return c.EditOrSend("Select currency", m)
}

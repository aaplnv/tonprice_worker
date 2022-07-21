package handlers

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/layout"
	"main/currencies"
	"main/database"
)

func OnSelect(c telebot.Context) error {
	// Log this event
	log.WithFields(log.Fields{
		"ID":       c.Sender().ID,
		"Username": c.Sender().Username,
	}).Info("New select request")

	// Layout temporary initiates here
	lt, err := layout.NewDefault("settings.yml", "default")
	if err != nil {
		log.Fatal(err)
	}

	// Get user from database
	user := database.GetUser(c.Sender().ID)
	currs := currencies.Create(user)

	// If request contains a currency, add or remove it from user's list
	if HasCallbackData(c.Callback()) {
		if currs.IsAdded(c.Callback().Data) {
			currs.RemoveStable(c.Callback().Data)
		} else {
			currs.AddStable(c.Callback().Data)
		}
	}

	// Get supported currencies list
	items := lt.Strings("currencies")

	// Build a row of buttons
	btns := make([]telebot.Btn, len(items))

	// For each currency, add a button
	for i, item := range items {

		// If item is empty continue
		if item == "" {
			continue
		}

		// Draw an emoji if currency already selected
		text := ""
		if currs.IsAdded(item) {
			text = "✅" + " " + item
		} else {
			text = item
		}

		// Add button to row
		btns[i] = *lt.Button("selector", struct {
			Ticker string
			Text   string
		}{
			Ticker: item,
			Text:   text,
		})
	}

	m := c.Bot().NewMarkup()

	// if user has any currency, allow him to exit
	if len(currs.All) != 0 {
		buttons := m.Split(4, btns)
		buttons = append(buttons, m.Row(*lt.Button("done", lt.Text("exit_settings_btn"))))
		m.Inline(buttons...)
	} else {
		m.Inline(m.Split(4, btns)...)
	}

	return c.EditOrSend(lt.Text("select_currency_row"), m)
}

func HasCallbackData(c *telebot.Callback) bool {
	return c != nil && c.Data != ""
}

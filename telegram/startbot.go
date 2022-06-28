package telegram

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	botapi "gopkg.in/telebot.v3"
	"main/cache"
	"os"
	"time"
)

var gotousd = (&botapi.ReplyMarkup{}).Data("🔄 USD", "goto_usd", "sjdgsdflgdfg")
var returntomaincurr = (&botapi.ReplyMarkup{}).Data("🔄 RUB", "returntomaincurr", "sdfgjsdkfg")
var subbtn = (&botapi.ReplyMarkup{}).Data("🛎 Subscribe", "subbtn", "shjfsdf")
var unsubbtn = (&botapi.ReplyMarkup{}).Data("🔕 Unsubscribe", "unsubbtn", "shdfjksldqwe")
var confirmsubbtn = (&botapi.ReplyMarkup{}).Data("Confirm subscription", "confirmsubbtn", "svfderwer")
var confirmusdbtn = (&botapi.ReplyMarkup{}).Data("Confirm subscription in USD", "confirmsubusdbtn", "svfderwer")
var cancelbtn = (&botapi.ReplyMarkup{}).Data("Cancel", "cancel", "sgdhfsjkewrvb")

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
		return c.Send(proceedRequest("RUB"))
	})

	bot.Handle(&gotousd, func(c botapi.Context) error {
		log.WithFields(log.Fields{
			"ID":       c.Message().Sender.ID,
			"Username": c.Message().Sender.Username,
		}).Info("New price request")
		c.Edit(proceedRequest("USD"))
		return c.Respond()
	})

	bot.Handle(&subbtn, func(c botapi.Context) error {
		r := botapi.ReplyMarkup{}
		r.Inline(r.Row(confirmsubbtn), r.Row(confirmusdbtn), r.Row(cancelbtn))
		c.Edit("We will notify you when price changes more than for 10% at any direction from your active currency (RUB)\n\nDo you confirm your subscription?", &r)
		return c.Respond()
	})

	bot.Handle(&returntomaincurr, func(c botapi.Context) error {
		log.WithFields(log.Fields{
			"ID":       c.Message().Sender.ID,
			"Username": c.Message().Sender.Username,
		}).Info("New price request")
		c.Edit(proceedRequest("RUB"))
		return c.Respond()
	})

	bot.Handle(&cancelbtn, func(c botapi.Context) error {
		log.WithFields(log.Fields{
			"ID":       c.Message().Sender.ID,
			"Username": c.Message().Sender.Username,
		}).Info("New price request")
		c.Edit(proceedRequest("RUB"))
		return c.Respond()
	})

	log.Info("Starting a bot...")
	bot.Start()
}

func proceedRequest(currency string) (string, *botapi.ReplyMarkup) {
	answer := buildPriceRow(currency) + "\n\n" + os.Getenv("EXCHANGES_ROW") + "\n\n" + os.Getenv("AD_ROW")

	r := botapi.ReplyMarkup{}

	if currency == "USD" {
		r.Inline(r.Row(returntomaincurr, subbtn))
	} else {
		r.Inline(r.Row(gotousd, subbtn))
	}

	return answer, &r
}

func buildPriceRow(currency string) string {
	var price float64
	switch currency {
	case "AED":
		price = cache.AEDLatest.Price
	case "ARS":
		price = cache.ARSLatest.Price
	case "AUD":
		price = cache.AUDLatest.Price
	case "BHD":
		price = cache.BHDLatest.Price
	case "BRL":
		price = cache.BRLLatest.Price
	case "BTC":
		price = cache.BTCLatest.Price
	case "CAD":
		price = cache.CADLatest.Price
	case "CHF":
		price = cache.CHFLatest.Price
	case "CLP":
		price = cache.CLPLatest.Price
	case "CNY":
		price = cache.CNYLatest.Price
	case "CZK":
		price = cache.CZKLatest.Price
	case "EUR":
		price = cache.EUROLatest.Price
	case "GBP":
		price = cache.GBPLatest.Price
	case "HKD":
		price = cache.HKDLatest.Price
	case "HUF":
		price = cache.HUFLatest.Price
	case "IDR":
		price = cache.IDRLatest.Price
	case "ILS":
		price = cache.ILSLatest.Price
	case "INR":
		price = cache.INRLatest.Price
	case "JPY":
		price = cache.JPYLatest.Price
	case "MXN":
		price = cache.MXNLatest.Price
	case "NOK":
		price = cache.NOKLatest.Price
	case "NZD":
		price = cache.NZDLatest.Price
	case "PKR":
		price = cache.PKRLatest.Price
	case "PLN":
		price = cache.PLNLatest.Price
	case "RUB":
		price = cache.RUBLatest.Price
	case "SAR":
		price = cache.SARLatest.Price
	case "SEK":
		price = cache.SEKLatest.Price
	case "TRY":
		price = cache.TRYLatest.Price
	case "TWD":
		price = cache.TWDLatest.Price
	case "UAH":
		price = cache.UAHLatest.Price
	case "USD":
		price = cache.USDLatest.Price
	case "ZAR":
		price = cache.ZARLatest.Price
	}
	return fmt.Sprint(os.Getenv("PRICE_ROW"), " ", price, " ", currency)
}

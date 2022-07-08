package telegram

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/layout"
	"main/cache"
	"main/ent"
	"main/ent/user"
	"time"
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

	bot.Handle("/select", func(c telebot.Context) error {
		user := getUser(c.Sender().ID)
		fmt.Println(user)
		items := []string{"USD", "RUB", "EUR", "GBP"}
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

		m := bot.NewMarkup()
		m.Inline(m.Row(btns...))

		return c.Send("Select currency", m)
	})

	log.Info("Logged in!")

	bot.Handle("/start", func(c telebot.Context) error {
		log.WithFields(log.Fields{
			"ID":       c.Message().Sender.ID,
			"Username": c.Message().Sender.Username,
		}).Info("New price request")

		return c.Send(proceedRequest("RUB", lt))
	})

	bot.Handle(lt.Callback("swap"), func(c telebot.Context) error {
		log.WithFields(log.Fields{
			"ID":       c.Message().Sender.ID,
			"Username": c.Message().Sender.Username,
		}).Info("New price request")

		print(lt.Get("layouts"))
		return c.Edit(proceedRequest(c.Callback().Data, lt))
	})

	log.Info("Starting a bot...")
	bot.Start()
}

func proceedRequest(currency string, lt *layout.DefaultLayout) (string, *telebot.ReplyMarkup) {
	answer := buildPriceRow(currency, lt) + "\n\n" + lt.Text("exchanges_row") + "\n\n" + lt.Text("ad_row")
	if currency == "USD" {
		return answer, lt.Markup("swap", "RUB")
	}
	return answer, lt.Markup("swap", "USD")
}

func getUser(id int64) *ent.User {
	lt, err := layout.NewDefault("settings.yml", "default")
	if err != nil {
		log.Fatal(err)
	}
	client, err := ent.Open("mysql", fmt.Sprint(lt.String("mariadb.login"), ":", lt.String("mariadb.password"), "@tcp(", lt.String("mariadb.host"), ":", lt.String("mariadb.port"), ")/testone?parseTime=True"))
	if err != nil {
		log.Error("Can't connect to MySQL", err)
	}
	defer client.Close()

	// TODO: Automatic tables migration

	user, err := client.User.Query().Where(user.TelegramId(id)).Only(context.Background())
	if err != nil {
		user, err = client.User.Create().SetTelegramId(id).SetRegTime(time.Now()).Save(context.Background())
		if err != nil {
			log.Error(err)
		}
	}

	return user
}

func setAllCurrencies(currencies string, id int64) {
	lt, err := layout.NewDefault("settings.yml", "default")
	if err != nil {
		log.Fatal(err)
	}
	client, err := ent.Open("mysql", fmt.Sprint(lt.String("mariadb.login"), ":", lt.String("mariadb.password"), "@tcp(", lt.String("mariadb.host"), ":", lt.String("mariadb.port"), ")/testone?parseTime=True"))
	if err != nil {
		log.Error("Can't connect to MySQL", err)
	}
	defer client.Close()

	_, err = client.User.Update().Where(user.TelegramId(id)).SetAllStables(currencies).Save(context.Background())

	if err != nil {
		log.Error(err)
	}
}

func buildPriceRow(currency string, lt *layout.DefaultLayout) string {
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
	default:
		log.Error("Unknown currency: " + currency)
	}
	return fmt.Sprint(lt.Text("price_row"), " ", price, " ", currency)
}

package cache

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"gopkg.in/telebot.v3/layout"
	"main/ent"
	"main/ent/aedquote"
	"main/ent/arsquote"
	"main/ent/audquote"
	"main/ent/bhdquote"
	"main/ent/brlquote"
	"main/ent/btcquote"
	"main/ent/cadquote"
	"main/ent/chfquote"
	"main/ent/clpquote"
	"main/ent/cnyquote"
	"main/ent/czkquote"
	"main/ent/euroquote"
	"main/ent/gbpquote"
	"main/ent/hkdquote"
	"main/ent/hufquote"
	"main/ent/idrquote"
	"main/ent/ilsquote"
	"main/ent/inrquote"
	"main/ent/jpyquote"
	"main/ent/mxnquote"
	"main/ent/nokquote"
	"main/ent/nzdquote"
	"main/ent/pkrquote"
	"main/ent/plnquote"
	"main/ent/rubquote"
	"main/ent/sarquote"
	"main/ent/sekquote"
	"main/ent/tryquote"
	"main/ent/twdquote"
	"main/ent/uahquote"
	"main/ent/usdquote"
	"main/ent/zarquote"
	"time"
)

func UpdateQuotes() {
	log.Info("Updating cache...")
	lt, err := layout.NewDefault("settings.yml", "default")
	if err != nil {
		log.Fatal(err)
	}
	client, err := ent.Open("mysql", fmt.Sprint(lt.String("mariadb.login"), ":", lt.String("mariadb.password"), "@tcp(", lt.String("mariadb.host"), ":", lt.String("mariadb.port"), ")/charts?parseTime=True"))
	if err != nil {
		log.Error("Can't connect to MySQL", err)
		return
	}
	defer client.Close()

	currencies := lt.Strings("currencies")

	for _, currency := range currencies {

		if currency == "AED" {
			query, err := client.AEDQuote.Query().Where(aedquote.TimestampGTE(time.Now().Add(-time.Hour * 24))).All(context.Background())
			if err != nil {
				log.Error("Can't get AED rates", err)
				return
			}

			//
			// Somewhere here we should build charts
			//

			LatestQuotes["AED"] = LatestQuote{
				Price:     query[len(query)-1:][0].Price,
				Timestamp: query[len(query)-1:][0].Timestamp,
			}
		}

		if currency == "ARS" {
			query, err := client.ARSQuote.Query().Where(arsquote.TimestampGTE(time.Now().Add(-time.Hour * 24))).All(context.Background())
			if err != nil {
				log.Error("Can't get ARS rates", err)
				return
			}

			//
			// Somewhere here we should build charts
			//

			LatestQuotes["ARS"] = LatestQuote{
				Price:     query[len(query)-1:][0].Price,
				Timestamp: query[len(query)-1:][0].Timestamp,
			}
		}

		if currency == "AUD" {
			query, err := client.AUDQuote.Query().Where(audquote.TimestampGTE(time.Now().Add(-time.Hour * 24))).All(context.Background())
			if err != nil {
				log.Error("Can't get AUD rates", err)
				return
			}

			//
			// Somewhere here we should build charts
			//

			LatestQuotes["AUD"] = LatestQuote{
				Price:     query[len(query)-1:][0].Price,
				Timestamp: query[len(query)-1:][0].Timestamp,
			}
		}

		if currency == "BHD" {
			query, err := client.BHDQuote.Query().Where(bhdquote.TimestampGTE(time.Now().Add(-time.Hour * 24))).All(context.Background())
			if err != nil {
				log.Error("Can't get BHD rates", err)
				return
			}

			//
			// Somewhere here we should build charts
			//

			LatestQuotes["BHD"] = LatestQuote{
				Price:     query[len(query)-1:][0].Price,
				Timestamp: query[len(query)-1:][0].Timestamp,
			}
		}

		if currency == "BRL" {
			query, err := client.BRLQuote.Query().Where(brlquote.TimestampGTE(time.Now().Add(-time.Hour * 24))).All(context.Background())
			if err != nil {
				log.Error("Can't get BRL rates", err)
				return
			}

			//
			// Somewhere here we should build charts
			//

			LatestQuotes["BRL"] = LatestQuote{
				Price:     query[len(query)-1:][0].Price,
				Timestamp: query[len(query)-1:][0].Timestamp,
			}

		}

		if currency == "BTC" {
			query, err := client.BTCQuote.Query().Where(btcquote.TimestampGTE(time.Now().Add(-time.Hour * 24))).All(context.Background())
			if err != nil {
				log.Error("Can't get BTC rates", err)
				return
			}

			//
			// Somewhere here we should build charts
			//

			LatestQuotes["BTC"] = LatestQuote{
				Price:     query[len(query)-1:][0].Price,
				Timestamp: query[len(query)-1:][0].Timestamp,
			}

		}

		if currency == "CAD" {
			query, err := client.CADQuote.Query().Where(cadquote.TimestampGTE(time.Now().Add(-time.Hour * 24))).All(context.Background())
			if err != nil {
				log.Error("Can't get CAD rates", err)
				return
			}

			//
			// Somewhere here we should build charts
			//

			LatestQuotes["CAD"] = LatestQuote{
				Price:     query[len(query)-1:][0].Price,
				Timestamp: query[len(query)-1:][0].Timestamp,
			}

		}

		if currency == "CHF" {
			query, err := client.CHFQuote.Query().Where(chfquote.TimestampGTE(time.Now().Add(-time.Hour * 24))).All(context.Background())
			if err != nil {
				log.Error("Can't get CHF rates", err)
				return
			}

			//
			// Somewhere here we should build charts
			//

			LatestQuotes["CHF"] = LatestQuote{
				Price:     query[len(query)-1:][0].Price,
				Timestamp: query[len(query)-1:][0].Timestamp,
			}
		}

		if currency == "CLP" {
			query, err := client.CLPQuote.Query().Where(clpquote.TimestampGTE(time.Now().Add(-time.Hour * 24))).All(context.Background())
			if err != nil {
				log.Error("Can't get CLP rates", err)
				return
			}

			//
			// Somewhere here we should build charts
			//

			LatestQuotes["CLP"] = LatestQuote{
				Price:     query[len(query)-1:][0].Price,
				Timestamp: query[len(query)-1:][0].Timestamp,
			}
		}

		if currency == "CNY" {
			query, err := client.CNYQuote.Query().Where(cnyquote.TimestampGTE(time.Now().Add(-time.Hour * 24))).All(context.Background())
			if err != nil {
				log.Error("Can't get CNY rates", err)
				return
			}

			//
			// Somewhere here we should build charts
			//

			LatestQuotes["CNY"] = LatestQuote{
				Price:     query[len(query)-1:][0].Price,
				Timestamp: query[len(query)-1:][0].Timestamp,
			}
		}

		if currency == "CZK" {
			query, err := client.CZKQuote.Query().Where(czkquote.TimestampGTE(time.Now().Add(-time.Hour * 24))).All(context.Background())
			if err != nil {
				log.Error("Can't get CZK rates", err)
				return
			}

			//
			// Somewhere here we should build charts
			//

			LatestQuotes["CZK"] = LatestQuote{
				Price:     query[len(query)-1:][0].Price,
				Timestamp: query[len(query)-1:][0].Timestamp,
			}
		}

		if currency == "EUR" {
			query, err := client.EUROQuote.Query().Where(euroquote.TimestampGTE(time.Now().Add(-time.Hour * 24))).All(context.Background())
			if err != nil {
				log.Error("Can't get EUR rates", err)
				return
			}

			//
			// Somewhere here we should build charts
			//

			LatestQuotes["EUR"] = LatestQuote{
				Price:     query[len(query)-1:][0].Price,
				Timestamp: query[len(query)-1:][0].Timestamp,
			}
		}

		if currency == "GBP" {
			query, err := client.GBPQuote.Query().Where(gbpquote.TimestampGTE(time.Now().Add(-time.Hour * 24))).All(context.Background())
			if err != nil {
				log.Error("Can't get GBP rates", err)
				return
			}

			//
			// Somewhere here we should build charts
			//

			LatestQuotes["GBP"] = LatestQuote{
				Price:     query[len(query)-1:][0].Price,
				Timestamp: query[len(query)-1:][0].Timestamp,
			}
		}

		if currency == "HKD" {
			query, err := client.HKDQuote.Query().Where(hkdquote.TimestampGTE(time.Now().Add(-time.Hour * 24))).All(context.Background())
			if err != nil {
				log.Error("Can't get HKD rates", err)
				return
			}

			//
			// Somewhere here we should build charts
			//

			LatestQuotes["HKD"] = LatestQuote{
				Price:     query[len(query)-1:][0].Price,
				Timestamp: query[len(query)-1:][0].Timestamp,
			}
		}

		if currency == "HUF" {
			query, err := client.HUFQuote.Query().Where(hufquote.TimestampGTE(time.Now().Add(-time.Hour * 24))).All(context.Background())
			if err != nil {
				log.Error("Can't get HUF rates", err)
				return
			}

			//
			// Somewhere here we should build charts
			//

			LatestQuotes["HUF"] = LatestQuote{
				Price:     query[len(query)-1:][0].Price,
				Timestamp: query[len(query)-1:][0].Timestamp,
			}
		}

		if currency == "IDR" {
			query, err := client.IDRQuote.Query().Where(idrquote.TimestampGTE(time.Now().Add(-time.Hour * 24))).All(context.Background())
			if err != nil {
				log.Error("Can't get IDR rates", err)
				return
			}

			//
			// Somewhere here we should build charts
			//

			LatestQuotes["IDR"] = LatestQuote{
				Price:     query[len(query)-1:][0].Price,
				Timestamp: query[len(query)-1:][0].Timestamp,
			}
		}

		if currency == "ILS" {
			query, err := client.ILSQuote.Query().Where(ilsquote.TimestampGTE(time.Now().Add(-time.Hour * 24))).All(context.Background())
			if err != nil {
				log.Error("Can't get ILS rates", err)
				return
			}

			//
			// Somewhere here we should build charts
			//

			LatestQuotes["ILS"] = LatestQuote{
				Price:     query[len(query)-1:][0].Price,
				Timestamp: query[len(query)-1:][0].Timestamp,
			}
		}

		if currency == "INR" {
			query, err := client.INRQuote.Query().Where(inrquote.TimestampGTE(time.Now().Add(-time.Hour * 24))).All(context.Background())
			if err != nil {
				log.Error("Can't get INR rates", err)
				return
			}

			//
			// Somewhere here we should build charts
			//

			LatestQuotes["INR"] = LatestQuote{
				Price:     query[len(query)-1:][0].Price,
				Timestamp: query[len(query)-1:][0].Timestamp,
			}
		}

		if currency == "JPY" {
			query, err := client.JPYQuote.Query().Where(jpyquote.TimestampGTE(time.Now().Add(-time.Hour * 24))).All(context.Background())
			if err != nil {
				log.Error("Can't get JPY rates", err)
				return
			}

			//
			// Somewhere here we should build charts
			//

			LatestQuotes["JPY"] = LatestQuote{
				Price:     query[len(query)-1:][0].Price,
				Timestamp: query[len(query)-1:][0].Timestamp,
			}
		}

		if currency == "MXN" {
			query, err := client.MXNQuote.Query().Where(mxnquote.TimestampGTE(time.Now().Add(-time.Hour * 24))).All(context.Background())
			if err != nil {
				log.Error("Can't get MXN rates", err)
				return
			}

			//
			// Somewhere here we should build charts
			//

			LatestQuotes["MXN"] = LatestQuote{
				Price:     query[len(query)-1:][0].Price,
				Timestamp: query[len(query)-1:][0].Timestamp,
			}
		}

		if currency == "NOK" {
			query, err := client.NOKQuote.Query().Where(nokquote.TimestampGTE(time.Now().Add(-time.Hour * 24))).All(context.Background())
			if err != nil {
				log.Error("Can't get NOK rates", err)
				return
			}

			//
			// Somewhere here we should build charts
			//

			LatestQuotes["NOK"] = LatestQuote{
				Price:     query[len(query)-1:][0].Price,
				Timestamp: query[len(query)-1:][0].Timestamp,
			}
		}

		if currency == "NZD" {
			query, err := client.NZDQuote.Query().Where(nzdquote.TimestampGTE(time.Now().Add(-time.Hour * 24))).All(context.Background())
			if err != nil {
				log.Error("Can't get NZD rates", err)
				return
			}

			//
			// Somewhere here we should build charts
			//

			LatestQuotes["NZD"] = LatestQuote{
				Price:     query[len(query)-1:][0].Price,
				Timestamp: query[len(query)-1:][0].Timestamp,
			}
		}

		if currency == "PKR" {
			query, err := client.PKRQuote.Query().Where(pkrquote.TimestampGTE(time.Now().Add(-time.Hour * 24))).All(context.Background())
			if err != nil {
				log.Error("Can't get PKR rates", err)
				return
			}

			//
			// Somewhere here we should build charts
			//

			LatestQuotes["PKR"] = LatestQuote{
				Price:     query[len(query)-1:][0].Price,
				Timestamp: query[len(query)-1:][0].Timestamp,
			}
		}

		if currency == "PLN" {
			query, err := client.PLNQuote.Query().Where(plnquote.TimestampGTE(time.Now().Add(-time.Hour * 24))).All(context.Background())
			if err != nil {
				log.Error("Can't get PLN rates", err)
				return
			}

			//
			// Somewhere here we should build charts
			//

			LatestQuotes["PLN"] = LatestQuote{
				Price:     query[len(query)-1:][0].Price,
				Timestamp: query[len(query)-1:][0].Timestamp,
			}
		}

		if currency == "RUB" {
			query, err := client.RUBQuote.Query().Where(rubquote.TimestampGTE(time.Now().Add(-time.Hour * 24))).All(context.Background())
			if err != nil {
				log.Error("Can't get RUB rates", err)
				return
			}

			//
			// Somewhere here we should build charts
			//

			LatestQuotes["RUB"] = LatestQuote{
				Price:     query[len(query)-1:][0].Price,
				Timestamp: query[len(query)-1:][0].Timestamp,
			}

		}

		if currency == "SAR" {
			query, err := client.SARQuote.Query().Where(sarquote.TimestampGTE(time.Now().Add(-time.Hour * 24))).All(context.Background())
			if err != nil {
				log.Error("Can't get SAR rates", err)
				return
			}

			//
			// Somewhere here we should build charts
			//

			LatestQuotes["SAR"] = LatestQuote{
				Price:     query[len(query)-1:][0].Price,
				Timestamp: query[len(query)-1:][0].Timestamp,
			}
		}

		if currency == "SEK" {
			query, err := client.SEKQuote.Query().Where(sekquote.TimestampGTE(time.Now().Add(-time.Hour * 24))).All(context.Background())
			if err != nil {
				log.Error("Can't get SEK rates", err)
				return
			}

			//
			// Somewhere here we should build charts
			//

			LatestQuotes["SEK"] = LatestQuote{
				Price:     query[len(query)-1:][0].Price,
				Timestamp: query[len(query)-1:][0].Timestamp,
			}
		}

		if currency == "TRY" {
			query, err := client.TRYQuote.Query().Where(tryquote.TimestampGTE(time.Now().Add(-time.Hour * 24))).All(context.Background())
			if err != nil {
				log.Error("Can't get TRY rates", err)
				return
			}

			//
			// Somewhere here we should build charts
			//

			LatestQuotes["TRY"] = LatestQuote{
				Price:     query[len(query)-1:][0].Price,
				Timestamp: query[len(query)-1:][0].Timestamp,
			}
		}

		if currency == "TWD" {
			query, err := client.TWDQuote.Query().Where(twdquote.TimestampGTE(time.Now().Add(-time.Hour * 24))).All(context.Background())
			if err != nil {
				log.Error("Can't get TWD rates", err)
				return
			}

			//
			// Somewhere here we should build charts
			//

			LatestQuotes["TWD"] = LatestQuote{
				Price:     query[len(query)-1:][0].Price,
				Timestamp: query[len(query)-1:][0].Timestamp,
			}
		}

		if currency == "UAH" {
			query, err := client.UAHQuote.Query().Where(uahquote.TimestampGTE(time.Now().Add(-time.Hour * 24))).All(context.Background())
			if err != nil {
				log.Error("Can't get UAH rates", err)
				return
			}

			//
			// Somewhere here we should build charts
			//

			LatestQuotes["UAH"] = LatestQuote{
				Price:     query[len(query)-1:][0].Price,
				Timestamp: query[len(query)-1:][0].Timestamp,
			}
		}

		if currency == "USD" {
			query, err := client.USDQuote.Query().Where(usdquote.TimestampGTE(time.Now().Add(-time.Hour * 24))).All(context.Background())
			if err != nil {
				log.Error("Can't get USD rates", err)
				return
			}

			//
			// Somewhere here we should build charts
			//

			LatestQuotes["USD"] = LatestQuote{
				Price:     query[len(query)-1:][0].Price,
				Timestamp: query[len(query)-1:][0].Timestamp,
			}
		}

		if currency == "ZAR" {
			query, err := client.ZARQuote.Query().Where(zarquote.TimestampGTE(time.Now().Add(-time.Hour * 24))).All(context.Background())
			if err != nil {
				log.Error("Can't get ZAR rates", err)
				return
			}

			//
			// Somewhere here we should build charts
			//

			LatestQuotes["ZAR"] = LatestQuote{
				Price:     query[len(query)-1:][0].Price,
				Timestamp: query[len(query)-1:][0].Timestamp,
			}
		}

	}

	log.Info("Cache updated!")
}

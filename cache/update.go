package cache

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"main/ent"
	"main/ent/usdquote"
	"os"
	"time"
)

func UpdateCache() {
	log.Info("Updating cache...")
	client, err := ent.Open("mysql", fmt.Sprint(os.Getenv("MDB_LOGIN"), ":", os.Getenv("MDB_PASSWORD"), "@tcp(", os.Getenv("MDB_HOST"), ":", os.Getenv("MDB_PORT"), ")/charts?parseTime=True"))
	if err != nil {
		log.Error("Can't connect to MySQL", err)
		return
	}
	defer client.Close()

	usd, err := client.USDQuote.Query().Where(usdquote.TimestampGTE(time.Now().Add(-time.Hour * 24))).All(context.Background())
	if err != nil {
		log.Error("Can't get USD rates", err)
		return
	}
	USDLatest = LatestQuote{
		Price:     usd[len(usd)-1:][0].Price,
		Timestamp: usd[len(usd)-1:][0].Timestamp,
	}
	log.Info("Cache updated!")
}

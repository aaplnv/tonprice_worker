package cache

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"main/ent"
	"os"
	"time"
)

func UpdateCache() {
	client, err := ent.Open("mysql", fmt.Sprint(os.Getenv("MDB_LOGIN"), ":", os.Getenv("MDB_PASSWORD"), "@tcp(", os.Getenv("MDB_HOST"), ":", os.Getenv("MDB_PORT"), ")/charts?parseTime=True"))
	if err != nil {
		log.Error("Can't connect to MySQL", err)
		return
	}
	defer client.Close()

	usd, err := client.USDQuote.Query().Order(ent.Desc()).First(context.Background())
	if err != nil {
		log.Error("Can't get USD rates", err)
		return
	}
	USDLatest = LatestQuote{
		Price:     usd.Price,
		Timestamp: time.Time{},
	}
}

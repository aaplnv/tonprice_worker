package database

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/telebot.v3/layout"
	"main/ent"
	"main/ent/user"
)

func UpdateUser(telegramID int64, all string, active string) error {
	lt, err := layout.NewDefault("settings.yml", "default")
	if err != nil {
		log.Fatal(err)
	}
	client, err := ent.Open("mysql", fmt.Sprint(lt.String("mariadb.login"), ":", lt.String("mariadb.password"), "@tcp(", lt.String("mariadb.host"), ":", lt.String("mariadb.port"), ")/testone?parseTime=True"))
	if err != nil {
		log.Error("Can't connect to MySQL", err)
	}
	defer client.Close()

	_, err = client.User.Update().Where(user.TelegramId(telegramID)).SetAllStables(all).SetActiveStable(active).Save(context.Background())

	if err != nil {
		log.Error(err)
	}
	return err
}

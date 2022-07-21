package database

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/telebot.v3/layout"
	"main/ent"
	"main/ent/user"
	"time"
)

func GetUser(id int64) *ent.User {
	lt, err := layout.NewDefault("settings.yml", "default")
	if err != nil {
		log.Fatal(err)
	}
	client, err := ent.Open("mysql", fmt.Sprint(lt.String("mariadb.login"), ":", lt.String("mariadb.password"), "@tcp(", lt.String("mariadb.host"), ":", lt.String("mariadb.port"), ")/test-one?parseTime=True"))
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

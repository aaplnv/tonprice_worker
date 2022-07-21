package currencies

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/telebot.v3/layout"
	"main/database"
	"main/ent"
	"strings"
)

func (s *settings) IsAdded(stable string) bool {
	if getIndex(len(s.All), func(i int) bool {
		return s.All[i] == stable
	}) == -1 {
		return false
	} else {
		return true
	}
}

func (s *settings) AddStable(stable string) {
	s.All = append(s.All, stable)
	s.Save()
}

func (s *settings) Save() {
	database.UpdateUser(s.TelegramId, strings.Join(s.All, " "))
}

func (s *settings) RemoveStable(stable string) {
	index := getIndex(len(s.All), func(i int) bool {
		return s.All[i] == stable
	})

	ret := make([]string, 0)
	ret = append(ret, s.All[:index]...)
	s.All = append(ret, s.All[index+1:]...)

	s.Save()
}

func Create(user *ent.User) (s settings) {
	saveRequired := false
	s.TelegramId = user.TelegramId

	for _, stable := range strings.Split(user.AllStables, " ") {
		if stable == "" {
			continue
			saveRequired = true
		}

		lt, err := layout.NewDefault("settings.yml", "default")
		if err != nil {
			log.Fatal(err)
		}

		if getIndex(len(lt.Strings("currencies")), func(i int) bool {
			return lt.Strings("currencies")[i] == stable
		}) == -1 {
			continue
			saveRequired = true
		}

		s.All = append(s.All, stable)
	}

	if saveRequired {
		s.Save()
	}
	return
}

func (s *settings) NextTicker(activeTicker string) string {
	index := getIndex(len(s.All), func(i int) bool {
		return s.All[i] == activeTicker
	})

	if (index + 1) == len(s.All) {
		return s.All[0]
	} else {
		return s.All[index+1]
	}
}

type settings struct {
	TelegramId int64
	All        []string
}

func getIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

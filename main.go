package main

import (
	"main/telegram"
	"sync"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var wg sync.WaitGroup

func init() {
	err := godotenv.Load("settings.env")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.Info("Starting application...")
	wg.Add(1)
	go func() {
		defer wg.Done()
		telegram.StartBot()
	}()
	wg.Wait()
}

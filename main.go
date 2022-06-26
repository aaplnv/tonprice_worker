package main

import (
	"main/cache"
	"main/telegram"
	"sync"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var wg sync.WaitGroup

func init() {
	// Set up logging level here
	log.SetLevel(log.InfoLevel)

	err := godotenv.Load("settings.env")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.Info("Starting application...")
	cache.UpdateCache()
	wg.Add(1)
	go func() {
		defer wg.Done()
		telegram.StartBot()
	}()
	wg.Wait()
}

package main

import (
	"github.com/jasonlvhit/gocron"
	"main/cache"
	"main/telegram"
	"sync"

	log "github.com/sirupsen/logrus"
)

var wg sync.WaitGroup

func init() {
	// Set up logging level here
	log.SetLevel(log.InfoLevel)
}

func main() {
	log.Info("Starting application...")
	cache.UpdateCache()

	wg.Add(2)
	go func() {
		defer wg.Done()
		telegram.StartBot()
	}()
	gocron.Every(10).Seconds().Do(cache.UpdateCache)
	go gocron.Start()
	wg.Wait()
}

package main

import (
	"main/telegram"
	"sync"

	log "github.com/sirupsen/logrus"
)

var wg sync.WaitGroup

func main() {
	log.Info("Starting application...")
	wg.Add(1)
	go func() {
		defer wg.Done()
		telegram.StartBot()
	}()
	wg.Wait()
}

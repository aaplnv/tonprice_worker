package main

import (
	"main/telegram"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go func() {
		defer wg.Done()
		telegram.StartBot()
	}()
	wg.Wait()
}

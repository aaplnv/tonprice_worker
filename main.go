package main

import (
	"fmt"
	"main/markets"
)

func main() {
	fmt.Println("TON price:", markets.GetPrice())
}

package cache

import (
	"time"
)

var LatestQuotes = map[string]LatestQuote{}

type LatestQuote struct {
	Price     float64
	Timestamp time.Time
}

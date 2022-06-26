package cache

import (
	"time"
)

var USDLatest = LatestQuote{}
var RUBLatest = LatestQuote{}

type LatestQuote struct {
	Price     float64
	Timestamp time.Time
}

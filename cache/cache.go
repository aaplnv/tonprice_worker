package cache

import (
	"time"
)

var AEDLatest = LatestQuote{}
var ARSLatest = LatestQuote{}
var AUDLatest = LatestQuote{}
var BHDLatest = LatestQuote{}
var BRLLatest = LatestQuote{}
var BTCLatest = LatestQuote{}
var CADLatest = LatestQuote{}
var CHFLatest = LatestQuote{}
var CLPLatest = LatestQuote{}
var CNYLatest = LatestQuote{}
var CZKLatest = LatestQuote{}
var EUROLatest = LatestQuote{}
var GBPLatest = LatestQuote{}
var HKDLatest = LatestQuote{}
var HUFLatest = LatestQuote{}
var IDRLatest = LatestQuote{}
var ILSLatest = LatestQuote{}
var INRLatest = LatestQuote{}
var JPYLatest = LatestQuote{}
var MXNLatest = LatestQuote{}
var NOKLatest = LatestQuote{}
var NZDLatest = LatestQuote{}
var PKRLatest = LatestQuote{}
var PLNLatest = LatestQuote{}
var RUBLatest = LatestQuote{}
var SARLatest = LatestQuote{}
var SEKLatest = LatestQuote{}
var TRYLatest = LatestQuote{}
var TWDLatest = LatestQuote{}
var UAHLatest = LatestQuote{}
var USDLatest = LatestQuote{}
var ZARLatest = LatestQuote{}

var LatestQuotes = map[string]LatestQuote{}

type LatestQuote struct {
	Price     float64
	Timestamp time.Time
}

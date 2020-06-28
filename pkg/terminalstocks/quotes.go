// Copyright (c) 2013-2016 by Michael Dvorkin. All Rights Reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package TerminalStocks

import (
	//"bytes"
	//"encoding/json"
	"fmt"
	//"reflect"
	//"strconv"
)

type APISourceType int32

const (
	API_VENDOR_YAHOO APISourceType = iota
	API_VENDOR_QQ
	API_VENDOR_SINA
	API_VENDOR_NETEASE
	API_VENDOR_EASTMONEY
	API_VENDOR_LIMITUP_EASTMONEY
	API_VENDOR_LHB_EASTMONEY
)

// Stock stores quote information for the particular stock ticker. The data
// for all the fields except 'Advancing' is fetched using Yahoo market API.
type Stock struct {
	Ticker     string `json:"symbol"` // Stock ticker.
	Gid        string `json:"gid"`
	LastTrade  string `json:"regularMarketPrice"`          // l1: last trade.
	Change     string `json:"regularMarketChange"`         // c6: change real time.
	ChangePct  string `json:"regularMarketChangePercent"`  // k2: percent change real time.
	Open       string `json:"regularMarketOpen"`           // o: market open price.
	Low        string `json:"regularMarketDayLow"`         // g: day's low.
	High       string `json:"regularMarketDayHigh"`        // h: day's high.
	Low52      string `json:"fiftyTwoWeekLow"`             // j: 52-weeks low.
	High52     string `json:"fiftyTwoWeekHigh"`            // k: 52-weeks high.
	Volume     string `json:"regularMarketVolume"`         // v: volume.
	AvgVolume  string `json:"averageDailyVolume10Day"`     // a2: average volume.
	PeRatio    string `json:"trailingPE"`                  // r2: P/E ration real time.
	PeRatioX   string `json:"trailingPE"`                  // r: P/E ration (fallback when real time is N/A).
	Dividend   string `json:"trailingAnnualDividendRate"`  // d: dividend.
	Yield      string `json:"trailingAnnualDividendYield"` // y: dividend yield.
	MarketCap  string `json:"marketCap"`                   // j3: market cap real time.
	MarketCapX string `json:"marketCap"`                   // j1: market cap (fallback when real time is N/A).
	Advancing  bool   // True when change is >= $0.
}

// Quotes stores relevant pointers as well as the array of stock quotes for
// the tickers we are tracking.
type Quotes struct {
	market  *Market       // Pointer to Market.
	profile *Profile      // Pointer to Profile.
	vendor  APISourceType // Pointer to Profile.
	stocks  []Stock       // Array of stock quote data.
	errors  string        // Error string if any.
}

// Sets the initial values and returns new Quotes struct.
func NewQuotes(market *Market, profile *Profile) *Quotes {
	return &Quotes{
		market:  market,
		profile: profile,
		errors:  ``,
	}
}

// Fetch the latest stock quotes and parse raw fetched data into array of
// []Stock structs.
func (quotes *Quotes) Fetch() (self *Quotes) {
	self = quotes // <-- This ensures we return correct quotes after recover() from panic().
	switch quotes.market.vendor {
	case API_VENDOR_YAHOO:
		quotes.FetchYahoo()
	case API_VENDOR_QQ:
		quotes.FetchQQ()
	case API_VENDOR_SINA:
		quotes.FetchSina()
	case API_VENDOR_NETEASE:
		quotes.FetchNetease()
	case API_VENDOR_EASTMONEY:
		quotes.FetchEastmoney()
	case API_VENDOR_LIMITUP_EASTMONEY:
		quotes.FetchLimitupEastmoney()
	default:
		quotes.errors = fmt.Sprintf("\n\n\n\nError Not Support : %v", quotes.vendor)
	}

	return quotes
}

// Ok returns two values: 1) boolean indicating whether the error has occured,
// and 2) the error text itself.
func (quotes *Quotes) Ok() (bool, string) {
	return quotes.errors == ``, quotes.errors
}

// AddTickers saves the list of tickers and refreshes the stock data if new
// tickers have been added. The function gets called from the line editor
// when user adds new stock tickers.
func (quotes *Quotes) AddTickers(tickers []string) (added int, err error) {
	if added, err = quotes.profile.AddTickers(tickers); err == nil && added > 0 {
		quotes.stocks = nil // Force fetch.
	}
	return
}

// RemoveTickers saves the list of tickers and refreshes the stock data if some
// tickers have been removed. The function gets called from the line editor
// when user removes existing stock tickers.
func (quotes *Quotes) RemoveTickers(tickers []string) (removed int, err error) {
	if removed, err = quotes.profile.RemoveTickers(tickers); err == nil && removed > 0 {
		quotes.stocks = nil // Force fetch.
	}
	return
}

// isReady returns true if we haven't fetched the quotes yet *or* the stock
// market is still open and we might want to grab the latest quotes. In both
// cases we make sure the list of requested tickers is not empty.
func (quotes *Quotes) isReady() bool {
	return (quotes.stocks == nil || !quotes.market.IsClosed) && len(quotes.profile.Tickers) > 0
}

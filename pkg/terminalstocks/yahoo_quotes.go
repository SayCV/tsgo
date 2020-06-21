// Copyright (c) 2013-2016 by Michael Dvorkin. All Rights Reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package TerminalStocks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

const quotesURLv7 = `https://query1.finance.yahoo.com/v7/finance/quote?symbols=%s`
const quotesURLv7QueryParts = `&range=1d&interval=5m&indicators=close&includeTimestamps=false&includePrePost=false&corsDomain=finance.yahoo.com&.tsrc=finance`

const noDataIndicator = `N/A`

// Fetch the latest stock quotes and parse raw fetched data into array of
// []Stock structs.
func (quotes *Quotes) FetchYahoo() (self *Quotes) {
	self = quotes // <-- This ensures we return correct quotes after recover() from panic().
	if quotes.isReady() {
		defer func() {
			if err := recover(); err != nil {
				quotes.errors = fmt.Sprintf("\n\n\n\nError fetching stock quotes...\n%s", err)
			}
		}()

		url := fmt.Sprintf(quotesURLv7, strings.Join(quotes.profile.Tickers, `,`))
		response, err := http.Get(url + quotesURLv7QueryParts)
		if err != nil {
			panic(err)
		}

		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}

		quotes.parse2(body)
	}

	return quotes
}

// this will parse the json objects
func (quotes *Quotes) parse2(body []byte) (*Quotes, error) {
	// response -> quoteResponse -> result|error (array) -> map[string]interface{}
	// Stocks has non-int things
	// d := map[string]map[string][]Stock{}
	// some of these are numbers vs strings
	// d := map[string]map[string][]map[string]string{}
	d := map[string]map[string][]map[string]interface{}{}
	err := json.Unmarshal(body, &d)
	if err != nil {
		return nil, err
	}
	results := d["quoteResponse"]["result"]

	quotes.stocks = make([]Stock, len(results))
	for i, raw := range results {
		result := map[string]string{}
		for k, v := range raw {
			switch v.(type) {
			case string:
				result[k] = v.(string)
			case float64:
				result[k] = float2Str(v.(float64))
			default:
				result[k] = fmt.Sprintf("%v", v)
			}

		}
		quotes.stocks[i].Ticker = result["symbol"]
		quotes.stocks[i].LastTrade = result["regularMarketPrice"]
		quotes.stocks[i].Change = result["regularMarketChange"]
		quotes.stocks[i].ChangePct = result["regularMarketChangePercent"]
		quotes.stocks[i].Open = result["regularMarketOpen"]
		quotes.stocks[i].Low = result["regularMarketDayLow"]
		quotes.stocks[i].High = result["regularMarketDayHigh"]
		quotes.stocks[i].Low52 = result["fiftyTwoWeekLow"]
		quotes.stocks[i].High52 = result["fiftyTwoWeekHigh"]
		quotes.stocks[i].Volume = result["regularMarketVolume"]
		quotes.stocks[i].AvgVolume = result["averageDailyVolume10Day"]
		quotes.stocks[i].PeRatio = result["trailingPE"]
		// TODO calculate rt
		quotes.stocks[i].PeRatioX = result["trailingPE"]
		quotes.stocks[i].Dividend = result["trailingAnnualDividendRate"]
		quotes.stocks[i].Yield = result["trailingAnnualDividendYield"]
		quotes.stocks[i].MarketCap = result["marketCap"]
		// TODO calculate rt?
		quotes.stocks[i].MarketCapX = result["marketCap"]

		/*
			fmt.Println(i)
			fmt.Println("-------------------")
			for k, v := range result {
				fmt.Println(k, v)
			}
			fmt.Println("-------------------")
		*/
		adv, err := strconv.ParseFloat(quotes.stocks[i].Change, 64)
		if err == nil {
			quotes.stocks[i].Advancing = adv >= 0.0
		}
	}
	return quotes, nil
}

// Use reflection to parse and assign the quotes data fetched using the Yahoo
// market API.
func (quotes *Quotes) parse(body []byte) *Quotes {
	lines := bytes.Split(body, []byte{'\n'})
	quotes.stocks = make([]Stock, len(lines))
	//
	// Get the total number of fields in the Stock struct. Skip the last
	// Advancing field which is not fetched.
	//
	fieldsCount := reflect.ValueOf(quotes.stocks[0]).NumField() - 1
	//
	// Split each line into columns, then iterate over the Stock struct
	// fields to assign column values.
	//
	for i, line := range lines {
		columns := bytes.Split(bytes.TrimSpace(line), []byte{','})
		for j := 0; j < fieldsCount; j++ {
			// ex. quotes.stocks[i].Ticker = string(columns[0])
			reflect.ValueOf(&quotes.stocks[i]).Elem().Field(j).SetString(string(columns[j]))
		}
		//
		// Try realtime value and revert to the last known if the
		// realtime is not available.
		//
		if quotes.stocks[i].PeRatio == noDataIndicator && quotes.stocks[i].PeRatioX != noDataIndicator {
			quotes.stocks[i].PeRatio = quotes.stocks[i].PeRatioX
		}
		if quotes.stocks[i].MarketCap == noDataIndicator && quotes.stocks[i].MarketCapX != noDataIndicator {
			quotes.stocks[i].MarketCap = quotes.stocks[i].MarketCapX
		}
		//
		// Stock is advancing if the change is not negative (i.e. $0.00
		// is also "advancing").
		//
		quotes.stocks[i].Advancing = (quotes.stocks[i].Change[0:1] != `-`)
	}

	return quotes
}

//-----------------------------------------------------------------------------
func sanitize(body []byte) []byte {
	return bytes.Replace(bytes.TrimSpace(body), []byte{'"'}, []byte{}, -1)
}

func float2Str(v float64) string {
	unit := ""
	switch {
	case v > 1.0e12:
		v = v / 1.0e12
		unit = "T"
	case v > 1.0e9:
		v = v / 1.0e9
		unit = "B"
	case v > 1.0e6:
		v = v / 1.0e6
		unit = "M"
	case v > 1.0e5:
		v = v / 1.0e3
		unit = "K"
	default:
		unit = ""
	}
	// parse
	return fmt.Sprintf("%0.3f%s", v, unit)
}

// Copyright (c) 2013-2016 by Michael Dvorkin. All Rights Reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package TerminalStocks

import (
	"strconv"
	"strings"

	util "github.com/saycv/tsgo/pkg/utils"
)

// Fetch downloads HTML page from the 'marketURL', parses it, and stores resulting data
// in internal hashes. If download or data parsing fails Fetch populates 'market.errors'.
func (market *Market) FetchQQ() (self *Market) {
	self = market // <-- This ensures we return correct market after recover() from panic().
	//defer func() {
	//	if err := recover(); err != nil {
	//		market.errors = fmt.Sprintf("Error fetching market data...\n%s", err)
	//	}
	//}()

	codes := []string{"sh000001", "sz399001", "sh000300", "sz399006"}
	codes = util.StockWithPrefix(codes)
	//fmt.Println(code)
	results := GetRealtime(strings.Join(codes, ","))
	//fmt.Println(results)

	realTime := results[0]

	market.Szzs[`change`] = strconv.FormatFloat(realTime.Change, 'f', -1, 64)
	market.Szzs[`latest`] = strconv.FormatFloat(realTime.NowPri, 'f', -1, 64)
	market.Szzs[`percent`] = strconv.FormatFloat(realTime.ChangePer, 'f', -1, 64)

	realTime = results[1]

	market.Szcz[`change`] = strconv.FormatFloat(realTime.Change, 'f', -1, 64)
	market.Szcz[`latest`] = strconv.FormatFloat(realTime.NowPri, 'f', -1, 64)
	market.Szcz[`percent`] = strconv.FormatFloat(realTime.ChangePer, 'f', -1, 64)

	realTime = results[2]

	market.Hs300[`change`] = strconv.FormatFloat(realTime.Change, 'f', -1, 64)
	market.Hs300[`latest`] = strconv.FormatFloat(realTime.NowPri, 'f', -1, 64)
	market.Hs300[`percent`] = strconv.FormatFloat(realTime.ChangePer, 'f', -1, 64)

	realTime = results[3]

	market.Cybz[`change`] = strconv.FormatFloat(realTime.Change, 'f', -1, 64)
	market.Cybz[`latest`] = strconv.FormatFloat(realTime.NowPri, 'f', -1, 64)
	market.Cybz[`percent`] = strconv.FormatFloat(realTime.ChangePer, 'f', -1, 64)

	return market
}

package TerminalStocks

import (
	"fmt"
	"testing"
	"time"

	//TerminalStocks "github.com/saycv/tsgo/pkg/terminalstocks"
	util "github.com/saycv/tsgo/pkg/utils"
)

func TestEastmoneyQuotesCase2(t *testing.T) {
	now := time.Now()
	year, month, day := now.Date()
	yearstart := fmt.Sprintf("%04d%02d%02d", year, month, day)

	//t.Log(util.IsTradeDay(yearstart))
	//dt.Log(util.GetLatestTradeDay(yearstart))
	lhblist := GetLhbEastmoney(util.GetLatestTradeDay(yearstart))
	t.Log(lhblist)

}

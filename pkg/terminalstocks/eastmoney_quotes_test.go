package TerminalStocks

import (
	//"fmt"
	//"strings"

	//"io/ioutil"
	//"strings"
	"testing"

	util "github.com/saycv/tsgo/pkg/utils"

	//"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEastmoneyQuotesCase0(t *testing.T) {
	market := NewMarket(API_VENDOR_EASTMONEY)
	profile := NewProfile(API_VENDOR_EASTMONEY)

	code := []string{"600519", "601318", "601066", "002958", "000878", "600121", "603121"}
	profile.Tickers = util.StockWithPrefixEastmoney(code)

	quotes := NewQuotes(market, profile)
	require.NotNil(t, quotes)

	results := quotes.FetchLimitupEastmoney()
	t.Log(results)

}

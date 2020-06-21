package TerminalStocks

import (
	//"fmt"
	//"io/ioutil"
	//"strings"
	"testing"

	util "github.com/saycv/tsgo/pkg/utils"

	//"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNeteaseQuotesCase0(t *testing.T) {
	market := NewMarket(API_VENDOR_NETEASE)
	profile := NewProfile()

	code := []string{"600519", "601318", "601066", "002958", "000878", "600121", "603121"}
	profile.Tickers = util.StockWithPrefixNetease(code)

	quotes := NewQuotes(market, profile)
	require.NotNil(t, quotes)

	results := quotes.FetchNetease()
	t.Log(results)

}

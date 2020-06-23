package TerminalStocks

import (
	"fmt"
	"strings"

	//"io/ioutil"
	//"strings"
	"testing"

	util "github.com/saycv/tsgo/pkg/utils"

	//"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNeteaseQuotesCase0(t *testing.T) {
	market := NewMarket(API_VENDOR_NETEASE)
	profile := NewProfile(API_VENDOR_NETEASE)

	code := []string{"600519", "601318", "601066", "002958", "000878", "600121", "603121"}
	profile.Tickers = util.StockWithPrefixNetease(code)

	quotes := NewQuotes(market, profile)
	require.NotNil(t, quotes)

	results := quotes.FetchNetease()
	t.Log(results)

}

func TestNeteaseQuotesCase2(t *testing.T) {
	token := "中国"
	t.Log(strings.Join(strings.Split(token, ""), " "))
	for i, char := range token {
		t.Log(i)
		t.Log(char)
		t.Log(fmt.Sprintf("%c", char))
	}

}

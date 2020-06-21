package TerminalStocks

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/ShawnRong/tushare-go"
	util "github.com/saycv/tsgo/pkg/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	TUSHARE_TOKEN = "4c010225d485a8db581030b9c04f634d14d3b2bd92fa2e3546a77bbe"
)

func TestQQQuotesCase0(t *testing.T) {

	c := tushare.New(TUSHARE_TOKEN)
	// 参数
	params := make(map[string]string)
	// 字段
	var fields []string = []string{
		"ts_code", "symbol", "name", "area", "industry", "list_date",
	}
	// 根据api 请求对应的接口
	data, _ := c.StockBasic(params, fields)

	t.Log(data)
}

func TestQQQuotesCase1(t *testing.T) {
	code := []string{"600519", "300519"}
	code = util.StockWithPrefix(code)
	fmt.Println(code)
	realTime := GetRealtime(strings.Join(code, ","))
	//pk := GetPK(code)
	//funFlow := GetFundFlow(code)
	//info := GetInfo(code)
	fmt.Println(util.JsonEncodeS(realTime))
	//fmt.Println(util.JsonEncodeS(pk))
	//fmt.Println(util.JsonEncodeS(funFlow))
	//fmt.Println(util.JsonEncodeS(info))
	//	list := GetDaily(code, 15)
	//list := GetWeekly(code)
	//fmt.Println(util.JsonEncodeS(list))
}

func TestQQQuotesCase2(t *testing.T) {
	market := NewMarket(API_VENDOR_QQ)
	profile := NewProfile()

	profile.Tickers = []string{"GOOG", "BA"}

	quotes := NewQuotes(market, profile)
	require.NotNil(t, quotes)

	data, err := ioutil.ReadFile("./yahoo_quotes_sample.json")
	require.Nil(t, err)
	require.NotNil(t, data)

	require.True(t, quotes.isReady())
	//quotes.Fetch(data)
	_, err = quotes.parse2(data)
	assert.NoError(t, err)

	require.Equal(t, 2, len(quotes.stocks))
	assert.Equal(t, "BA", quotes.stocks[0].Ticker)
	assert.Equal(t, "331.76", quotes.stocks[0].LastTrade)
	assert.Equal(t, "GOOG", quotes.stocks[1].Ticker)
	assert.Equal(t, "1214.38", quotes.stocks[1].LastTrade)
}

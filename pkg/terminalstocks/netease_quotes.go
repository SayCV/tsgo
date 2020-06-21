package TerminalStocks

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"net/http"
	//"regexp"
	"strconv"
	"strings"

	//log "github.com/sirupsen/logrus"
	"github.com/axgle/mahonia"
	//"github.com/axgle/pinyin"
	util "github.com/saycv/tsgo/pkg/utils"
)

const (
	URL_NETEASE_REAL_TIME = "http://api.money.126.net/data/feed/%s,money.api"
	URL_NETEASE_FUND_FLOW = "http://qt.gtimg.cn/q=ff_%s"
	URL_NETEASE_PK        = "http://qt.gtimg.cn/q=s_pk%s"
	URL_NETEASE_INFO      = "http://qt.gtimg.cn/q=s_%s"
	URL_NETEASE_DAILY     = "http://data.gtimg.cn/flashdata/hushen/daily/%v/%s.js"
	URL_NETEASE_WEEKLY    = "http://data.gtimg.cn/flashdata/hushen/weekly/%s.js"
)

func GetRealtimeNetease(code string) []*RealTimeData {
	//fmt.Println(fmt.Sprintf(URL_NETEASE_REAL_TIME, code))
	resp, err := http.DefaultClient.Get(fmt.Sprintf(URL_NETEASE_REAL_TIME, code))
	checkErr(err)
	if resp == nil || resp.StatusCode != http.StatusOK {
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	res := mahonia.NewDecoder("gbk").ConvertString(string(body))

	var dataList []*RealTimeData
	//fmt.Println(res)
	results := strings.Split(res, ";")
	for _, raw := range results {
		if !strings.Contains(raw, "~") {
			break
		}
		dataArray := strings.Split(raw, "~")
		//fmt.Println(strings.Join(dataArray, "\n"))
		data := new(RealTimeData)
		data.Name = strings.Replace(dataArray[1], " ", "", -1)
		data.Gid = dataArray[2]
		data.NowPri, _ = strconv.ParseFloat(dataArray[3], 64)
		data.YestClosePri, _ = strconv.ParseFloat(dataArray[4], 64)
		data.OpeningPri, _ = strconv.ParseFloat(dataArray[5], 64)
		data.TraNumber, _ = strconv.ParseInt(dataArray[6], 10, 64)
		data.Outter, _ = strconv.ParseInt(dataArray[7], 10, 64)
		data.Inner, _ = strconv.ParseInt(dataArray[8], 10, 64)
		data.BuyOnePri, _ = strconv.ParseFloat(dataArray[9], 64)
		data.BuyOne, _ = strconv.ParseInt(dataArray[10], 10, 64)
		data.BuyTwoPri, _ = strconv.ParseFloat(dataArray[11], 64)
		data.BuyTwo, _ = strconv.ParseInt(dataArray[12], 10, 64)
		data.BuyThreePri, _ = strconv.ParseFloat(dataArray[13], 64)
		data.BuyThree, _ = strconv.ParseInt(dataArray[14], 10, 64)
		data.BuyFourPri, _ = strconv.ParseFloat(dataArray[15], 64)
		data.BuyFour, _ = strconv.ParseInt(dataArray[16], 10, 64)
		data.BuyFivePri, _ = strconv.ParseFloat(dataArray[17], 64)
		data.BuyFive, _ = strconv.ParseInt(dataArray[18], 10, 64)
		data.SellOnePri, _ = strconv.ParseFloat(dataArray[19], 64)
		data.SellOne, _ = strconv.ParseInt(dataArray[20], 10, 64)
		data.SellTwoPri, _ = strconv.ParseFloat(dataArray[21], 64)
		data.SellTwo, _ = strconv.ParseInt(dataArray[22], 10, 64)
		data.SellThreePri, _ = strconv.ParseFloat(dataArray[23], 64)
		data.SellThree, _ = strconv.ParseInt(dataArray[24], 10, 64)
		data.SellFourPri, _ = strconv.ParseFloat(dataArray[25], 64)
		data.SellFour, _ = strconv.ParseInt(dataArray[26], 10, 64)
		data.SellFivePri, _ = strconv.ParseFloat(dataArray[27], 64)
		data.SellFive, _ = strconv.ParseInt(dataArray[28], 10, 64)
		data.Time = dataArray[30]
		data.Change, _ = strconv.ParseFloat(dataArray[31], 64)
		data.ChangePer, _ = strconv.ParseFloat(dataArray[32], 64)
		data.TodayMax, _ = strconv.ParseFloat(dataArray[33], 64)
		data.TodayMin, _ = strconv.ParseFloat(dataArray[34], 64)
		data.TradeCount, _ = strconv.ParseInt(dataArray[36], 10, 64)
		data.TradeAmont, _ = strconv.ParseInt(dataArray[37], 10, 64)
		data.ChangeRate, _ = strconv.ParseFloat(dataArray[38], 64)
		data.PERatio, _ = strconv.ParseFloat(dataArray[39], 64)
		data.MaxMinChange, _ = strconv.ParseFloat(dataArray[43], 64)
		data.MarketAmont, _ = strconv.ParseFloat(dataArray[44], 64)
		data.TotalAmont, _ = strconv.ParseFloat(dataArray[45], 64)
		data.PBRatio, _ = strconv.ParseFloat(dataArray[46], 64)
		data.HighPri, _ = strconv.ParseFloat(dataArray[47], 64)
		data.LowPri, _ = strconv.ParseFloat(dataArray[48], 64)

		dataList = append(dataList, data)
	}
	return dataList
}

func GetPKNetease(code string) *PKData {
	resp, err := http.DefaultClient.Get(fmt.Sprintf(URL_NETEASE_PK, code))
	checkErr(err)
	if resp == nil || resp.StatusCode != http.StatusOK {
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	res := mahonia.NewDecoder("gbk").ConvertString(string(body))
	dataArray := strings.Split(res, "~")
	data := new(PKData)
	data.BuyBig, _ = strconv.ParseFloat(strings.Split(dataArray[0], "\"")[1], 64)
	data.BuySmall, _ = strconv.ParseFloat(dataArray[1], 64)
	data.SellBig, _ = strconv.ParseFloat(dataArray[2], 64)
	data.SellSmall, _ = strconv.ParseFloat(strings.Split(dataArray[3], "\"")[0], 64)
	return data

}
func GetFundFlowNetease(code string) *FundFlow {
	resp, err := http.DefaultClient.Get(fmt.Sprintf(URL_NETEASE_FUND_FLOW, code))
	checkErr(err)
	if resp == nil || resp.StatusCode != http.StatusOK {
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	res := mahonia.NewDecoder("gbk").ConvertString(string(body))
	dataArray := strings.Split(res, "~")
	data := new(FundFlow)
	data.Gid = code
	data.BigIn, _ = strconv.ParseFloat(dataArray[1], 64)
	data.BigOut, _ = strconv.ParseFloat(dataArray[2], 64)
	data.SmallIn, _ = strconv.ParseFloat(dataArray[5], 64)
	data.SmallOut, _ = strconv.ParseFloat(dataArray[6], 64)
	data.Name = dataArray[12]
	data.Date = dataArray[13]
	return data
}

func GetInfoNetease(code string) *StockInfo {
	resp, err := http.DefaultClient.Get(fmt.Sprintf(URL_NETEASE_INFO, code))
	checkErr(err)
	if resp == nil || resp.StatusCode != http.StatusOK {
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	res := mahonia.NewDecoder("gbk").ConvertString(string(body))
	dataArray := strings.Split(res, "~")
	data := new(StockInfo)
	data.Name = dataArray[1]
	data.Gid = dataArray[2]
	data.Price, _ = strconv.ParseFloat(dataArray[3], 64)
	data.Change, _ = strconv.ParseFloat(dataArray[4], 64)
	data.ChangePer, _ = strconv.ParseFloat(dataArray[5], 64)
	data.TradeCount, _ = strconv.ParseFloat(dataArray[6], 64)
	data.TradeAmont, _ = strconv.ParseFloat(dataArray[7], 64)
	data.TotalAmont, _ = strconv.ParseFloat(strings.Split(dataArray[9], "\"")[0], 64)
	return data
}

func GetDailyNetease(code string, year int) []*HistoryData {
	resp, err := http.DefaultClient.Get(fmt.Sprintf(URL_NETEASE_DAILY, year, code))
	checkErr(err)
	if resp == nil || resp.StatusCode != http.StatusOK {
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	res := string(body)
	dataArray := strings.Split(res, "\\n\\")
	list := []*HistoryData{}
	for index, str := range dataArray {
		if index == 0 || index == len(dataArray)-1 {

		} else {
			data := strings.Split(str, " ")
			entity := new(HistoryData)
			entity.Date = strings.Replace(data[0], "\n", "", -1)
			entity.Open, _ = strconv.ParseFloat(data[1], 64)
			entity.Close, _ = strconv.ParseFloat(data[2], 64)
			entity.Max, _ = strconv.ParseFloat(data[3], 64)
			entity.Min, _ = strconv.ParseFloat(data[4], 64)
			entity.Trade, _ = strconv.ParseFloat(data[5], 64)
			list = append(list, entity)
		}
	}
	return list
}

func GetWeeklyNetease(code string) []*HistoryData {
	resp, err := http.DefaultClient.Get(fmt.Sprintf(URL_NETEASE_WEEKLY, code))
	checkErr(err)
	if resp == nil || resp.StatusCode != http.StatusOK {
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	res := string(body)
	dataArray := strings.Split(res, "\\n\\")
	list := []*HistoryData{}
	for index, str := range dataArray {
		if index == 0 || index == len(dataArray)-1 {

		} else {
			data := strings.Split(str, " ")
			entity := new(HistoryData)
			entity.Date = strings.Replace(data[0], "\n", "", -1)
			entity.Open, _ = strconv.ParseFloat(data[1], 64)
			entity.Close, _ = strconv.ParseFloat(data[2], 64)
			entity.Max, _ = strconv.ParseFloat(data[3], 64)
			entity.Min, _ = strconv.ParseFloat(data[4], 64)
			entity.Trade, _ = strconv.ParseFloat(data[5], 64)
			list = append(list, entity)
		}
	}
	return list
}

// Fetch the latest stock quotes and parse raw fetched data into array of
// []Stock structs.
func (quotes *Quotes) FetchNetease() (self *Quotes) {
	self = quotes // <-- This ensures we return correct quotes after recover() from panic().
	if quotes.isReady() {
		defer func() {
			if err := recover(); err != nil {
				quotes.errors = fmt.Sprintf("\n\n\n\nError fetching stock quotes...\n%s", err)
			}
		}()

		code := util.StockWithPrefix(quotes.profile.Tickers)
		url := fmt.Sprintf(URL_NETEASE_REAL_TIME, strings.Join(code, ","))
		response, err := http.Get(url)
		if err != nil {
			panic(err)
		}

		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}

		quotes.parseNetease(body)
	}

	return quotes
}

// this will parse the json objects
func (quotes *Quotes) parseNetease(body []byte) (*Quotes, error) {
	// response -> quoteResponse -> result|error (array) -> map[string]interface{}
	// Stocks has non-int things
	// d := map[string]map[string][]Stock{}
	// some of these are numbers vs strings
	// d := map[string]map[string][]map[string]string{}
	d := map[string]interface{}{}
	//reusltString := strings.Trim(string(body), "_ntes_quote_callback();")
	reusltString := strings.TrimLeft(string(body), "_ntes_quote_callback(")
	reusltString = strings.TrimRight(reusltString, ");")
	err := json.Unmarshal([]byte(reusltString), &d)
	if err != nil {
		return nil, err
	}
	results := d
	//fmt.Println(len(results))
	//log.Info(results)

	i := 0
	quotes.stocks = make([]Stock, len(results))
	for _, raw := range results {
		//log.Info(j, raw)
		result := map[string]string{}
		for k, v := range raw.(map[string]interface{}) {
			switch v.(type) {
			case string:
				result[k] = v.(string)
			case float64:
				result[k] = float2Str(v.(float64))
			default:
				result[k] = fmt.Sprintf("%v", v)
			}

		}
		
		yestclose, _ := strconv.ParseFloat(result["yestclose"], 64)
		now, _ := strconv.ParseFloat(result["price"], 64)

		quotes.stocks[i].Ticker = result["name"]
		quotes.stocks[i].LastTrade = result["price"]
		quotes.stocks[i].Change = fmt.Sprintf("%.2f", now-yestclose)
		quotes.stocks[i].ChangePct = fmt.Sprintf("%.2f", 100.0*(now-yestclose)/yestclose)
		quotes.stocks[i].Open = result["open"]
		quotes.stocks[i].Low = result["low"]
		quotes.stocks[i].High = result["high"]
		quotes.stocks[i].Low52 = ""
		quotes.stocks[i].High52 = ""
		quotes.stocks[i].Volume = result["volume"]
		quotes.stocks[i].AvgVolume = ""
		quotes.stocks[i].PeRatio = ""
		// TODO calculate rt
		quotes.stocks[i].PeRatioX = ""
		quotes.stocks[i].Dividend = ""
		quotes.stocks[i].Yield = ""
		quotes.stocks[i].MarketCap = ""
		// TODO calculate rt?
		quotes.stocks[i].MarketCapX = ""

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
		i++
	}
	return quotes, nil
}

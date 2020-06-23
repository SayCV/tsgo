package TerminalStocks

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/axgle/mahonia"
	"github.com/axgle/pinyin"
	util "github.com/saycv/tsgo/pkg/utils"
)

//某只某年股票日行情
//http://data.gtimg.cn/flashdata/hushen/daily/15/sz000750.js
//某只某年股票周行情
//http://data.gtimg.cn/flashdata/hushen/weekly/sz000868.js
//实时行情
//http://qt.gtimg.cn/q=sz000858
//资金流向
//http://qt.gtimg.cn/q=ff_sz000858
//盘口分析
//http://qt.gtimg.cn/q=s_pksz000858
//简要信息
//http://qt.gtimg.cn/q=s_sz000858

const (
	URL_QQ_REAL_TIME = "http://qt.gtimg.cn/q=%s"
	URL_QQ_FUND_FLOW = "http://qt.gtimg.cn/q=ff_%s"
	URL_QQ_PK        = "http://qt.gtimg.cn/q=s_pk%s"
	URL_QQ_INFO      = "http://qt.gtimg.cn/q=s_%s"
	URL_QQ_DAILY     = "http://data.gtimg.cn/flashdata/hushen/daily/%v/%s.js"
	URL_QQ_WEEKLY    = "http://data.gtimg.cn/flashdata/hushen/weekly/%s.js"
)

type StockGetter struct {
	client *http.Client
}

func checkErr(err error) {}

func GetRealtime(code string) []*RealTimeData {
	//fmt.Println(fmt.Sprintf(URL_QQ_REAL_TIME, code))
	resp, err := http.DefaultClient.Get(fmt.Sprintf(URL_QQ_REAL_TIME, code))
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

func GetPK(code string) *PKData {
	resp, err := http.DefaultClient.Get(fmt.Sprintf(URL_QQ_PK, code))
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
func GetFundFlow(code string) *FundFlow {
	resp, err := http.DefaultClient.Get(fmt.Sprintf(URL_QQ_FUND_FLOW, code))
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

func GetInfo(code string) *StockInfo {
	resp, err := http.DefaultClient.Get(fmt.Sprintf(URL_QQ_INFO, code))
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

func GetDaily(code string, year int) []*HistoryData {
	resp, err := http.DefaultClient.Get(fmt.Sprintf(URL_QQ_DAILY, year, code))
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

func GetWeekly(code string) []*HistoryData {
	resp, err := http.DefaultClient.Get(fmt.Sprintf(URL_QQ_WEEKLY, code))
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
func (quotes *Quotes) FetchQQ() (self *Quotes) {
	self = quotes // <-- This ensures we return correct quotes after recover() from panic().
	if quotes.isReady() {
		//defer func() {
		//	if err := recover(); err != nil {
		//		quotes.errors = fmt.Sprintf("\n\n\n\nError fetching stock quotes...\n%s", err)
		//	}
		//}()

		code := util.StockWithPrefix(quotes.profile.Tickers)
		results := GetRealtime(strings.Join(code, ","))

		quotes.stocks = make([]Stock, len(results))
		for i, raw := range results {
			realTime := raw

			pyStr := pinyin.Convert(realTime.Name)
			reg, err := regexp.Compile("[a-z]+")
			if err != nil {
				log.Fatal(err)
			}
			pyCapStr := reg.ReplaceAllString(pyStr, "")
			if true {
				if false {
					quotes.stocks[i].Ticker = realTime.Name
				} else {
					tmp := ""
					for _, char := range realTime.Name {
						tmp = tmp + string(char) + string(char)
					}
					quotes.stocks[i].Ticker = tmp
				}
			} else {
				quotes.stocks[i].Ticker = pyCapStr
			}
			quotes.stocks[i].LastTrade = strconv.FormatFloat(realTime.NowPri, 'f', -1, 64)
			quotes.stocks[i].Change = strconv.FormatFloat(realTime.Change, 'f', -1, 64)
			quotes.stocks[i].ChangePct = strconv.FormatFloat(realTime.ChangePer, 'f', -1, 64)
			quotes.stocks[i].Open = strconv.FormatFloat(realTime.OpeningPri, 'f', -1, 64)
			quotes.stocks[i].Low = strconv.FormatFloat(realTime.TodayMin, 'f', -1, 64)
			quotes.stocks[i].High = strconv.FormatFloat(realTime.TodayMax, 'f', -1, 64)
			quotes.stocks[i].Low52 = ""
			quotes.stocks[i].High52 = ""
			quotes.stocks[i].Volume = strconv.Itoa(int(realTime.TradeAmont))
			quotes.stocks[i].AvgVolume = ""
			quotes.stocks[i].PeRatio = strconv.FormatFloat(realTime.PERatio, 'f', -1, 64)
			// TODO calculate rt
			quotes.stocks[i].PeRatioX = ""
			quotes.stocks[i].Dividend = ""
			quotes.stocks[i].Yield = ""
			quotes.stocks[i].MarketCap = ""
			// TODO calculate rt?
			quotes.stocks[i].MarketCapX = ""

			adv := realTime.Change
			quotes.stocks[i].Advancing = adv >= 0.0
		}
	}

	return quotes
}

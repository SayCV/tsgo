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

const (
	URL_SINA_REAL_TIME = "http://hq.sinajs.cn/list=%s"
	URL_SINA_FUND_FLOW = "http://qt.gtimg.cn/q=ff_%s"
	URL_SINA_PK        = "http://qt.gtimg.cn/q=s_pk%s"
	URL_SINA_INFO      = "http://qt.gtimg.cn/q=s_%s"
	URL_SINA_DAILY     = "http://image.sinajs.cn/newchart/%d/%s/n/"
	URL_SINA_WEEKLY    = "http://image.sinajs.cn/newchart/%s/n/"
)

func GetRealtimeSina(code string) []*RealTimeData {
	//fmt.Println(fmt.Sprintf(URL_SINA_REAL_TIME, code))
	resp, err := http.DefaultClient.Get(fmt.Sprintf(URL_SINA_REAL_TIME, code))
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
		if !strings.Contains(raw, ",") {
			break
		}
		dataArray := strings.Split(raw, ",")
		//fmt.Println(strings.Join(dataArray, "\n"))
		data := new(RealTimeData)
		name := strings.Split(dataArray[0], "\"")
		data.Name = strings.Replace(name[1], " ", "", -1)
		//data.Gid = dataArray[2]
		data.NowPri, _ = strconv.ParseFloat(dataArray[3], 64)
		data.YestClosePri, _ = strconv.ParseFloat(dataArray[2], 64)
		data.OpeningPri, _ = strconv.ParseFloat(dataArray[1], 64)
		data.TraNumber, _ = strconv.ParseInt(dataArray[8], 10, 64)
		//data.Outter, _ = strconv.ParseInt(dataArray[7], 10, 64)
		//data.Inner, _ = strconv.ParseInt(dataArray[8], 10, 64)
		data.BuyOnePri, _ = strconv.ParseFloat(dataArray[11], 64)
		data.BuyOne, _ = strconv.ParseInt(dataArray[10], 10, 64)
		data.BuyTwoPri, _ = strconv.ParseFloat(dataArray[13], 64)
		data.BuyTwo, _ = strconv.ParseInt(dataArray[12], 10, 64)
		data.BuyThreePri, _ = strconv.ParseFloat(dataArray[15], 64)
		data.BuyThree, _ = strconv.ParseInt(dataArray[14], 10, 64)
		data.BuyFourPri, _ = strconv.ParseFloat(dataArray[17], 64)
		data.BuyFour, _ = strconv.ParseInt(dataArray[16], 10, 64)
		data.BuyFivePri, _ = strconv.ParseFloat(dataArray[19], 64)
		data.BuyFive, _ = strconv.ParseInt(dataArray[18], 10, 64)
		data.SellOnePri, _ = strconv.ParseFloat(dataArray[21], 64)
		data.SellOne, _ = strconv.ParseInt(dataArray[20], 10, 64)
		data.SellTwoPri, _ = strconv.ParseFloat(dataArray[23], 64)
		data.SellTwo, _ = strconv.ParseInt(dataArray[22], 10, 64)
		data.SellThreePri, _ = strconv.ParseFloat(dataArray[24], 64)
		data.SellThree, _ = strconv.ParseInt(dataArray[24], 10, 64)
		data.SellFourPri, _ = strconv.ParseFloat(dataArray[27], 64)
		data.SellFour, _ = strconv.ParseInt(dataArray[26], 10, 64)
		data.SellFivePri, _ = strconv.ParseFloat(dataArray[29], 64)
		data.SellFive, _ = strconv.ParseInt(dataArray[28], 10, 64)
		data.Time = dataArray[30] + "-" + dataArray[31]
		data.Change, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", data.NowPri-data.YestClosePri), 64)
		data.ChangePer, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", 100.0*(data.NowPri-data.YestClosePri)/data.YestClosePri), 64)
		data.TodayMax, _ = strconv.ParseFloat(dataArray[4], 64)
		data.TodayMin, _ = strconv.ParseFloat(dataArray[5], 64)
		data.TradeCount, _ = strconv.ParseInt(dataArray[8], 10, 64)
		data.TradeAmont, _ = strconv.ParseInt(dataArray[9], 10, 64)
		//data.ChangeRate, _ = strconv.ParseFloat(dataArray[38], 64)
		//data.PERatio, _ = strconv.ParseFloat(dataArray[39], 64)
		//data.MaxMinChange, _ = strconv.ParseFloat(dataArray[43], 64)
		//data.MarketAmont, _ = strconv.ParseFloat(dataArray[44], 64)
		//data.TotalAmont, _ = strconv.ParseFloat(dataArray[45], 64)
		//data.PBRatio, _ = strconv.ParseFloat(dataArray[46], 64)
		//data.HighPri, _ = strconv.ParseFloat(dataArray[47], 64)
		//data.LowPri, _ = strconv.ParseFloat(dataArray[48], 64)

		dataList = append(dataList, data)
	}
	return dataList
}

func GetPKSina(code string) *PKData {
	resp, err := http.DefaultClient.Get(fmt.Sprintf(URL_SINA_PK, code))
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
func GetFundFlowSina(code string) *FundFlow {
	resp, err := http.DefaultClient.Get(fmt.Sprintf(URL_SINA_FUND_FLOW, code))
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

func GetInfoSina(code string) *StockInfo {
	resp, err := http.DefaultClient.Get(fmt.Sprintf(URL_SINA_INFO, code))
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

func GetDailySina(code string, year int) []*HistoryData {
	resp, err := http.DefaultClient.Get(fmt.Sprintf(URL_SINA_DAILY, year, code))
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

func GetWeeklySina(code string) []*HistoryData {
	resp, err := http.DefaultClient.Get(fmt.Sprintf(URL_SINA_WEEKLY, code))
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
func (quotes *Quotes) FetchSina() (self *Quotes) {
	self = quotes // <-- This ensures we return correct quotes after recover() from panic().
	if quotes.isReady() {
		//defer func() {
		//	if err := recover(); err != nil {
		//		quotes.errors = fmt.Sprintf("\n\n\n\nError fetching stock quotes...\n%s", err)
		//	}
		//}()

		codes := util.StockWithPrefix(quotes.profile.Tickers)
		results := GetRealtimeSina(strings.Join(codes, ","))

		quotes.stocks = make([]Stock, len(results))
		for i, raw := range results {
			realTime := raw

			pyStr := pinyin.Convert(realTime.Name)
			reg, err := regexp.Compile("[a-z]+")
			if err != nil {
				log.Fatal(err)
			}
			pyCapStr := reg.ReplaceAllString(pyStr, "")
			if false {
				quotes.stocks[i].Ticker = realTime.Name
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

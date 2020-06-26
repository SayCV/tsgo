package TerminalStocks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"
	"unicode"

	//"regexp"
	"strconv"
	"strings"

	//log "github.com/sirupsen/logrus"
	"github.com/axgle/mahonia"
	"github.com/axgle/pinyin"

	//"github.com/axgle/pinyin"
	util "github.com/saycv/tsgo/pkg/utils"
)

const (
	// URL_EASTMONEY_REAL_TIME = "http://21.push2.eastmoney.com/api/qt/clist/get?cb=jQuery112407201580659678162_1581950914193&pn=1&pz=20000&po=1&np=1&ut=bd1d9ddb04089700cf9c27f6f7426281&fltt=2&invt=2&fid=f3&fs=m:0+t:6,m:0+t:13,m:0+t:80,m:1+t:2,m:1+t:23&fields=f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f23,f24,f25,f22,f11,f62,f128,f136,f115,f152"
	URL_EASTMONEY_REAL_TIME = "http://push2.eastmoney.com/api/qt/ulist.np/get?ut=bd1d9ddb04089700cf9c27f6f7426281&fields=f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f23,f24,f25,f22,f11,f62,f128,f136,f115,f152&fltt=2&secids=%s&cb=jQuery112403947390609559913_1593128751691&_=1593128751711"
	URL_EASTMONEY_FUND_FLOW = "http://qt.gtimg.cn/q=ff_%s"
	URL_EASTMONEY_PK        = "http://qt.gtimg.cn/q=s_pk%s"
	URL_EASTMONEY_INFO      = "http://qt.gtimg.cn/q=s_%s"
	URL_EASTMONEY_DAILY     = "http://push2his.eastmoney.com/api/qt/stock/kline/get?fields1=f1,f2,f3,f4,f5,f6&fields2=f51,f52,f53,f54,f55,f56,f57,f58,f61&klt=101&fqt=1&secid=%s&beg=%s&end=%s&_=1591683995756"
	URL_EASTMONEY_WEEKLY    = "http://data.gtimg.cn/flashdata/hushen/weekly/%s.js"
)

func GetRealtimeEastmoney(code string) []*RealTimeData {
	//fmt.Println(fmt.Sprintf(URL_EASTMONEY_REAL_TIME, code))
	resp, err := http.DefaultClient.Get(fmt.Sprintf(URL_EASTMONEY_REAL_TIME, code))
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

func GetPKEastmoney(code string) *PKData {
	resp, err := http.DefaultClient.Get(fmt.Sprintf(URL_EASTMONEY_PK, code))
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
func GetFundFlowEastmoney(code string) *FundFlow {
	resp, err := http.DefaultClient.Get(fmt.Sprintf(URL_EASTMONEY_FUND_FLOW, code))
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

func GetInfoEastmoney(code string) *StockInfo {
	resp, err := http.DefaultClient.Get(fmt.Sprintf(URL_EASTMONEY_INFO, code))
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

func GetDailyEastmoney(code string) []*HistoryData {
	now := time.Now()
	year, month, day := now.Date()
	yearstart := fmt.Sprintf("%04d%02d%02d", year-2, month, day)
	yearend := fmt.Sprintf("%04d%02d%02d", year+1, month, day)
	url := fmt.Sprintf(URL_EASTMONEY_DAILY, code, yearstart, yearend)
	//fmt.Println(url)
	resp, err := http.DefaultClient.Get(url)
	checkErr(err)
	if resp == nil || resp.StatusCode != http.StatusOK {
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	res := string(body)
	d := map[string]interface{}{}
	//dataArray := strings.Split(res, "\\n\\")
	list := []*HistoryData{}
	err = json.Unmarshal([]byte(res), &d)
	if err != nil {
		return list
	}
	results := d["data"].(map[string]interface{})["klines"].([]interface{})
	dataArray := results

	for index, str := range dataArray {
		if index == -1 || index == -len(results)-1 {

		} else {
			data := strings.Split(str.(string), ",")
			entity := new(HistoryData)
			entity.Date = strings.Replace(data[0], "\n", "", -1)
			entity.Open, _ = strconv.ParseFloat(data[1], 64)
			entity.Close, _ = strconv.ParseFloat(data[2], 64)
			entity.Max, _ = strconv.ParseFloat(data[3], 64)
			entity.Min, _ = strconv.ParseFloat(data[4], 64)
			entity.Trade, _ = strconv.ParseFloat(data[5], 64)
			list = append(list, entity)
		}
		//fmt.Println("index = ", index, ", str = ", str)
	}

	return list
}

func GetWeeklyEastmoney(code string) []*HistoryData {
	resp, err := http.DefaultClient.Get(fmt.Sprintf(URL_EASTMONEY_WEEKLY, code))
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
func (quotes *Quotes) FetchEastmoney() (self *Quotes) {
	self = quotes // <-- This ensures we return correct quotes after recover() from panic().
	if quotes.isReady() {
		defer func() {
			if err := recover(); err != nil {
				quotes.errors = fmt.Sprintf("\n\n\n\nError fetching stock quotes...\n%s", err)
			}
		}()

		codes := util.StockWithPrefixEastmoney(quotes.profile.Tickers)
		url := fmt.Sprintf(URL_EASTMONEY_REAL_TIME, strings.Join(codes, ","))
		//fmt.Println(url)
		//response, err := http.Get(url)
		request, err := http.NewRequest("GET", url, nil)
		if err != nil {
			panic(err)
		}
		request.Header.Add("Cookie", "name=anny")
		request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.113 Safari/537.36")
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			panic(err)
		}

		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}

		quotes.parseEastmoney(body, codes)
	}

	return quotes
}

// this will parse the json objects
func (quotes *Quotes) parseEastmoney(body []byte, codes []string) (*Quotes, error) {
	// response -> quoteResponse -> result|error (array) -> map[string]interface{}
	// Stocks has non-int things
	// d := map[string]map[string][]Stock{}
	// some of these are numbers vs strings
	// d := map[string]map[string][]map[string]string{}
	d := map[string]interface{}{}
	reusltString := strings.TrimLeft(string(body), "jQuery112403947390609559913_1593128751691(")
	reusltString = strings.TrimRight(reusltString, ");")
	err := json.Unmarshal([]byte(reusltString), &d)
	if err != nil {
		return nil, err
	}
	results := d["data"].(map[string]interface{})["diff"].([]interface{})
	//fmt.Println(len(results))
	//fmt.Println(results)

	quotes.stocks = make([]Stock, len(results))
	for i, raw := range results {
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

		name := result["f14"]
		pyStr := pinyin.Convert(name)
		reg, err := regexp.Compile("[a-z]+")
		if err != nil {
			log.Fatal(err)
		}
		pyCapStr := reg.ReplaceAllString(pyStr, "")
		if true {
			if true {
				//quotes.stocks[i].Ticker = strings.Join(strings.Split(name, ""), " ")
				newstring := ""
				for _, char := range name {
					if unicode.Is(unicode.Scripts["Han"], char) {
						newstring = newstring + string(char) + " "
					} else {
						newstring = newstring + string(char)
					}
				}
				quotes.stocks[i].Ticker = newstring
			} else {
				quotes.stocks[i].Ticker = name
			}
		} else {
			quotes.stocks[i].Ticker = pyCapStr
		}

		yestclose, _ := strconv.ParseFloat(result["f18"], 64)
		now, _ := strconv.ParseFloat(result["f2"], 64)

		quotes.stocks[i].LastTrade = result["f2"]
		quotes.stocks[i].Change = fmt.Sprintf("%.2f", now-yestclose)
		quotes.stocks[i].ChangePct = fmt.Sprintf("%.2f", 100.0*(now-yestclose)/yestclose)
		quotes.stocks[i].Open = result["f17"]
		quotes.stocks[i].Low = result["f16"]
		quotes.stocks[i].High = result["f15"]
		quotes.stocks[i].Low52 = ""
		quotes.stocks[i].High52 = ""
		quotes.stocks[i].Volume = result["f5"]
		quotes.stocks[i].AvgVolume = ""
		quotes.stocks[i].PeRatio = result["f115"]
		// TODO calculate rt
		quotes.stocks[i].PeRatioX = result["f9"]
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

	for i, code := range codes {
		weekly := GetDailyEastmoney(code)
		low52, high52 := func(list []*HistoryData) (min float64, max float64) {
			listnbr := len(list)
			startIndex := 0
			if len(weekly) > 52*5 {
				startIndex = listnbr - 52*5
			} else if len(weekly) < 1 {
				return 0, 0
			} else {
				startIndex = 0
			}
			min = list[startIndex].Min
			max = list[startIndex].Max

			for _, entity := range list[startIndex:] {
				valueMin := entity.Min
				valueMax := entity.Max
				if valueMin < min {
					min = valueMin
				}
				if valueMax > max {
					max = valueMax
				}
			}
			return min, max
		}(weekly)

		quotes.stocks[i].Low52 = strconv.FormatFloat(low52, 'f', -1, 64)
		quotes.stocks[i].High52 = strconv.FormatFloat(high52, 'f', -1, 64)
	}

	return quotes, nil
}

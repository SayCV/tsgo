package TerminalStocks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/axgle/pinyin"
	util "github.com/saycv/tsgo/pkg/utils"
)

const (
	URL_EASTMONEY_LIMITUP = "http://93.push2.eastmoney.com/api/qt/clist/get?cb=jQuery112406773975227501083_1593238765028&pn=1&pz=%d&po=1&np=1&ut=bd1d9ddb04089700cf9c27f6f7426281&fltt=2&invt=2&fid=f3&fs=m:0+t:6,m:0+t:13,m:0+t:80,m:1+t:2,m:1+t:23&fields=f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f23,f24,f25,f22,f11,f62,f128,f136,f115,f152&_=1593238765029"
)

func GetLimitupEastmoney() ([]Stock, error) {
	queryNbr := 50
	url := fmt.Sprintf(URL_EASTMONEY_LIMITUP, queryNbr)
	//fmt.Println(url)
	resp, err := http.DefaultClient.Get(url)
	checkErr(err)
	if resp == nil || resp.StatusCode != http.StatusOK {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	res := body //mahonia.NewDecoder("gbk").ConvertString(string(body))
	d := map[string]interface{}{}
	reusltString := strings.TrimLeft(string(res), "jQuery112403947390609559913_1593128751691(")
	reusltString = strings.TrimRight(reusltString, ");")
	err = json.Unmarshal([]byte(reusltString), &d)
	if err != nil {
		return nil, err
	}
	results := d["data"].(map[string]interface{})["diff"].([]interface{})
	//fmt.Println(len(results))
	//fmt.Println(results)

	stocks := make([]Stock, len(results))
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
				//stocks[i].Ticker = strings.Join(strings.Split(name, ""), " ")
				newstring := ""
				for _, char := range name {
					if unicode.Is(unicode.Scripts["Han"], char) {
						newstring = newstring + string(char) + " "
					} else {
						newstring = newstring + string(char)
					}
				}
				stocks[i].Ticker = newstring
			} else {
				stocks[i].Ticker = name
			}
		} else {
			stocks[i].Ticker = pyCapStr
		}
		stocks[i].Gid = result["f12"]

		yestclose, _ := strconv.ParseFloat(result["f18"], 64)
		now, _ := strconv.ParseFloat(result["f2"], 64)

		stocks[i].LastTrade = result["f2"]
		stocks[i].Change = fmt.Sprintf("%.2f", now-yestclose)
		stocks[i].ChangePct = fmt.Sprintf("%.2f", 100.0*(now-yestclose)/yestclose)
		stocks[i].Open = result["f17"]
		stocks[i].Low = result["f16"]
		stocks[i].High = result["f15"]
		stocks[i].Low52 = ""
		stocks[i].High52 = ""
		stocks[i].Volume = result["f5"]
		stocks[i].AvgVolume = ""
		stocks[i].PeRatio = result["f115"]
		// TODO calculate rt
		stocks[i].PeRatioX = result["f9"]
		stocks[i].Dividend = ""
		stocks[i].Yield = ""
		stocks[i].MarketCap = ""
		// TODO calculate rt?
		stocks[i].MarketCapX = ""

		/*
			fmt.Println(i)
			fmt.Println("-------------------")
			for k, v := range result {
				fmt.Println(k, v)
			}
			fmt.Println("-------------------")
		*/
		adv, err := strconv.ParseFloat(stocks[i].Change, 64)
		if err == nil {
			stocks[i].Advancing = adv >= 0.0
		}
		i++
	}

	for i, stock := range stocks {
		code := util.StockWithPrefixEastmoney([]string{stock.Gid})
		weekly := GetDailyEastmoney(code[0])
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

		stocks[i].Low52 = strconv.FormatFloat(low52, 'f', -1, 64)
		stocks[i].High52 = strconv.FormatFloat(high52, 'f', -1, 64)
	}

	return stocks, nil
}

package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// StockWithPrefix autoset stockcode prefix
// 600***, 601*** 上证A股;
// 000***, 002*** 深圳A股;
// 400*** 三板市场股票
// 300*** 创业板
func StockWithPrefix(code []string) []string {
	var results []string
	for _, one := range code {
		new := ""
		switch {
		case strings.HasPrefix(one, "000"):
			new = fmt.Sprintf("sz%s", one)
		case strings.HasPrefix(one, "002"):
			new = fmt.Sprintf("sz%s", one)
		case strings.HasPrefix(one, "300"):
			new = fmt.Sprintf("sz%s", one)
		case strings.HasPrefix(one, "600"):
			new = fmt.Sprintf("sh%s", one)
		case strings.HasPrefix(one, "601"):
			new = fmt.Sprintf("sh%s", one)
		case strings.HasPrefix(one, "603"):
			new = fmt.Sprintf("sh%s", one)
		case strings.HasPrefix(one, "605"):
			new = fmt.Sprintf("sh%s", one)
		case strings.HasPrefix(one, "688"):
			new = fmt.Sprintf("sh%s", one)
		default:
			new = one
		}
		//fmt.Println(new)
		results = append(results, new)
	}
	return results
}

func StockWithPrefixNetease(code []string) []string {
	var results []string
	for _, one := range code {
		new := ""
		switch {
		case strings.HasPrefix(one, "000"):
			new = fmt.Sprintf("1%s", one)
		case strings.HasPrefix(one, "002"):
			new = fmt.Sprintf("1%s", one)
		case strings.HasPrefix(one, "300"):
			new = fmt.Sprintf("1%s", one)
		case strings.HasPrefix(one, "600"):
			new = fmt.Sprintf("0%s", one)
		case strings.HasPrefix(one, "601"):
			new = fmt.Sprintf("0%s", one)
		case strings.HasPrefix(one, "603"):
			new = fmt.Sprintf("0%s", one)
		case strings.HasPrefix(one, "605"):
			new = fmt.Sprintf("0%s", one)
		case strings.HasPrefix(one, "688"):
			new = fmt.Sprintf("0%s", one)
		default:
			new = one
		}
		//fmt.Println(new)
		results = append(results, new)
	}
	return results
}

func StockWithPrefixEastmoney(code []string) []string {
	var results []string
	for _, one := range code {
		new := ""
		switch {
		case strings.HasPrefix(one, "000"):
			new = fmt.Sprintf("0.%s", one)
		case strings.HasPrefix(one, "002"):
			new = fmt.Sprintf("0.%s", one)
		case strings.HasPrefix(one, "300"):
			new = fmt.Sprintf("0.%s", one)
		case strings.HasPrefix(one, "600"):
			new = fmt.Sprintf("1.%s", one)
		case strings.HasPrefix(one, "601"):
			new = fmt.Sprintf("1.%s", one)
		case strings.HasPrefix(one, "603"):
			new = fmt.Sprintf("1.%s", one)
		case strings.HasPrefix(one, "605"):
			new = fmt.Sprintf("1.%s", one)
		case strings.HasPrefix(one, "688"):
			new = fmt.Sprintf("1.%s", one)
		default:
			new = one
		}
		//fmt.Println(new)
		results = append(results, new)
	}
	return results
}

func JsonDecodeS(jstr string, obj interface{}) {
	json.Unmarshal([]byte(jstr), obj)
}

func JsonDecodeB(value []byte, obj interface{}) {
	json.Unmarshal(value, obj)
}

func JsonEncodeS(obj interface{}) string {
	res, _ := json.Marshal(obj)
	return string(res)
}

func JsonEncodeB(obj interface{}) []byte {
	res, _ := json.Marshal(obj)
	return res
}

func GetCurrentTimeStamp() string {
	return time.Now().Format("2020-06-21 06:21:00")
}

/*
 * @query a single date: string '20200627';
 * @api return day_type: 0 workday 1 weekend 2 holiday -1 err
 * @function return day_type: 1 workday 0 weekend&holiday
 */
func get_day_type(query_date string) int {
	URL_EASTMONEY_DAILY := "http://push2his.eastmoney.com/api/qt/stock/kline/get?fields1=f1,f2,f3,f4,f5,f6&fields2=f51,f52,f53,f54,f55,f56,f57,f58,f61&klt=101&fqt=1&secid=%s&beg=%s&end=%s&_=1591683995756"
	yearstart := query_date
	yearend := query_date
	url := fmt.Sprintf(URL_EASTMONEY_DAILY, "1.000001", yearstart, yearend)
	//fmt.Println(url)
	resp, err := http.DefaultClient.Get(url)
	if err != nil || resp == nil || resp.StatusCode != http.StatusOK {
		return -1
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	res := string(body)
	d := map[string]interface{}{}
	err = json.Unmarshal([]byte(res), &d)
	if err != nil {
		return -1
	}
	results := d["data"].(map[string]interface{})["klines"].([]interface{})

	nbr := len(results)
	if nbr < 1 {
		return 2
	}
	return 0
}

func IsTradeDay(query_date string) bool {
	if get_day_type(query_date) == 0 {
		return true
	}
	return false
}

func GetLatestTradeDay(query_date string) string {
	//now := time.Now()
	now, _ := time.ParseInLocation("20060102 15:04:05", query_date+" 15:00:00", time.Local)
	loop := 5
	for {
		if loop < 0 {
			//fmt.Println("20060102")
			return "20060102"
		}
		year, month, day := now.Date()
		tstr := fmt.Sprintf("%04d%02d%02d", year, month, day)
		//fmt.Println(tstr)
		if IsTradeDay(tstr) {
			return tstr
		}
		now = now.AddDate(0, 0, -1)
		loop--
	}
}

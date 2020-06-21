package util

import (
	"encoding/json"
	"fmt"
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

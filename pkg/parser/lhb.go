package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/axgle/mahonia"
	TerminalStocks "github.com/saycv/tsgo/pkg/terminalstocks"
	util "github.com/saycv/tsgo/pkg/utils"
)

const (
	URL_EASTMONEY_LHB = "http://data.eastmoney.com/DataCenter_V3/stock2016/TradeDetail/pagesize=200,page=1,sortRule=-1,sortType=,startDate=%s,endDate=%s,gpfw=0,js=var data_tab_1.html?rt=26553644"
)

func checkErr(err error) {}

func GetLhbEastmoney(date string) []TerminalStocks.LhbData {
	//now := time.Now()
	//year, month, day := now.Date()
	//yearstart := fmt.Sprintf("%04d-%02d-%02d", year, month, day)
	//yearend := fmt.Sprintf("%04d-%02d-%02d", year, month, day)
	yearstart := util.GetLatestTradeDay(date)
	yearend := yearstart
	url := fmt.Sprintf(URL_EASTMONEY_LHB, yearstart, yearend)
	fmt.Println(url)
	resp, err := http.DefaultClient.Get(url)
	checkErr(err)
	if resp == nil || resp.StatusCode != http.StatusOK {
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	res := mahonia.NewDecoder("gbk").ConvertString(string(body))
	d := map[string]interface{}{}
	reusltString := strings.TrimLeft(string(res), "var data_tab_1=")
	reusltString = strings.TrimRight(reusltString, ");")
	err = json.Unmarshal([]byte(reusltString), &d)
	list := []TerminalStocks.LhbData{}
	if err != nil {
		return list
	}
	results := d["data"].([]interface{})
	dataArray := results

	for index, str := range dataArray {
		if index == -1 || index == -len(results)-1 {

		} else {
			data := map[string]string{}
			for k, v := range str.(map[string]interface{}) {
				switch v.(type) {
				case string:
					data[k] = v.(string)
				case float64:
					data[k] = float2Str(v.(float64))
				default:
					data[k] = fmt.Sprintf("%v", v)
				}

			}
			entity := new(TerminalStocks.LhbData)
			entity.SCode = data["SCode"]
			entity.SName = data["SName"]
			entity.Price, _ = strconv.ParseFloat(data["ClosePrice"], 64)
			list = append(list, *entity)
		}
		//fmt.Println("index = ", index, ", str = ", str)
	}

	return list
}

func float2Str(v float64) string {
	unit := ""
	switch {
	case v > 1.0e12:
		v = v / 1.0e12
		unit = "T"
	case v > 1.0e9:
		v = v / 1.0e9
		unit = "B"
	case v > 1.0e6:
		v = v / 1.0e6
		unit = "M"
	case v > 1.0e5:
		v = v / 1.0e3
		unit = "K"
	default:
		unit = ""
	}
	// parse
	return fmt.Sprintf("%0.3f%s", v, unit)
}

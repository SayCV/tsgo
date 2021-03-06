package TerminalStocks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/axgle/mahonia"
	util "github.com/saycv/tsgo/pkg/utils"
)

const (
	URL_EASTMONEY_LHB = "http://data.eastmoney.com/DataCenter_V3/stock2016/TradeDetail/pagesize=200,page=1,sortRule=-1,sortType=,startDate=%s,endDate=%s,gpfw=0,js=var data_tab_1.html?rt=26553644"
)

func GetLhbEastmoney(date string) []LhbData {
	//now := time.Now()
	//year, month, day := now.Date()
	//yearstart := fmt.Sprintf("%04d-%02d-%02d", year, month, day)
	//yearend := fmt.Sprintf("%04d-%02d-%02d", year, month, day)
	yearstart := util.GetLatestTradeDay(date)
	yearend := yearstart
	url := fmt.Sprintf(URL_EASTMONEY_LHB, yearstart, yearend)
	//fmt.Println(url)
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
	list := []LhbData{}
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
			entity := new(LhbData)
			entity.SCode = data["SCode"]
			entity.SName = data["SName"]
			entity.Price, _ = strconv.ParseFloat(data["ClosePrice"], 64)
			entity.ChangePer, _ = strconv.ParseFloat(data["Chgradio"], 64)
			entity.ChangeRate, _ = strconv.ParseFloat(data["Dchratio"], 64)
			entity.LhbInMoney, _ = strconv.ParseFloat(data["JmMoney"], 64)
			entity.TradeAmont, _ = strconv.ParseFloat(data["Turnover"], 64)
			//entity.LhbInMoney, _ = strconv.ParseFloat(data["JmMoney"], 64)
			entity.LhbCauses = data["Ctypedes"]
			entity.LhbSellMoney, _ = strconv.ParseFloat(data["Smoney"], 64)
			entity.LhbBuyMoney, _ = strconv.ParseFloat(data["Bmoney"], 64)
			entity.LhbTradeAmont, _ = strconv.ParseFloat(data["ZeMoney"], 64)
			entity.LhbNotes = data["JD"]

			list = append(list, *entity)
		}
		//fmt.Println("index = ", index, ", str = ", str)
	}

	return list
}

package stocks

import (
	"testing"

	"github.com/ShawnRong/tushare-go"
	//"github.com/stretchr/testify/require"
)

const (
	TUSHARE_TOKEN = "4c010225d485a8db581030b9c04f634d14d3b2bd92fa2e3546a77bbe"
)

func TestQuotes(t *testing.T) {

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

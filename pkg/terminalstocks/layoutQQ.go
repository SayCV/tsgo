// Copyright (c) 2013-2016 by Michael Dvorkin. All Rights Reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package TerminalStocks

import (
	"text/template"
)

//-----------------------------------------------------------------------------
func buildQQMarketTemplate() *template.Template {
	markup := `<yellow>Dow</> {{.Dow.change}} ({{.Dow.percent}}) at {{.Dow.latest}} <yellow>NASDAQ</> {{.Nasdaq.change}} ({{.Nasdaq.percent}}) at {{.Nasdaq.latest}}
<yellow>上 指</> {{.Szzs.change}} ({{.Szzs.percent}}%) at {{.Szzs.latest}} <yellow>深 成</> {{.Szcz.change}} ({{.Szcz.percent}}%) at {{.Szcz.latest}}
<yellow>沪 深</> {{.Hs300.change}} ({{.Hs300.percent}}%) at {{.Hs300.latest}} <yellow>创 指</> {{.Cybz.change}} ({{.Cybz.percent}}%) at {{.Cybz.latest}}
<yellow>Oil</> ${{.Oil.latest}} ({{.Oil.change}}%) <yellow>Gold</> ${{.Gold.latest}} ({{.Gold.change}}%)`

	return template.Must(template.New(`market`).Parse(markup))
}

//-----------------------------------------------------------------------------
func buildQQQuotesTemplate() *template.Template {
	markup := `<right><white>{{.Now}}</></right>



{{.Header}}
{{range.Stocks}}{{if .Advancing}}<cyan>{{else}}<green>{{end}}{{.Ticker}}{{.LastTrade}}{{.Change}}{{.ChangePct}}{{.Open}}{{.Low}}{{.High}}{{.Low52}}{{.High52}}{{.Volume}}{{.AvgVolume}}{{.PeRatio}}{{.Dividend}}{{.Yield}}{{.MarketCap}}</>
{{end}}`

	return template.Must(template.New(`quotes`).Parse(markup))
}

//-----------------------------------------------------------------------------
func buildLhbQuotesTemplate() *template.Template {
	markup := `<right><white>{{.Now}}</></right>



{{.Header}}
{{range.Stocks}}{{if .Advancing}}<cyan>{{else}}<green>{{end}}{{.Ticker}}{{.LastTrade}}{{.Change}}{{.ChangePct}}{{.Open}}{{.Low}}{{.High}}{{.Low52}}{{.High52}}{{.Volume}}{{.Gid}}{{.Yield}}{{.MarketCap}}</>
{{end}}`

	return template.Must(template.New(`quotes`).Parse(markup))
}

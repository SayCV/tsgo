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
<yellow>SZZS</> {{.Szzs.change}} ({{.Szzs.percent}}) at {{.Szzs.latest}} <yellow>SZCZ</> {{.Szcz.change}} ({{.Szcz.percent}}) at {{.Szcz.latest}}
<yellow>Oil</> ${{.Oil.latest}} ({{.Oil.change}}%) <yellow>Gold</> ${{.Gold.latest}} ({{.Gold.change}}%)`

	return template.Must(template.New(`market`).Parse(markup))
}

//-----------------------------------------------------------------------------
func buildQQQuotesTemplate() *template.Template {
	markup := `<right><white>{{.Now}}</></right>



{{.Header}}
{{range.Stocks}}{{if .Advancing}}<green>{{else}}<red>{{end}}{{.Ticker}}{{.LastTrade}}{{.Change}}{{.ChangePct}}{{.Open}}{{.Low}}{{.High}}{{.Low52}}{{.High52}}{{.Volume}}{{.AvgVolume}}{{.PeRatio}}{{.Dividend}}{{.Yield}}{{.MarketCap}}</>
{{end}}`

	return template.Must(template.New(`quotes`).Parse(markup))
}

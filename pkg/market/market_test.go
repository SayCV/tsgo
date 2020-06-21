package market

import (
	"testing"

	TerminalStocks "github.com/saycv/tsgo/pkg/terminalstocks"
	"github.com/stretchr/testify/require"
)

func TestQuotes(t *testing.T) {
	market := TerminalStocks.NewMarket(TerminalStocks.API_VENDOR_YAHOO)
	profile := TerminalStocks.NewProfile()

	profile.Tickers = []string{"GOOG", "BA"}

	quotes := TerminalStocks.NewQuotes(market, profile)
	require.NotNil(t, quotes)

	t.Log(market.Fetch())
	t.Log(quotes.Fetch())
}

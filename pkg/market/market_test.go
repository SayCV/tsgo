package market

import (
	"testing"

	"github.com/brandleesee/TerminalStocks"
	"github.com/stretchr/testify/require"
)

func TestQuotes(t *testing.T) {
	market := TerminalStocks.NewMarket()
	profile := TerminalStocks.NewProfile()

	profile.Tickers = []string{"GOOG", "BA"}

	quotes := TerminalStocks.NewQuotes(market, profile)
	require.NotNil(t, quotes)

	t.Log(market.Fetch())
	t.Log(quotes.Fetch())
}

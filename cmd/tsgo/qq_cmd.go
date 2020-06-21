package main

import (
	"fmt"
	"time"

	//"github.com/saycv/tsgo"
	//"github.com/saycv/tsgo/pkg/configuration"
	"github.com/nsf/termbox-go"
	logsupport "github.com/saycv/tsgo/pkg/log"
	TerminalStocks "github.com/saycv/tsgo/pkg/terminalstocks"
	util "github.com/saycv/tsgo/pkg/utils"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// NewQQAPICmd returns the root command
func NewQQAPICmd() *cobra.Command {

	var logLevel string

	rootCmd := &cobra.Command{
		Use:   "qq",
		Short: `Fetch finance from QQ API`,
		Args:  cobra.ArbitraryArgs,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			lvl, err := log.ParseLevel(logLevel)
			if err != nil {
				fmt.Fprintf(cmd.OutOrStderr(), "unable to parse log level '%v'", logLevel)
				return err
			}
			logsupport.Setup()
			log.SetLevel(lvl)
			log.SetOutput(cmd.OutOrStdout())
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			//if len(args) <= 1 {
			//	return helpCommand.RunE(cmd, args)
			//}
			//attrs := parseAttributes(attributes)

			//for _, sourcePath := range args {
			//}
			screen := TerminalStocks.NewScreen(TerminalStocks.API_VENDOR_QQ)
			defer screen.Close()

			profile := TerminalStocks.NewProfile()
			code := []string{"600519", "601318", "601066", "002958", "000878", "600121", "603121"}
			profile.Tickers = util.StockWithPrefix(code)
			profile.SortColumn = 3

			mainLoopQQ(screen, profile)
			return nil
		},
	}
	rootCmd.SilenceUsage = true
	flags := rootCmd.Flags()

	flags.StringVarP(&logLevel, "log", "l", "debug", "log level to set [debug|info|warning|error|fatal|panic]")

	// rootCmd.MarkFlagRequired("logLevel")

	return rootCmd
}

//-----------------------------------------------------------------------------
func mainLoopQQ(screen *TerminalStocks.Screen, profile *TerminalStocks.Profile) {
	var lineEditor *TerminalStocks.LineEditor
	var columnEditor *TerminalStocks.ColumnEditor

	keyboardQueue := make(chan termbox.Event)
	timestampQueue := time.NewTicker(1 * time.Second)
	quotesQueue := time.NewTicker(5 * time.Second)
	marketQueue := time.NewTicker(12 * time.Second)
	showingHelp := false
	paused := false

	go func() {
		for {
			keyboardQueue <- termbox.PollEvent()
		}
	}()

	market := TerminalStocks.NewMarket(TerminalStocks.API_VENDOR_QQ)
	quotes := TerminalStocks.NewQuotes(market, profile, TerminalStocks.API_VENDOR_QQ)
	screen.Draw(market, quotes)

loop:
	for {
		select {
		case event := <-keyboardQueue:
			switch event.Type {
			case termbox.EventKey:
				if lineEditor == nil && columnEditor == nil && !showingHelp {
					if event.Key == termbox.KeyEsc || event.Ch == 'q' || event.Ch == 'Q' {
						break loop
					} else if event.Ch == '+' || event.Ch == '-' {
						lineEditor = TerminalStocks.NewLineEditor(screen, quotes)
						lineEditor.Prompt(event.Ch)
					} else if event.Ch == 'o' || event.Ch == 'O' {
						columnEditor = TerminalStocks.NewColumnEditor(screen, quotes)
					} else if event.Ch == 'g' || event.Ch == 'G' {
						if profile.Regroup() == nil {
							screen.Draw(quotes)
						}
					} else if event.Ch == 'p' || event.Ch == 'P' {
						paused = !paused
						screen.Pause(paused).Draw(time.Now())
					} else if event.Ch == '?' || event.Ch == 'h' || event.Ch == 'H' {
						showingHelp = true
						screen.Clear().Draw(helpYahoo)
					}
				} else if lineEditor != nil {
					if done := lineEditor.Handle(event); done {
						lineEditor = nil
					}
				} else if columnEditor != nil {
					if done := columnEditor.Handle(event); done {
						columnEditor = nil
					}
				} else if showingHelp {
					showingHelp = false
					screen.Clear().Draw(market, quotes)
				}
			case termbox.EventResize:
				screen.Resize()
				if !showingHelp {
					screen.Draw(market, quotes)
				} else {
					screen.Draw(helpYahoo)
				}
			}

		case <-timestampQueue.C:
			if !showingHelp && !paused {
				screen.Draw(time.Now())
			}

		case <-quotesQueue.C:
			if !showingHelp && !paused {
				screen.Draw(quotes)
			}

		case <-marketQueue.C:
			if !showingHelp && !paused {
				screen.Draw(market)
			}
		}
	}
}

package main

import (
	"fmt"

	//"github.com/saycv/tsgo"
	//"github.com/saycv/tsgo/pkg/configuration"

	logsupport "github.com/saycv/tsgo/pkg/log"
	TerminalStocks "github.com/saycv/tsgo/pkg/terminalstocks"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// NewSinaAPICmd returns the root command
func NewSinaAPICmd() *cobra.Command {

	var logLevel string

	rootCmd := &cobra.Command{
		Use:   "sina",
		Short: `Fetch finance from Sina API`,
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
			screen := TerminalStocks.NewScreen(TerminalStocks.API_VENDOR_SINA)
			defer screen.Close()

			profile := TerminalStocks.NewProfile(TerminalStocks.API_VENDOR_SINA)

			mainLoop(screen, profile)
			return nil
		},
	}
	rootCmd.SilenceUsage = true
	flags := rootCmd.Flags()

	flags.StringVarP(&logLevel, "log", "l", "debug", "log level to set [debug|info|warning|error|fatal|panic]")

	// rootCmd.MarkFlagRequired("logLevel")

	return rootCmd
}

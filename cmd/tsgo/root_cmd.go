package main

import (
	"fmt"
	"strings"

	//"github.com/saycv/tsgo"
	//"github.com/saycv/tsgo/pkg/configuration"

	logsupport "github.com/saycv/tsgo/pkg/log"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// NewRootCmd returns the root command
func NewRootCmd() *cobra.Command {

	var logLevel string

	rootCmd := &cobra.Command{
		Use:   "tsgo [-logLevel]",
		Short: `tsgo is a command-line utility that displays stocks`,
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

			return nil
		},
	}
	rootCmd.SilenceUsage = true
	flags := rootCmd.Flags()

	flags.StringVarP(&logLevel, "log", "l", "debug", "log level to set [debug|info|warning|error|fatal|panic]")

	// rootCmd.MarkFlagRequired("logLevel")

	return rootCmd
}

// converts the `name`, `!name` and `name=value` into a map
func parseAttributes(attributes []string) map[string]string {
	result := make(map[string]string, len(attributes))
	for _, attr := range attributes {
		data := strings.Split(attr, "=")
		if len(data) > 1 {
			result[data[0]] = data[1]
		} else {
			result[data[0]] = ""
		}
	}
	return result
}

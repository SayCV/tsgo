package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := NewRootCmd()
	versionCmd := NewVersionCmd()
	yahooAPICmd := NewYahooAPICmd()
	qqAPICmd := NewQQAPICmd()
	sinaAPICmd := NewSinaAPICmd()
	neteaseAPICmd := NewNeteaseAPICmd()
	eastmoneyAPICmd := NewEastmoneyAPICmd()
	eastmoneyLimitupAPICmd := NewEastmoneyLimitupAPICmd()
	eastmoneyLhbAPICmd := NewEastmoneyLhbAPICmd()
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(yahooAPICmd)
	rootCmd.AddCommand(qqAPICmd)
	rootCmd.AddCommand(sinaAPICmd)
	rootCmd.AddCommand(neteaseAPICmd)
	rootCmd.AddCommand(eastmoneyAPICmd)
	rootCmd.AddCommand(eastmoneyLimitupAPICmd)
	rootCmd.AddCommand(eastmoneyLhbAPICmd)
	rootCmd.SetHelpCommand(helpCommand)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var helpCommand = &cobra.Command{
	Use:   "help [command]",
	Short: "Help about the command",
	RunE: func(c *cobra.Command, args []string) error {
		cmd, args, e := c.Root().Find(args)
		if cmd == nil || e != nil || len(args) > 0 {
			return errors.Errorf("unknown help topic: %v", strings.Join(args, " "))
		}
		helpFunc := cmd.HelpFunc()
		helpFunc(cmd, args)
		return nil
	},
}

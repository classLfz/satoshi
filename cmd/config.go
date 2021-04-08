package cmd

import (
	"github.com/spf13/cobra"
)

var (
	configCmd = &cobra.Command{
		Use:   "config",
		Short: "Satoshi config subcommand",
	}
)

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.PersistentFlags().StringP("config", "c", "~/.satoshi/config.yaml", "Location of satoshi config file")
}

package cmd

import (
	"github.com/classlfz/satoshi/pkg/http"
	"github.com/classlfz/satoshi/pkg/siri"
	"github.com/spf13/cobra"
)

var (
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "serve controller",
		RunE:  satoshiSummon,
	}
)

func satoshiSummon(cmd *cobra.Command, args []string) error {
	fs := cmd.Flags()
	configStr, _ := fs.GetString("config")
	siriEnabled, _ := fs.GetBool("siri")
	httpEnabled, _ := fs.GetBool("http")

	go func() {
		if siriEnabled {
			siri.Start(configStr)
		}
	}()

	if httpEnabled {
		http.Start(configStr)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.PersistentFlags().StringP("config", "c", "~/.satoshi/config.yaml", "Location of satoshi config file")
	serveCmd.PersistentFlags().Bool("siri", true, "Whether to enable docking with Siri")
	serveCmd.PersistentFlags().Bool("http", true, "Whether to open http service")
}

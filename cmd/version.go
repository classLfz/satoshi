package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "satoshi version",
		RunE:  satoshiVersion,
	}
)

func satoshiVersion(cmd *cobra.Command, args []string) error {
	fmt.Printf("v0.3.0\n")

	return nil
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

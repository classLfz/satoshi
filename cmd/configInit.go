package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/classlfz/satoshi/cmd/config"
	"github.com/spf13/cobra"
)

var (
	configInitCmd = &cobra.Command{
		Use:   "init",
		Short: "Initial config file",
		RunE:  satoshiConfigInit,
	}
)

func satoshiConfigInit(cmd *cobra.Command, args []string) error {
	var loaded bool
	var err error

	fs := cmd.Flags()
	cp, _ := fs.GetString("config")
	overwrite, _ := fs.GetBool("overwrite")

	if _, loaded, err = config.Load(cp); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}

	log.Printf("overwrite: %v", overwrite)

	if !overwrite && loaded {
		return fmt.Errorf("Config already existed")
	}

	cfg := config.NewDefaultConfig()

	if err = cfg.Dump(cp); err != nil {
		return err
	}

	log.Printf("Satoshi config initialized\n")

	return nil
}

func init() {
	configCmd.AddCommand(configInitCmd)

	configInitCmd.PersistentFlags().Bool("overwrite", false, "Overwrite config file")
}

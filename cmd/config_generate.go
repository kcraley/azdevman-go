package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// configGenerateCmd represents the generate subcommand
var configGenerateCmd = &cobra.Command{
	Use:     "generate",
	Short:   "Generate a new configuration context connection.",
	Example: `  azdevman config generate`,
	RunE:    configGenerateFunc,
}

func init() {
	// Add `generate` subcommand to `config`
	configCmd.AddCommand(configGenerateCmd)
}

// configCurrentContextFunc is the main entrypoint for `azdevman config generate`
func configGenerateFunc(cmd *cobra.Command, args []string) error {
	log.Info("azdevman config generate was called")
	return nil
}

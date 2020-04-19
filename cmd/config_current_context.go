package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// configCurrentContextCmd represents the current-context subcommand
var configCurrentContextCmd = &cobra.Command{
	Use:     "current-context",
	Short:   "Retrieve the currently configured context.",
	Example: `  azdevman config current-context`,
	RunE:    configCurrentContextFunc,
}

func init() {
	// Add `current-context` subcommand to `config`
	configCmd.AddCommand(configCurrentContextCmd)
}

// configCurrentContextFunc is the main entrypoint for `azdevman config current-context`
func configCurrentContextFunc(cmd *cobra.Command, args []string) error {
	log.Info("azdevman config current-context was called")
	return nil
}

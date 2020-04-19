package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	// configCurrentContextCmd represents the current-context subcommand
	configCurrentContextCmd = &cobra.Command{
		Use:     "current-context",
		Short:   "Retrieve the name of the current connection context.",
		Example: `  azdevman config current-context`,
		Run:     configCurrentContextFunc,
	}
)

func init() {
	// Add `current-context` subcommand to `config`
	configCmd.AddCommand(configCurrentContextCmd)
}

// configCurrentContextFunc is the main entrypoint for `azdevman config current-context`
func configCurrentContextFunc(cmd *cobra.Command, args []string) {
	log.Info(options.GetCurrentContextName())
}

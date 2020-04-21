package cmd

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	// configViewCmd represents the view subcommand
	configViewCmd = &cobra.Command{
		Use:     "view",
		Short:   "Display the context of the configuration file.",
		Example: `  azdevman config view`,
		Run:     configconfigViewFunc,
	}
)

func init() {
	// Add `view` subcommand to `config`
	configCmd.AddCommand(configViewCmd)
}

// configCurrentContextFunc is the main entrypoint for `azdevman config view`
func configconfigViewFunc(cmd *cobra.Command, args []string) {
	log.Debugf("Reading configuration file %q", configFile)
	contents, err := options.ViewConfig()
	if err != nil {
		log.Fatalf("Error reading configuration file: %s", err)
	}
	fmt.Println(string(contents))
}

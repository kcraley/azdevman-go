package cmd

import "github.com/spf13/cobra"

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage the configuration to connect to various Azure DevOps organizations or projects.",
}

func init() {
	// Add `config` subcommand to `azdevman`
	rootCmd.AddCommand(configCmd)
}

package cmd

import (
	"context"

	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	ctx        context.Context
	coreClient core.Client
	err        error

	// projectCmd represents the command command
	projectCmd = &cobra.Command{
		Use:               "project",
		Short:             "Interact with projects in an organization.",
		PersistentPreRunE: createCoreClient,
	}
)

func init() {
	// Add `project` subcommand to `azdevman`
	rootCmd.AddCommand(projectCmd)

	// Add all subcommands for `azdevman project`
	projectCmd.AddCommand(projectListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// commandCmd.PersistentFlags().String("foo", "", "A help for foo")
}

func createCoreClient(cmd *cobra.Command, args []string) error {
	// Call parent command PersistentPreRun to chain PersistentPreRun
	// commands together without overriding each other
	// Ref https://github.com/spf13/cobra/issues/216
	rootCmd.PersistentPreRun(cmd, args)

	// Create a core client to handle projects
	coreClient, err = core.NewClient(ctx, connection)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

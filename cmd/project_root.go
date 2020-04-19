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

	// projectCmd represents the project subcommand
	projectCmd = &cobra.Command{
		Use:               "project",
		Short:             "Interact with projects in an organization.",
		PersistentPreRunE: createCoreClient,
	}
)

func init() {
	// Add `project` subcommand to `azdevman`
	rootCmd.AddCommand(projectCmd)
}

// createCoreClient creates a new core client used to interact with project resources
func createCoreClient(cmd *cobra.Command, args []string) error {
	// Call parent command PersistentPreRun to chain PersistentPreRun
	// commands together without overriding each other
	// Ref https://github.com/spf13/cobra/issues/216
	rootCmd.PersistentPreRun(cmd, args)

	coreClient, err = core.NewClient(ctx, connection)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

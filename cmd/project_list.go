package cmd

import (
	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var projectListCmd = &cobra.Command{
	Use:     "list",
	Short:   "Retrieves a list of projects within the organization.",
	Long:    `Retrieves a list of projects within the organization.`,
	Example: `  azdevman project list`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get first page of the list of team projects for your organization
		responseValue, err := coreClient.GetProjects(ctx, core.GetProjectsArgs{})
		if err != nil {
			log.Fatal(err)
			return err
		}

		index := 0
		for responseValue != nil {
			// Log the page of team project names
			for _, teamProjectReference := range (*responseValue).Value {
				log.WithFields(log.Fields{
					"ID":          *teamProjectReference.Id,
					"Name":        *teamProjectReference.Name,
					"Description": *teamProjectReference.Description,
				}).Info("")
				index++
			}

			// if continuationToken has a value, then there is at least one more page of projects to get
			if responseValue.ContinuationToken != "" {
				// Get next page of team projects
				projectArgs := core.GetProjectsArgs{
					ContinuationToken: &responseValue.ContinuationToken,
				}
				responseValue, err = coreClient.GetProjects(ctx, projectArgs)
				if err != nil {
					log.Fatal(err)
					return err
				}
			} else {
				responseValue = nil
			}
		}
		return nil
	},
}

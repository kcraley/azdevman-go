package cmd

import (
	"os"

	"github.com/kcraley/azdevman-go/pkg/config"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Default values for flags
const (
	defaultConfigFile = "~/.config/azdevman/config.json"
	azureDevOpsURL    = "https://dev.azure.com/"
)

var (
	// Variables for command line flags
	connection   *azuredevops.Connection
	configFile   string
	verbose      bool
	organization string
	token        string

	// azdevman root command
	rootCmd = &cobra.Command{
		Use:   "azdevman",
		Short: "azdevman manages instances of Azure DevOps",
		Long:  "azdevman is a simple CLI tool to interact with Microsoft's Azure DevOps.",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// Create a connection object everytime
			connection = azuredevops.NewPatConnection(azureDevOpsURL+organization, token)
		},
	}
)

func init() {
	// Initialize Viper config on execution
	cobra.OnInitialize(initConfig)

	// Root level flags for the cli
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "config file")
	rootCmd.PersistentFlags().StringVarP(&organization, "organization", "o", "", "sets the Azure DevOps organization to connect to")
	rootCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "the personal access token used to connect to the organization")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose logging by including additional output")
}

// initViperConfig is responsible for configuring Viper
// anytime each command's Execute method is called
func initConfig() {
	// Configure logging
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.JSONFormatter{})
	// Set verbose logging if true
	if verbose {
		log.SetLevel(log.DebugLevel)
	}

	// Initialize configuration and set values
	// if they were not set via the CLI flags
	options, err := config.Init(configFile)
	if err != nil {
		log.Fatal(err)
	}
	if organization == "" {
		organization = options.GetCurrentContext().Organization
	}
	if token == "" {
		token = options.GetCurrentContext().Token
	}
	log.Debugf("Organization: %s, Token: %s", organization, token)
}

// Execute is the entrypoint function which runs our command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

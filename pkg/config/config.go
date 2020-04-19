package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
)

const (
	projectName       = "AZDEVMAN"
	defaultConfigPath = "~/.config/azdevman/azdevman.json"
)

// Profile is a single configuration instance for an Azure DevOps organization
type Profile struct {
	Name         string `json:"name,omitempty"`
	Organization string `json:"organization,omitempty"`
	Token        string `json:"token,omitempty"`
	Project      string `json:"project,omitempty"`
}

// Context is the entire configuration which contains multiple profiles
type Context struct {
	Current  string     `json:"current,omitempty"`  // Sets the current context to a profile name
	Profiles *[]Profile `json:"profiles,omitempty"` // A list of Profiles that are configured
}

// Options represents the option set and its methods for the cli
type Options interface {
	GetCurrentContext() *Profile
	GetCurrentContextName() string
	SetCurrentContext(name string)
}

// GetCurrentContext returns the Profile type which is configured
func (c *Context) GetCurrentContext() *Profile {
	currentContextName := c.Current
	for _, context := range *c.Profiles {
		if context.Name == currentContextName {
			return &context
		}
	}
	return nil
}

// GetCurrentContextName returns the name of the current Context
func (c *Context) GetCurrentContextName() string {
	return c.Current
}

// SetCurrentContext sets the current context to an existing context
// which connects to a different organization and/or project with a token
func (c *Context) SetCurrentContext(name string) {
	c.Current = name
}

// Init initializes the config for Azdevman
func Init(configFile string) (context *Context, err error) {
	// Create a new options set
	context = &Context{}

	// Set the default config file path if not specified
	if configFile == "" {
		configFile, err = homedir.Expand(defaultConfigPath)
		if err != nil {
			log.Warnf("Unable to expand the relative config file path: %q", configFile)
			return nil, err
		}
	}

	// Read in the data from the config file
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Failed to read file: %q", configFile)
	}
	if err := json.Unmarshal(data, &context); err != nil {
		log.Fatalf("Unable to process the structure of the file: %q", configFile)
	}

	return context, nil
}

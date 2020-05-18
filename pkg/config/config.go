package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
)

const (
	projectName       = "AZDEVMAN"
	defaultConfigPath = "~/.config/azdevman/azdevman.json"
)

// Profile is a single configuration instance for an Azure DevOps organization
type Profile struct {
	Name         string `json:"name,omitempty"`         // Generic name of the profile
	Organization string `json:"organization,omitempty"` // Azure DevOps Organization perform operations against
	Token        string `json:"token,omitempty"`        // A Personal Access Token of the user making the connection
	Project      string `json:"project,omitempty"`      // Azure DevOps Project within the Organization to interact in
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
	SetCurrentContext(string) bool
	ViewConfig() ([]byte, error)
	Exists(string) bool
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
func (c *Context) SetCurrentContext(name string) bool {
	for _, v := range *c.Profiles {
		if v.Name == name {
			c.Current = name
			return true
		}
	}
	return false
}

// ViewConfig returns the entire contents of a configuration file
func (c *Context) ViewConfig() (contents []byte, err error) {
	contents, err = json.MarshalIndent(c, "", "  ")
	if err != nil {
		return nil, err
	}
	return contents, err
}

// Exists verifies that the configuration file existin on the local filesystem
func (c *Context) Exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
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

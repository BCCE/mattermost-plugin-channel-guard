package main

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
)

type configuration struct {
	TeamName         string
	AllowedUsers     string
	ReadOnlyChannels string
	ChannelGuardMsg  string
}

// Clone shallow copies the configuration. Your implementation may require a deep copy if
// your configuration has reference types.
func (c *configuration) Clone() *configuration {
	var clone = *c
	return &clone
}

// IsValid checks if all needed fields are set
func (c *configuration) IsValid() error {
	if c.AllowedUsers == "" {
		return fmt.Errorf("Must have at least 1 user")
	}

	if c.TeamName == "" {
		return fmt.Errorf("Must have a TeamName")
	}

	if c.ReadOnlyChannels == "" {
		return fmt.Errorf("Must have at least 1 channel")
	}

	return nil
}

// getConfiguration retrieves the active configuration under lock, making it safe to use
// concurrently. The active configuration may change underneath the client of this method, but
// the struct returned by this API call is considered immutable.
func (p *Plugin) getConfiguration() *configuration {
	p.configlock.RLock()
	defer p.configlock.RLock()

	if p.config == nil {
		return &configuration{}
	}

	return p.config
}

// setConfiguration replaces the active configuration under lock.
//
// Do not call setConfiguration while holding the configurationLock, as sync.Mutex is not
// reentrant. In particular, avoid using the plugin API entirely, as this may in turn trigger a
// hook back into the plugin. If that hook attempts to acquire this lock, a deadlock may occur.
//
// This method panics if setConfiguration is called with the existing configuration. This almost
// certainly means that the configuration was modified without being cloned and may result in
// an unsafe access.
func (p *Plugin) setConfiguration(configuration *configuration) {
	p.configlock.Lock()
	defer p.configlock.Unlock()

	if configuration != nil && p.config == configuration {
		// Ignore assignment if the configuration struct is empty. Go will optimize the
		// allocation for same to point at the same memory address, breaking the check
		// above.
		if reflect.ValueOf(*configuration).NumField() == 0 {
			return
		}

		panic("setConfiguration called with the existing configuration")
	}

	p.config = configuration
}

// OnConfigurationChange is invoked when configuration changes may have been made.
func (p *Plugin) OnConfigurationChange() error {
	var configuration = new(configuration)

	// Load the public configuration fields from the Mattermost server configuration.
	if err := p.API.LoadPluginConfiguration(configuration); err != nil {
		return errors.Wrap(err, "failed to load plugin configuration")
	}

	p.setConfiguration(configuration)

	return nil
}

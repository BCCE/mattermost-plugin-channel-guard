package main

type Configuration struct {
	Guards []*ConfigGuard
}

type ConfigGuard struct {
	TeamName string

	ChannelName string

	Allowed []string
}

// OnConfigurationChange is invoked when configuration changes may have been made.
//
// This demo implementation ensures the configured demo user and channel are created for use
// by the plugin.
func (p *guard) OnConfigurationChange() error {
	var c Configuration

	if err := p.API.LoadPluginConfiguration(&c); err != nil {
		p.API.LogError(err.Error())
		return err
	}

	p.guards.Store(c.Guards)

	return nil

}

// List of the welcome messages from the configuration

func (p *guard) getGuards() []*ConfigGuard {
	return p.guards.Load().([]*ConfigGuard)
}

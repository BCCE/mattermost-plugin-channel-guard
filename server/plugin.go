package main

import (
	"io/ioutil"
	"path/filepath"
	"sync"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
	"github.com/pkg/errors"
)

const (
	botUsername    = "guardbot"
	botDisplayName = "GuardBot"
	botDescription = "Guards the channels so you dont have to"
)

// Plugin struct for the plugin
type Plugin struct {
	plugin.MattermostPlugin

	configlock sync.RWMutex

	botUserID string

	config *configuration
}

// OnActivate registers the command
func (p *Plugin) OnActivate() error {
	bot := &model.Bot{
		Username:    botUsername,
		DisplayName: botDisplayName,
		Description: botDescription,
	}

	botUserID, apperr := p.Helpers.EnsureBot(bot)
	if apperr != nil {
		return errors.Wrap(apperr, "failed to ensure bot user")
	}

	if err := p.setBotIcon(botUserID); err != nil {
		p.API.LogWarn("Failed to set profile image for bot", "err", err)
	}

	p.botUserID = botUserID

	return nil
}

func (p *Plugin) setBotIcon(botUserID string) *model.AppError {
	bundlePath, err := p.API.GetBundlePath()
	p.Check(err)

	icon, err := ioutil.ReadFile(filepath.Join(bundlePath, "assets", "icon.png"))
	p.Check(err)

	return p.API.SetProfileImage(botUserID, icon)
}

// Check takes care of error handling
func (p *Plugin) Check(e error) *model.AppError {
	if e != nil {
		return &model.AppError{Message: e.Error()}
	}
	return nil
}

func main() {
	plugin.ClientMain(&Plugin{})
}

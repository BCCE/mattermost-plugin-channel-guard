package main

import (
	"sync/atomic"

	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"
	"github.com/pkg/errors"
)

const (
	botUsername    = "guardbot"
	botDisplayName = "GuardBot"
	botDescription = "Guards the channels so you dont have to"
)

// LanBotPlugin is the main plugin struct
type guard struct {
	plugin.MattermostPlugin
	botUserID string

	guards atomic.Value
}

// OnActivate registers the command
func (p *guard) OnActivate() error {
	bot := &model.Bot{
		Username:    botUsername,
		DisplayName: botDisplayName,
		Description: botDescription,
	}
	botUserID, apperr := p.Helpers.EnsureBot(bot)
	if apperr != nil {
		return errors.Wrap(apperr, "failed to ensure bot user")
	}

	p.botUserID = botUserID

	return nil
}

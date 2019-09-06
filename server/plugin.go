package main

import (
	"path/filepath"
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

	readFile func(path string) ([]byte, error)

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

	if err := p.setBotProfileImage(bot.UserId); err != nil {
		p.API.LogWarn("Failed to set profile image for bot", "err", err)
	}

	p.botUserID = botUserID

	return nil
}

func (p *guard) setBotIcon(botUserID string) *model.AppError {
	bundlePath, err := p.API.GetBundlePath()
	if err != nil {
		return &Model.AppError{Message: err.Error()}
	}
	icon, err := p.readFile(filepath.Join(bundlePath, "assets", "icon.png"))
	if err != nil {
		return &Model.AppError{Message: err.Error()}
	}

	return p.API.SetProfileImage(botUserID, icon)
}
